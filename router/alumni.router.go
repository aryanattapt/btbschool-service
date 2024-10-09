package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func AlumniRouter(app *fiber.App) {
	alumniRouter := app.Group("/alumni")
	alumniRouter.All("/fetch", middleware.JWTAuthMiddleware(), controller.GetAlumni)
	alumniRouter.All("/submit", middleware.BasicAuthMiddleware(), controller.SubmitAlumni)
	alumniRouter.All("/submit/validate", middleware.BasicAuthMiddleware(), controller.ValidateAlumni)
	alumniRouter.All("/verify", middleware.JWTAuthMiddleware(), controller.VerifyAlumni)
}
