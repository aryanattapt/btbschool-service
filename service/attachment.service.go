package service

import (
	"btb-service/pkg"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
)

func UploadAttachment(ctx *fiber.Ctx) error {
	var param string = ctx.Params("param")
	if pkg.IsEmptyString(param) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "ATTACHMENT.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
			"message": "Failed to submit attachment data!",
			"error":   err.Error(),
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
				"category":     param,
				"type":         formFieldName,
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

func UploadAttachmentS3(ctx *fiber.Ctx) error {
	var param string = ctx.Params("param")
	if pkg.IsEmptyString(param) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "ATTACHMENT.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
			"message": "Failed to submit attachment data!",
			"error":   err.Error(),
		})
	}

	session, err := CreateSessionS3()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
			"message": "Failed to open session!",
			"error":   err.Error(),
		})
	}

	var data []map[string]interface{}
	for formFieldName, fileHeaders := range form.File {
		for _, file := range fileHeaders {
			filename := strings.Replace(pkg.GenerateUUID(), "-", "", -1)
			fileExt := path.Ext(file.Filename)
			fileName := fmt.Sprintf("%s.%s", filename, fileExt)

			theFile, err := file.Open()
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
					"message": "Failed to open attachment data!",
					"error":   err.Error(),
				})
			}
			defer theFile.Close()

			uploader := s3manager.NewUploader(session)
			_, err = UploadFileS3(uploader, theFile, "uploads", fmt.Sprintf("%s/%s/%s", param, formFieldName, fileName))
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
					"message": "Failed to upload file!",
					"error":   err.Error(),
				})
			}

			temp := map[string]interface{}{
				"fileName":     fileName,
				"fileURL":      fmt.Sprintf("%s/%s/%s/%s", "https://w6i8.c1.e2-7.dev/uploads", param, formFieldName, fileName),
				"fileMetadata": file.Header,
				"fileSize":     file.Size,
				"category":     param,
				"type":         formFieldName,
			}
			data = append(data, temp)
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit attachment data!",
		"data":    data,
	})
}

func UploadAssetsS3(ctx *fiber.Ctx) error {
	var param string = ctx.Params("param")
	if pkg.IsEmptyString(param) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "ATTACHMENT.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
			"message": "Failed to submit attachment data!",
			"error":   err.Error(),
		})
	}

	session, err := CreateSessionS3()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
			"message": "Failed to open session!",
			"error":   err.Error(),
		})
	}

	var data []map[string]interface{}
	for formFieldName, fileHeaders := range form.File {
		for _, file := range fileHeaders {
			filename := strings.Replace(pkg.GenerateUUID(), "-", "", -1)
			fileExt := path.Ext(file.Filename)
			fileName := fmt.Sprintf("%s.%s", filename, fileExt)

			theFile, err := file.Open()
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
					"message": "Failed to open attachment data!",
					"error":   err.Error(),
				})
			}
			defer theFile.Close()

			uploader := s3manager.NewUploader(session)
			_, err = UploadFileS3(uploader, theFile, "assets", fmt.Sprintf("%s/%s/%s", param, formFieldName, fileName))
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "ATTACHMENT.UPLOAD.EXCEPTION",
					"message": "Failed to upload file!",
					"error":   err.Error(),
				})
			}

			temp := map[string]interface{}{
				"fileName":     fileName,
				"fileURL":      fmt.Sprintf("%s/%s/%s/%s", "https://w6i8.c1.e2-7.dev/uploads", param, formFieldName, fileName),
				"fileMetadata": file.Header,
				"fileSize":     file.Size,
				"category":     param,
				"type":         formFieldName,
			}
			data = append(data, temp)
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit attachment data!",
		"data":    data,
	})
}
