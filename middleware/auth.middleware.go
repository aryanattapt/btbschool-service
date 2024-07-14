package middleware

import (
	"btb-service/service"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"btbschool": "btbschool",
		},
		Realm: "Forbidden",
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Sorry, Unauthorized Access",
				"error":   "INVALID_AUTH",
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})
}

func JWTAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: service.JWT_SIGNATURE_KEY},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Sorry, Unauthorized Access",
				"error":   "INVALID_AUTH",
			})
		},
	})
}
