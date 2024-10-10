package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetCareerApplicantData(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetCareerApplicantData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func ValidateApplyCareer(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.ValidateApplyCareer(ctx)
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

func GetOrUpsertCareer(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetAllCareer(ctx)
	} else if ctx.Method() == "PUT" || ctx.Method() == "PATCH" {
		return service.UpsertCareer(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func GetActiveCareer(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetActiveCareer(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
