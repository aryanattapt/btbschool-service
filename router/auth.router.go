package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	authRouter := app.Group("/auth")
	authRouter.All("/signin", middleware.BasicAuthMiddleware(), controller.SignIn)
	authRouter.All("/signup", middleware.BasicAuthMiddleware(), controller.SignUp)
	authRouter.All("/validate", middleware.JWTAuthMiddleware(), controller.Validate)
	authRouter.All("/user", middleware.BasicAuthMiddleware(), controller.GetOrUpdateUser)
}
