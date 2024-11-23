package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserViewPayload struct {
	FirstName      string             `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	LastName       string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Username       string             `bson:"username,omitempty" json:"username,omitempty" validate:"required,alphanum,min=3,max=20"`
	Email          string             `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	RegisteredDate primitive.DateTime `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	UpdatedDate    primitive.DateTime `bson:"updateddate,omitempty" json:"updateddate,omitempty"`
	Role           string             `bson:"role,omitempty" json:"role,omitempty"`
	IsActive       string             `bson:"isactive,omitempty" json:"isactive,omitempty"`
	Permission     []interface{}      `bson:"permission,omitempty" json:"permission,omitempty" validate:"required"`
}

type UserInsertPayload struct {
	Password string `bson:"password,omitempty" json:"password,omitempty" validate:"required,alphanum,min=3,max=20"`
	UserViewPayload
}

type UserUpdatePayload struct {
	ID string `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	UserViewPayload
}
