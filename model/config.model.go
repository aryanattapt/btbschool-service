package model

type ConfigRequestPayload struct {
	Payload []interface{} `json:"data" validate:"required"`
}

type ConfigModel struct {
	Type    string
	Payload []map[string]interface{}
}
