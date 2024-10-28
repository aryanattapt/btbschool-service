package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CareerApplyInsertPayload struct {
	FirstName      string                   `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	LastName       string                   `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Phoneno        string                   `bson:"phoneno,omitempty" json:"phoneno,omitempty" validate:"required,e164"`
	Email          string                   `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Careerid       string                   `bson:"careerid,omitempty" json:"careerid,omitempty" validate:"required"`
	AdditionalInfo string                   `bson:"additionalInfo,omitempty" json:"additionalInfo,omitempty" validate:"required"`
	Country        string                   `bson:"country,omitempty" json:"country,omitempty" validate:"required"`
	Currentaddress string                   `bson:"currentaddress,omitempty" json:"currentaddress,omitempty" validate:"required"`
	Attachment     []map[string]interface{} `bson:"attachment,omitempty" json:"attachment,omitempty"`
}

type CareerUpsertPayload struct {
	ID               string             `bson:"_id,omitempty" json:"_id,omitempty"`
	Jobtitlename     string             `bson:"jobtitlename,omitempty" json:"jobtitlename,omitempty" validate:"required,min=3,max=30"`
	Experienced      []interface{}      `bson:"experienced,omitempty" json:"experienced,omitempty" validate:"required"`
	Jobcategory      string             `bson:"jobcategory,omitempty" json:"jobcategory,omitempty" validate:"required"`
	Location         string             `bson:"location,omitempty" json:"location,omitempty" validate:"required"`
	Jobsummary       string             `bson:"jobsummary,omitempty" json:"jobsummary,omitempty" validate:"required"`
	Responsibilities string             `bson:"responsibilites,omitempty" json:"responsibilites,omitempty" validate:"required"`
	MaximumApplyDate string             `bson:"maximumApplyDate,omitempty" json:"maximumApplyDate,omitempty" validate:"required"`
	RegisteredDate   primitive.DateTime `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	Jobtype          []interface{}      `bson:"jobtype,omitempty" json:"jobtype,omitempty" validate:"required"`
	Experiencelevel  []interface{}      `bson:"experiencelevel,omitempty" json:"experiencelevel,omitempty" validate:"required"`
	Requirement      string             `bson:"requirement,omitempty" json:"requirement,omitempty" validate:"required"`
}

type CareerDeletePayload struct {
	ID string `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
}
