package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactInsertPayload struct {
	FirstName      string             `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	LastName       string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Phoneno        string             `bson:"phoneno,omitempty" json:"phoneno,omitempty" validate:"required,e164"`
	Email          string             `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Message        string             `bson:"message,omitempty" json:"message,omitempty" validate:"required,min=3,max=30"`
	RegisteredDate primitive.DateTime `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	UpdatedDate    primitive.DateTime `bson:"updateddate,omitempty" json:"updateddate,omitempty"`
}
