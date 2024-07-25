package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRegistrationRouter(app *fiber.App) {
	studentRegistrationRouter := app.Group("/student/registration")
	studentRegistrationRouter.All("/outstanding", middleware.BasicAuthMiddleware(), controller.GetStudentRegistrationOutstandingData)
	studentRegistrationRouter.All("/submit", middleware.BasicAuthMiddleware(), controller.SubmitDataStudentRegistration)
}
