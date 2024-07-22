package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetCareerApplicantData(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetCareerApplicantData(&fiber.Ctx{})
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func ApplyCareer(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.ApplyCareer(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
