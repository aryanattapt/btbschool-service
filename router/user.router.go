package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	UserRouter := app.Group("/users")
	UserRouter.All("/search", middleware.JWTAuthMiddleware(), controller.GetAllUser)
	UserRouter.All("/", middleware.JWTAuthMiddleware(), controller.UpsertUser)
}
