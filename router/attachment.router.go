package router

import (
	"btb-service/controller"
	"btb-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func AttachmentsRouter(app *fiber.App) {
	attachmentRouter := app.Group("/attachment", middleware.BasicAuthMiddleware())
	attachmentRouter.All("/upload/:param", controller.UploadAttachment)
	attachmentRouter.All("/uploadassets/:param", controller.UploadAssets)
}
