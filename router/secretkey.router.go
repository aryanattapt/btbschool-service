package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func SecretKeyRouter(app *fiber.App) {
	secretKeyRouter := app.Group("/secretkey", middleware.JWTAuthMiddleware())
	secretKeyRouter.All("/emailconfig/fetch", controller.GetEmailConfig)
	secretKeyRouter.All("/emailconfig/update", controller.UpdateEmailConfig)
	secretKeyRouter.All("/recaptcha/fetch", controller.GetRecaptchaConfig)
	secretKeyRouter.All("/recaptcha/update", controller.UpdateRecaptchaConfig)
	secretKeyRouter.All("/instagram/fetch", controller.GetInstagramConfig)
	secretKeyRouter.All("/instagram/update", controller.UpdateInstagramConfig)
}
