package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// RecaptchaResponse represents the response from the Google reCAPTCHA API
type RecaptchaResponse struct {
	Success      bool    `json:"success"`
	ChallengeTs  string  `json:"challenge_ts"`
	Hostname     string  `json:"hostname"`
	Score        float64 `json:"score"`
	Action       string  `json:"action"`
	ScoreMessage string  `json:"score_message"`
}

func ValidateRecaptchaGoogle(ctx *fiber.Ctx) error {
	// Initialize an empty map to hold the parsed body
	var payload map[string]interface{}

	// Parse the request body into the map
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "GOOGLE.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, system can't parse your data! Please recheck!",
			"error":   err.Error(),
		})
	}

	// Extract the "data" field from the map and assert its type
	data, ok := payload["data"].(string)
	if !ok {
		log.Println("data field is either missing or not a string")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "GOOGLE.INVALIDPAYLOAD.DATA_FIELD_INVALID",
			"message": "Sorry, system can't parse your data! Please recheck!",
		})
	}

	// Prepare the request to Google's reCAPTCHA API
	finalPayload := fmt.Sprintf("secret=%s&response=%s", os.Getenv("RECAPTCHA_SECRET_KEY"), data)
	req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", bytes.NewBufferString(finalPayload))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "GOOGLE.REQUEST_ERROR",
			"message": "Failed to create request to Google reCAPTCHA.",
			"error":   err.Error(),
		})
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "GOOGLE.REQUEST_ERROR",
			"message": "Failed to send request to Google reCAPTCHA.",
			"error":   err.Error(),
		})
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "GOOGLE.RESPONSE_ERROR",
			"message": "Failed to read response from Google reCAPTCHA.",
			"error":   err.Error(),
		})
	}

	// Parse the JSON response
	var recaptchaResponse RecaptchaResponse
	if err := json.Unmarshal(body, &recaptchaResponse); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "GOOGLE.PARSE_ERROR",
			"message": "Failed to parse response from Google reCAPTCHA.",
			"error":   err.Error(),
		})
	}

	// Check the success status and return an appropriate response
	if !recaptchaResponse.Success {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    "GOOGLE.INVALID_RECAPTCHA",
			"message": "Invalid reCAPTCHA verification.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "reCAPTCHA validation successful",
		"data":    recaptchaResponse,
	})
}
