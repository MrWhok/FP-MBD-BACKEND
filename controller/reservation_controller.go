package controller

import (
	"fmt"

	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/middleware"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"github.com/gofiber/fiber/v2"
)

type ReservationController struct {
	service.ReservationService
	configuration.Config
}

func NewReservationController(reservationService service.ReservationService, config configuration.Config) *ReservationController {
	return &ReservationController{ReservationService: reservationService, Config: config}
}

func (r *ReservationController) Route(app *fiber.App) {
	app.Post("/v1/api/reservation", middleware.AuthenticateJWT("customer", r.Config), r.MakeReservation)
}

func (r *ReservationController) MakeReservation(c *fiber.Ctx) error {
	var req model.ReservationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
		})
	}

	customerIDValue := c.Locals("customer_id")
	customerID, ok := customerIDValue.(int)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "General Error",
			Data:    "Failed to parse customer_id from token. Got type: " + fmt.Sprintf("%T", customerIDValue),
		})
	}
	err := r.ReservationService.Reserve(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Reservation created successfully",
	})
}
