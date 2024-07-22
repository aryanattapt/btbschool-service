package model

type AlumniInsertPayload struct {
	FirstName          string                   `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	LastName           string                   `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Birthdate          string                   `bson:"birthdate,omitempty" json:"birthdate,omitempty"`
	Laststudentyear    string                   `bson:"laststudentyear,omitempty" json:"laststudentyear,omitempty" validate:"required"`
	EdukasiOptions     []string                 `bson:"edukasiOptions,omitempty" json:"edukasiOptions,omitempty" validate:"required"`
	StatusKerjaOptions string                   `bson:"statusKerjaOptions,omitempty" json:"statusKerjaOptions,omitempty" validate:"required"`
	Phoneno            string                   `bson:"phoneno,omitempty" json:"phoneno,omitempty" validate:"required,e164"`
	Email              string                   `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Attachment         []map[string]interface{} `bson:"attachment,omitempty" json:"attachment,omitempty"`
}
