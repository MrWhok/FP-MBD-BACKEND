package controller

import (
	"context"
	"database/sql"
	"errors"
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
	app.Put("/v1/api/reservation/reschedule", middleware.AuthenticateJWT("customer", r.Config), r.RescheduleReservation)
	app.Put("/v1/api/reservation/re", middleware.AuthenticateJWT("customer", r.Config), r.RescheduleReservation)
	app.Put("/v1/api/reservation/edit", middleware.AuthenticateJWT("customer", r.Config), r.EditReservation)
	app.Delete("/v1/api/reservation/:id", middleware.AuthenticateJWT("customer", r.Config), r.CancelReservation)
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
			Data:    "Failed to parse customer_id from token.",
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

func (r *ReservationController) EditReservation(c *fiber.Ctx) error {
	var req model.EditReservationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
			Data:    err.Error(),
		})
	}

	err := r.ReservationService.EditReservation(c.Context(), req.ReservationID, req.NewGuests)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Reservation successfully edited",
		Data:    fmt.Sprintf("Reservation ID %d edited successfully.", req.ReservationID),
	})
}

// RescheduleReservation handles the request to reschedule an existing reservation.
func (r *ReservationController) RescheduleReservation(c *fiber.Ctx) error {
	var req model.RescheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
			Data:    err.Error(),
		})
	}

	customerIDValue := c.Locals("customer_id")
	customerID, ok := customerIDValue.(int)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "General Error",
			Data:    "Failed to parse customer_id from token.",
		})
	}

	err := r.ReservationService.Reschedule(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Reservation successfully rescheduled",
		Data:    fmt.Sprintf("Reservation ID %d rescheduled successfully.", req.ReservationID),
	})
}

func (r *ReservationController) CancelReservation(c *fiber.Ctx) error {
	reservationID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid reservation ID",
			Data:    err.Error(),
		})
	}

	customerIDValue := c.Locals("customer_id")
	customerID, ok := customerIDValue.(int)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "General Error",
			Data:    "Failed to parse customer_id from token.",
		})
	}

	err = r.ReservationService.CancelReservation(c.Context(), reservationID, customerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Reservasi berhasil dibatalkan",
	})
}
