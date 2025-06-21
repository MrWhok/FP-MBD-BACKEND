package controller

import (
	"strconv"

	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/middleware"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	service.PaymentService
	configuration.Config
}

func NewPaymentController(paymentService service.PaymentService, config configuration.Config) *PaymentController {
	return &PaymentController{
		PaymentService: paymentService,
		Config:         config,
	}
}

func (p *PaymentController) Route(app *fiber.App) {
	app.Post("/v1/api/payment/upload/:reservation_id", middleware.AuthenticateJWT("customer", p.Config), p.UploadProof)
	app.Post("/v1/api/payment/confirm/:reservation_id", middleware.AuthenticateJWT("admin", p.Config), p.ConfirmPayment)
}

func (p *PaymentController) UploadProof(c *fiber.Ctx) error {
	reservationIDStr := c.Params("reservation_id")
	reservationID, err := strconv.Atoi(reservationIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid reservation ID",
		})
	}

	// Ambil file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "File not found in request",
		})
	}

	// Proses upload
	err = p.PaymentService.UploadPaymentProof(c.Context(), reservationID, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Payment proof uploaded successfully",
	})
}

func (p *PaymentController) ConfirmPayment(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("reservation_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid reservation ID",
		})
	}

	err = p.PaymentService.ConfirmPayment(c.Context(), reservationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Payment successfully confirmed",
	})
}
