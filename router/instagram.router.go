package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func InstagramRouter(app *fiber.App) {
	instagramRouter := app.Group("/instagram")
	instagramRouter.All("/feed", middleware.BasicAuthMiddleware(), controller.GetInstagramFeed)
}
