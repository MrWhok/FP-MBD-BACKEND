package middleware

import (
	"github.com/MrWhok/FP-MBD-BACKEND/common"
	"github.com/MrWhok/FP-MBD-BACKEND/configuration"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(role string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// ✅ Extract customer_id from JWT and set it to ctx.Locals
			if customerIDFloat, ok := claims["customer_id"].(float64); ok {
				customerID := int(customerIDFloat)
				ctx.Locals("customer_id", customerID)
				common.NewLogger().Info("✅ Set customer_id to Locals:", customerID)
			} else {
				common.NewLogger().Error("❌ Failed to extract customer_id from claims")
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid or missing customer ID",
				})
			}

			// ✅ Check role
			roles, ok := claims["roles"].([]interface{})
			if !ok {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Missing roles",
				})
			}

			common.NewLogger().Info("Required role:", role, " | User roles:", roles)
			for _, roleInterface := range roles {
				if roleMap, ok := roleInterface.(map[string]interface{}); ok {
					if roleMap["role"] == role {
						return ctx.Next()
					}
				}
			}

			return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid Role",
			})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
					Code:    400,
					Message: "Bad Request",
					Data:    "Missing or malformed JWT",
				})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid or expired JWT",
				})
			}
		},
	})
}
