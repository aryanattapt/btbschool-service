package model

type SecretKeyPayload struct {
	Payload map[string]interface{} `json:"data" validate:"required"`
}
