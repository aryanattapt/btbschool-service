package router

import (
	"btb-service/controller"

	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {
	configRouter := app.Group("/config")
	configRouter.All("/:type", controller.ConfigController)
}
