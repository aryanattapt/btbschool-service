package middleware

import (
	"btb-service/pkg"
	"btb-service/repository"
	"btb-service/service"
	"strings"

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
				"error":      "AUTH.INVALIDTOKEN.EXCEPTION",
				"message":    "Sorry, Unauthorized Access",
				"stacktrace": "Invalid token",
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
				"error":      "AUTH.INVALIDTOKEN.EXCEPTION",
				"message":    "Sorry, Unauthorized Access",
				"stacktrace": err.Error(),
			})
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			result, err := service.ValidateJWTToken(strings.Split(string(ctx.Request().Header.Peek("Authorization")), " ")[1])
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":      "AUTH.INVALIDTOKEN.EXCEPTION",
					"message":    "Invalid Token!",
					"stacktrace": err.Error(),
				})
			}

			userDataList, err := repository.GetUserById(pkg.DecodeBase64(result["aud"].(string)))
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":      "AUTH.USERQUERY.EXCEPTION",
					"message":    "Failed to validate user data!",
					"stacktrace": err.Error(),
				})
			}

			if len(userDataList) == 0 {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":      "AUTH.USERNOTEXIST.EXCEPTION",
					"message":    "Sorry, User is not exist in database!",
					"stacktrace": "User is not exist",
				})
			}

			var userData map[string]interface{} = userDataList[0]
			delete(userData, "password")
			delete(userData, "isactive")

			ctx.Locals("userauth", userData)
			return ctx.Next()
		},
	})
}
