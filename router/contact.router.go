package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func ContactRouter(app *fiber.App) {
	contactRouter := app.Group("/contact")
	contactRouter.All("/", middleware.BasicAuthMiddleware(), controller.GetContacts)
	contactRouter.All("/submit", middleware.BasicAuthMiddleware(), controller.SaveContacts)
}
