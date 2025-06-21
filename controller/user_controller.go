package controller

import (
	"github.com/MrWhok/FP-MBD-BACKEND/common"
	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/register", controller.Register)
	app.Post("/v1/api/login", controller.Login)
}

func (controller UserController) Register(c *fiber.Ctx) error {
	var request model.UserRegisterModel
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
		})
	}

	err := controller.UserService.Register(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "User registered successfully",
	})
}

func (controller UserController) Login(c *fiber.Ctx) error {
	var request model.UserLoginModel
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid request format",
		})
	}

	customerID, role, err := controller.UserService.Login(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:    401,
			Message: "Email or password incorrect",
		})
	}

	token := common.GenerateToken(customerID, role, controller.Config)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Login successful",
		Data: map[string]interface{}{
			"token":       token,
			"customer_id": customerID,
			"role":        role,
		},
	})
}
