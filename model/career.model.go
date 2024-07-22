package model

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
