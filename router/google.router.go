package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func GoogleRouter(app *fiber.App) {
	googleRouter := app.Group("/google")
	googleRouter.All("/validatecaptcha", middleware.BasicAuthMiddleware(), controller.ValidateRecaptchaGoogle)
}
