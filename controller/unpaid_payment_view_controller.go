package controller

import (
	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/middleware"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"github.com/gofiber/fiber/v2"
)

type UnpaidPaymentController struct {
	service service.UnpaidPaymentService
	config  configuration.Config
}

func NewUnpaidPaymentController(service service.UnpaidPaymentService, config configuration.Config) *UnpaidPaymentController {
	return &UnpaidPaymentController{
		service: service,
		config:  config,
	}
}

// ✅ Register route(s) here
func (c *UnpaidPaymentController) Route(app *fiber.App) {
	app.Get("/v1/api/unpaid-payments", middleware.AuthenticateJWT("admin", c.config), c.GetUnpaidPayments)
}

func (c *UnpaidPaymentController) GetUnpaidPayments(ctx *fiber.Ctx) error {
	data, err := c.service.GetAllUnpaidPayments()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch unpaid payments",
		})
	}
	return ctx.JSON(data)
}
