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
				"code":    "AUTH.INVALIDTOKEN.EXCEPTION",
				"message": "Sorry, Unauthorized Access",
				"error":   "Invalid token",
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
				"code":    "AUTH.INVALIDTOKEN.EXCEPTION",
				"message": "Sorry, Unauthorized Access",
				"error":   err.Error(),
			})
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			var splitedToken = strings.Split(string(ctx.Request().Header.Peek("Authorization")), " ")
			if len(splitedToken) == 0 {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    "AUTH.INVALIDTOKEN.EXCEPTION",
					"message": "Invalid Token!",
					"error":   "Invalid Pattern!",
				})
			}

			result, err := service.ValidateJWTToken(splitedToken[1])
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    "AUTH.INVALIDTOKEN.EXCEPTION",
					"message": "Invalid Token!",
					"error":   err.Error(),
				})
			}

			userDataList, err := repository.GetUserById(pkg.DecodeBase64(result["aud"].(string)))
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "AUTH.USERQUERY.EXCEPTION",
					"message": "Failed to validate user data!",
					"error":   err.Error(),
				})
			}

			if len(userDataList) == 0 {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    "AUTH.USERNOTEXIST.EXCEPTION",
					"message": "Sorry, User is not exist in database!",
					"error":   "User is not exist",
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
