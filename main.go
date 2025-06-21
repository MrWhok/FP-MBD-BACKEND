package main

import (
	"log"
	"time"

	"github.com/MrWhok/FP-MBD-BACKEND/client/restclient"
	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/controller"
	_ "github.com/MrWhok/FP-MBD-BACKEND/docs"
	"github.com/MrWhok/FP-MBD-BACKEND/exception"
	"github.com/MrWhok/FP-MBD-BACKEND/repository/impl"
	repository "github.com/MrWhok/FP-MBD-BACKEND/repository/impl"
	service "github.com/MrWhok/FP-MBD-BACKEND/service/impl"
	"github.com/MrWhok/FP-MBD-BACKEND/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Go Fiber Clean Architecture
// @version 1.0.0
// @description Baseline project using Go Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	// redis := configuration.NewRedis(config)

	//repository
	userRepository := repository.NewUserRepositoryImpl(database)
	reservationRepo := impl.NewReservationRepositoryImpl(database)
	unpaidPaymentRepo := impl.NewUnpaidPaymentRepositoryImpl(database)

	//rest client
	httpBinRestClient := restclient.NewHttpBinRestClient()

	//service
	userService := service.NewUserServiceImpl(&userRepository)
	httpBinService := service.NewHttpBinServiceImpl(&httpBinRestClient)
	reservationService := service.NewReservationServiceImpl(reservationRepo)
	unpaidPaymentService := service.NewUnpaidPaymentServiceImpl(unpaidPaymentRepo)

	//controller
	userController := controller.NewUserController(&userService, config)
	httpBinController := controller.NewHttpBinController(&httpBinService)
	reservationController := controller.NewReservationController(reservationService, config)
	unpaidPaymentController := controller.NewUnpaidPaymentController(unpaidPaymentService, config)

	notificationRepo := impl.NewNotificationRepositoryImpl(database)

	// Start background email sender
	go func() {
		for {
			notifications, err := notificationRepo.GetUnsentNotifications()
			if err != nil {
				log.Println("❌ Error fetching unsent notifications:", err)
			}

			for _, notif := range notifications {
				err := utils.SendEmail(notif.Email, "Reservasi Anda", notif.Message)
				if err != nil {
					log.Println("❌ Failed to send email to", notif.Email, ":", err)
					continue
				}
				err = notificationRepo.MarkAsSent(notif.ID)
				if err != nil {
					log.Println("❌ Failed to mark notification as sent:", err)
				}
			}

			time.Sleep(30 * time.Second) // check setiap 30 detik
		}
	}()

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	userController.Route(app)
	httpBinController.Route(app)
	reservationController.Route(app)
	unpaidPaymentController.Route(app)

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
