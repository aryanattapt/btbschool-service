package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func AlumniRouter(app *fiber.App) {
	alumniRouter := app.Group("/alumni")
	alumniRouter.All("/", middleware.BasicAuthMiddleware(), controller.GetAlumni)
	alumniRouter.All("/submit", middleware.BasicAuthMiddleware(), controller.SubmitAlumni)
}
