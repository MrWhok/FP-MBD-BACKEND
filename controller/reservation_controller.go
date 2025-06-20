package controller

import (
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"github.com/gofiber/fiber/v2"
)

type ReservationController struct {
	Service service.ReservationService
}

func NewReservationController(service service.ReservationService) *ReservationController {
	return &ReservationController{Service: service}
}

func (r *ReservationController) Route(app *fiber.App) {
	app.Post("/v1/api/reservation", r.MakeReservation)
}

func (r *ReservationController) MakeReservation(c *fiber.Ctx) error {
	var req model.ReservationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
		})
	}

	customerID := c.Locals("user_id").(int) // make sure your JWT middleware sets this

	err := r.Service.Reserve(c.Context(), customerID, req)
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
