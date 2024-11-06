package model

type SignInPayload struct {
	Username string `json:"username" validate:"required,alpha,min=3,max=20"`
	Password string `json:"password" validate:"required,alphanum,min=3,max=20"`
}

type UserCheckPermissionPayload struct {
	Permission string `json:"permission,omitempty" validate:"required"`
}
