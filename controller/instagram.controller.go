package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetInstagramFeed(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetInstagramFeed(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
