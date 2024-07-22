package service

import (
	"btb-service/pkg"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UploadAttachment(ctx *fiber.Ctx) error {
	var param string = ctx.Params("param")
	if pkg.IsEmptyString(param) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "ATTACHMENT.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Please define Param!",
			"stacktrace": "Param is not exist",
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ATTACHMENT.UPLOAD.EXCEPTION",
			"message":    "Failed to submit attachment data!",
			"stacktrace": err.Error(),
		})
	}

	var data []map[string]interface{}
	for formFieldName, fileHeaders := range form.File {
		for _, file := range fileHeaders {
			filename := strings.Replace(pkg.GenerateUUID(), "-", "", -1)
			fileExt := path.Ext(file.Filename)
			fileName := fmt.Sprintf("%s.%s", filename, fileExt)

			os.MkdirAll(fmt.Sprintf("%s/%s/%s", "./uploads", param, formFieldName), os.ModePerm)
			err = ctx.SaveFile(file, fmt.Sprintf("%s/%s/%s/%s", "./uploads", param, formFieldName, fileName))

			temp := map[string]interface{}{
				"fileName":     fileName,
				"fileURL":      fmt.Sprintf("%s/%s/%s/%s/%s", os.Getenv("API_BASE_URL"), "uploads", param, formFieldName, fileName),
				"fileMetadata": file.Header,
				"fileSize":     file.Size,
			}

			if err == nil {
				data = append(data, temp)
			} else {
				log.Printf("Failed upload data: %s", err.Error())
			}
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit attachment data!",
		"data":    data,
	})
}
