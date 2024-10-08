package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func CareerRouter(app *fiber.App) {
	careerRouter := app.Group("/career")
	careerRouter.All("/applicant", middleware.JWTAuthMiddleware(), controller.GetCareerApplicantData)
	careerRouter.All("/apply", middleware.BasicAuthMiddleware(), controller.ApplyCareer)
	careerRouter.All("/apply/validate", middleware.BasicAuthMiddleware(), controller.ValidateApplyCareer)
	careerRouter.All("/active", middleware.BasicAuthMiddleware(), controller.GetActiveCareer)
	careerRouter.All("/", middleware.BasicAuthMiddleware(), controller.GetOrUpsertCareer)
}
