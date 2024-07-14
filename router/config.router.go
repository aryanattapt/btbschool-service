package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {
	configRouter := app.Group("/config", middleware.JWTAuthMiddleware())
	configRouter.All("/:type", controller.ConfigController)
}
