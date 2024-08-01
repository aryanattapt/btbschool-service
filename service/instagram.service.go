package service

import (
	"btb-service/pkg"
	"btb-service/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RefreshTokenInstagram() {
	log.Println("Refresh instagram token")
	oldToken, err := repository.GetInstagramToken()
	if err == nil {
		result, err := pkg.CallURLGet(fmt.Sprintf("%s%s", "https://graph.instagram.com/refresh_access_token?grant_type=ig_refresh_token&access_token=", oldToken))
		log.Println("Hasil Refresh Token: ", result)
		if err != nil {
			log.Println("Failed to refresh token : ", err.Error())
		}

		var resultInJSON map[string]interface{}
		err = json.Unmarshal([]byte(result), &resultInJSON)
		if err != nil {
			log.Println("Failed to decode json : ", err.Error())
		} else {
			newToken, _ := resultInJSON["access_token"].(string)
			err = repository.UpdateInstagramToken(newToken)
			if err != nil {
				log.Println("Failed update token : ", err.Error())
			}
		}
	} else {
		log.Println("Failed get old token : ", err.Error())
	}
}

func GetInstagramFeed(ctx *fiber.Ctx) error {
	token, err := repository.GetInstagramToken()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "IG.QUERY.EXCEPTION",
			"message":    "Sorry, Failed get instagram data!",
			"stacktrace": err.Error(),
		})
	}

	result, err := pkg.CallURLGet(fmt.Sprintf("%s%s", "https://graph.instagram.com/me/media?fields=id,caption,media_type,media_url,thumbnail_url,permalink,timestamp&access_token=", token))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "IG.QUERY.EXCEPTION",
			"message":    "Sorry, Failed get instagram data!",
			"stacktrace": err.Error(),
		})
	}

	var resultInJSON map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultInJSON)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "IG.QUERY.EXCEPTION",
			"message":    "Sorry, Failed get instagram data!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get instagram feed data!",
		"data":    resultInJSON["data"],
	})
}
