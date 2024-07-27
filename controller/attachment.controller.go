package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func UploadAttachment(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.UploadAttachmentS3(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func UploadAssets(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.UploadAssetsS3(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
