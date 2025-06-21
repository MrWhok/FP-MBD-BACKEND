package middleware

import (
	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(requiredRole string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Extract customer_id
			if customerIDFloat, ok := claims["customer_id"].(float64); ok {
				ctx.Locals("customer_id", int(customerIDFloat))
			} else {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid or missing customer ID",
				})
			}

			// Extract and check role
			role, ok := claims["role"].(string)
			if !ok || role != requiredRole {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid Role",
				})
			}

			return ctx.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    err.Error(),
			})
		},
	})
}
