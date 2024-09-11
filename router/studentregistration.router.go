package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRegistrationRouter(app *fiber.App) {
	studentRegistrationRouter := app.Group("/student/registration")
	studentRegistrationRouter.All("/draft", middleware.BasicAuthMiddleware(), controller.GetDraftStudentRegistrationgData)
	studentRegistrationRouter.All("/submit", middleware.BasicAuthMiddleware(), controller.SubmitDataStudentRegistration)
	studentRegistrationRouter.All("/validate", middleware.BasicAuthMiddleware(), controller.ValidateDataStudentRegistration)
	studentRegistrationRouter.All("/outstanding", middleware.JWTAuthMiddleware(), controller.GetStudentRegistrationOutstandingData)
	studentRegistrationRouter.All("/outstanding/approve", middleware.JWTAuthMiddleware(), controller.ApproveStudentRegistrationOutstandingData)
	studentRegistrationRouter.All("/all", middleware.JWTAuthMiddleware(), controller.GetAllStudentRegistrationAuthData)
	studentRegistrationRouter.All("/detail", middleware.JWTAuthMiddleware(), controller.GetStudentRegistrationDetailData)
}
