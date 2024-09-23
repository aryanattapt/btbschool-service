package model

import "go.mongodb.org/mongo-driver/bson/primitive"

/* Generic Data */
type StudentRegistrationGenericPayload struct {
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	RegisteredDate primitive.DateTime `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	Updateddate    primitive.DateTime `bson:"updateddate,omitempty" json:"updateddate,omitempty"`
}

/* Page 0 */
type StudentRegistrationInsertPayloadPage0 struct {
	MainEmail        string `bson:"mainEmail,omitempty" json:"mainEmail,omitempty" validate:"required"`
	RegistrationCode string `bson:"registrationcode,omitempty" json:"registrationcode,omitempty"`
}

/* Page 1 */
type schoolInformationForm struct {
	Schoolyear string `bson:"schoolyear,omitempty" json:"schoolyear,omitempty" validate:"required"`
}

type studentDetailForm struct {
	Firstname      string        `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	Middlename     string        `bson:"middlename,omitempty" json:"middlename"`
	Lastname       string        `bson:"lastname,omitempty" json:"lastname"`
	Birthplace     string        `bson:"birthplace,omitempty" json:"birthplace" validate:"required"`
	Birthdate      string        `bson:"birthdate,omitempty" json:"birthdate" validate:"required"`
	Nationality    string        `bson:"nationality,omitempty" json:"nationality,omitempty" validate:"required"`
	Religion       string        `bson:"religion,omitempty" json:"religion,omitempty" validate:"required"`
	Gender         string        `bson:"gender,omitempty" json:"gender,omitempty" validate:"required"`
	Address        string        `bson:"address,omitempty" json:"address,omitempty" validate:"required"`
	Phoneno        string        `bson:"phoneno,omitempty" json:"phoneno,omitempty" validate:"required,e164"`
	Email          string        `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Languagespoken []interface{} `bson:"languagespoken,omitempty" json:"languagespoken,omitempty" validate:"required"`
}
type educationalBackgroundForm struct {
	Previousschoolname  string `bson:"previousschoolname,omitempty" json:"previousschoolname,omitempty" validate:"required"`
	Yearlevelprevschool string `bson:"yearlevelprevschool,omitempty" json:"yearlevelprevschool,omitempty" validate:"required"`
	Nextclass           string `bson:"nextclass,omitempty" json:"nextclass,omitempty" validate:"required"`
}

type parentsInformationForm struct {
	Fathername            string `bson:"fathername,omitempty" json:"fathername,omitempty" validate:"required"`
	Fatherbirthplace      string `bson:"fatherbirthplace,omitempty" json:"fatherbirthplace,omitempty" validate:"required"`
	Fatherbirthdate       string `bson:"fatherbirthdate,omitempty" json:"fatherbirthdate,omitempty" validate:"required"`
	Fatherphoneno         string `bson:"fatherphoneno,omitempty" json:"fatherphoneno,omitempty" validate:"required,e164"`
	Fatheremail           string `bson:"fatheremail,omitempty" json:"fatheremail,omitempty" validate:"required,email"`
	Fathermaritalstatus   string `bson:"fathermaritalstatus,omitempty" json:"fathermaritalstatus,omitempty" validate:"required"`
	Fatheroccupation      string `bson:"fatheroccupation,omitempty" json:"fatheroccupation,omitempty" validate:"required"`
	Fathercompanyname     string `bson:"fathercompanyname,omitempty" json:"fathercompanyname,omitempty"`
	FatherbusinessAddress string `bson:"fatherbusinessAddress,omitempty" json:"fatherbusinessAddress,omitempty"`
	Fathertelephone       string `bson:"fathertelephone,omitempty" json:"fathertelephone,omitempty"`
	Fatherfax             string `bson:"fatherfax,omitempty" json:"fatherfax,omitempty"`
	Mothername            string `bson:"mothername,omitempty" json:"mothername,omitempty" validate:"required"`
	Motherbirthplace      string `bson:"motherbirthplace,omitempty" json:"motherbirthplace,omitempty" validate:"required"`
	Motherbirthdate       string `bson:"motherbirthdate,omitempty" json:"motherbirthdate,omitempty" validate:"required"`
	Motherphoneno         string `bson:"motherphoneno,omitempty" json:"motherphoneno,omitempty" validate:"required,e164"`
	Motheremail           string `bson:"motheremail,omitempty" json:"motheremail,omitempty" validate:"required,email"`
	Mothermaritalstatus   string `bson:"mothermaritalstatus,omitempty" json:"mothermaritalstatus,omitempty" validate:"required"`
	Motheroccupation      string `bson:"motheroccupation,omitempty" json:"motheroccupation,omitempty" validate:"required"`
	Mothercompanyname     string `bson:"mothercompanyname,omitempty" json:"mothercompanyname,omitempty"`
	MotherbusinessAddress string `bson:"motherbusinessAddress,omitempty" json:"motherbusinessAddress,omitempty"`
	Mothertelephone       string `bson:"mothertelephone,omitempty" json:"mothertelephone,omitempty"`
	Motherfax             string `bson:"motherfax,omitempty" json:"motherfax,omitempty"`
}

type emergencyContactForm struct {
	Emergencycontactname      string `bson:"emergencycontactname,omitempty" json:"emergencycontactname,omitempty" validate:"required"`
	Emergencycontactrelaction string `bson:"emergencycontactrelaction,omitempty" json:"emergencycontactrelaction,omitempty" validate:"required"`
	Emergencycontactphoneno   string `bson:"emergencycontactphoneno,omitempty" json:"emergencycontactphoneno,omitempty" validate:"required,e164"`
	Emergencycontactaddress   string `bson:"emergencycontactaddress,omitempty" json:"emergencycontactaddress,omitempty" validate:"required"`
}

type StudentRegistrationInsertPayloadPage1 struct {
	schoolInformationForm
	studentDetailForm
	educationalBackgroundForm
	parentsInformationForm
	emergencyContactForm
	Siblinglist []interface{} `bson:"siblinglist,omitempty" json:"siblinglist,omitempty" validate:"required"`
	Ttdpage1    string        `bson:"ttdpage1,omitempty" json:"ttdpage1,omitempty" validate:"required"`
}

/* Page 2 */
type StudentRegistrationInsertPayloadPage2 struct {
	Ttdpage2 string `bson:"ttdpage2,omitempty" json:"ttdpage2,omitempty" validate:"required"`
}

/* Page 3 */
type personalHealthInformationForm struct {
	Bloodgroup               string `bson:"bloodgroup,omitempty" json:"bloodgroup,omitempty"`
	Doctorname               string `bson:"doctorname,omitempty" json:"doctorname,omitempty"`
	Doctorphone              string `bson:"doctorphone,omitempty" json:"doctorphone,omitempty"`
	Doctoraddress            string `bson:"doctoraddress,omitempty" json:"doctoraddress,omitempty"`
	Medicationoption         string `bson:"medicationoption,omitempty" json:"medicationoption,omitempty" validate:"required"`
	Isrecassmedicationoption string `bson:"isrecassmedicationoption,omitempty" json:"isrecassmedicationoption,omitempty" validate:"required"`
	Naturemedication         string `bson:"naturemedication,omitempty" json:"naturemedication,omitempty"`
}

type medicalProblemForm struct {
	Alergicoption               string   `bson:"alergicoption,omitempty" json:"alergicoption,omitempty"`
	Natureofallergy             string   `bson:"natureofallergy,omitempty" json:"natureofallergy,omitempty"`
	Limitationofphysical        string   `bson:"limitationofphysical,omitempty" json:"limitationofphysical,omitempty" validate:"required"`
	Limitationofphysicalexplain string   `bson:"limitationofphysicalexplain,omitempty" json:"limitationofphysicalexplain,omitempty"`
	Surgeryoperation            string   `bson:"surgeryoperation,omitempty" json:"surgeryoperation,omitempty" validate:"required"`
	Surgeryoperationexplain     string   `bson:"surgeryoperationexplain,omitempty" json:"surgeryoperationexplain,omitempty"`
	Medicalproblemoptions       []string `bson:"medicalproblemoptions,omitempty" json:"medicalproblemoptions,omitempty" validate:"required"`
	Specificdisability          string   `bson:"specificdisability,omitempty" json:"specificdisability,omitempty" validate:"required"`
}

type StudentRegistrationInsertPayloadPage3 struct {
	personalHealthInformationForm
	medicalProblemForm
	Ttdpage3 string `bson:"ttdpage3,omitempty" json:"ttdpage3,omitempty" validate:"required"`
}

/* Page 4 */
type recomendedForm struct {
	Recommendedoption        string `bson:"recommendedoption,omitempty" json:"recommendedoption,omitempty" validate:"required"`
	Btbparentnamerec         string `bson:"btbparentnamerec,omitempty" json:"btbparentnamerec,omitempty"`
	Btbstudentnamerec        string `bson:"btbstudentnamerec,omitempty" json:"btbstudentnamerec,omitempty"`
	Btbstudentgraderec       string `bson:"btbstudentgraderec,omitempty" json:"btbstudentgraderec,omitempty"`
	Btbstudentphonehomerec   string `bson:"btbstudentphonehomerec,omitempty" json:"btbstudentphonehomerec,omitempty"`
	Btbstudentphonemobilerec string `bson:"btbstudentphonemobilerec,omitempty" json:"btbstudentphonemobilerec,omitempty"`
}

type StudentRegistrationInsertPayloadPage4 struct {
	recomendedForm
	Attachment []interface{} `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Ttdpage4   string        `bson:"ttdpage4,omitempty" json:"ttdpage4,omitempty" validate:"required"`
}

/* Main Data */
type StudentRegistrationInsertPayload struct {
	StudentRegistrationGenericPayload
	StudentRegistrationInsertPayloadPage0
	StudentRegistrationInsertPayloadPage1
	StudentRegistrationInsertPayloadPage3
	StudentRegistrationInsertPayloadPage4
}

type DraftStudentRegistrationData struct {
	RegistrationCode string `bson:"registrationcode,omitempty" json:"registrationcode,omitempty"`
}
