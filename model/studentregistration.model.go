package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentRegistrationInsertPayload struct {
	RegistrationCode            string             `bson:"registrationcode,omitempty" json:"registrationcode,omitempty"`
	Status                      string             `bson:"status,omitempty" json:"status,omitempty"`
	Schoolyear                  string             `bson:"schoolyear,omitempty" json:"schoolyear,omitempty" validate:"required"`
	Firstname                   string             `bson:"firstname,omitempty" json:"firstname,omitempty" validate:"required,min=3,max=30"`
	Middlename                  string             `bson:"middlename,omitempty" json:"middlename"`
	Lastname                    string             `bson:"lastname,omitempty" json:"lastname"`
	Birthplace                  string             `bson:"birthplace,omitempty" json:"birthplace" validate:"required"`
	Birthdate                   string             `bson:"birthdate,omitempty" json:"birthdate" validate:"required"`
	Religion                    string             `bson:"religion,omitempty" json:"religion,omitempty" validate:"required"`
	Gender                      string             `bson:"gender,omitempty" json:"gender,omitempty" validate:"required"`
	Address                     string             `bson:"address,omitempty" json:"address,omitempty" validate:"required"`
	Phoneno                     string             `bson:"phoneno,omitempty" json:"phoneno,omitempty" validate:"required"`
	Email                       string             `bson:"email,omitempty" json:"email,omitempty" validate:"required"`
	Musicinstrument             string             `bson:"musicinstrument,omitempty" json:"musicinstrument,omitempty" validate:"required"`
	Languagespoken              string             `bson:"languagespoken,omitempty" json:"languagespoken,omitempty" validate:"required"`
	Previousschoolname          string             `bson:"previousschoolname,omitempty" json:"previousschoolname,omitempty" validate:"required"`
	Yearlevelprevschool         string             `bson:"yearlevelprevschool,omitempty" json:"yearlevelprevschool,omitempty" validate:"required"`
	Nextclass                   string             `bson:"nextclass,omitempty" json:"nextclass,omitempty" validate:"required"`
	Fathername                  string             `bson:"fathername,omitempty" json:"fathername,omitempty" validate:"required"`
	Fatherbirthplace            string             `bson:"fatherbirthplace,omitempty" json:"fatherbirthplace,omitempty" validate:"required"`
	Fatherbirthdate             string             `bson:"fatherbirthdate,omitempty" json:"fatherbirthdate,omitempty" validate:"required"`
	Fatherphoneno               string             `bson:"fatherphoneno,omitempty" json:"fatherphoneno,omitempty" validate:"required"`
	Fatheremail                 string             `bson:"fatheremail,omitempty" json:"fatheremail,omitempty" validate:"required"`
	Fathermaritalstatus         string             `bson:"fathermaritalstatus,omitempty" json:"fathermaritalstatus,omitempty" validate:"required"`
	Mothername                  string             `bson:"mothername,omitempty" json:"mothername,omitempty" validate:"required"`
	Motherbirthplace            string             `bson:"motherbirthplace,omitempty" json:"motherbirthplace,omitempty" validate:"required"`
	Motherbirthdate             string             `bson:"motherbirthdate,omitempty" json:"motherbirthdate,omitempty" validate:"required"`
	Motherphoneno               string             `bson:"motherphoneno,omitempty" json:"motherphoneno,omitempty" validate:"required"`
	Motheremail                 string             `bson:"motheremail,omitempty" json:"motheremail,omitempty" validate:"required"`
	Mothermaritalstatus         string             `bson:"mothermaritalstatus,omitempty" json:"mothermaritalstatus,omitempty" validate:"required"`
	Emergencycontactname        string             `bson:"emergencycontactname,omitempty" json:"emergencycontactname,omitempty" validate:"required"`
	Emergencycontactrelaction   string             `bson:"emergencycontactrelaction,omitempty" json:"emergencycontactrelaction,omitempty" validate:"required"`
	Emergencycontactphoneno     string             `bson:"emergencycontactphoneno,omitempty" json:"emergencycontactphoneno,omitempty" validate:"required"`
	Emergencycontactaddress     string             `bson:"emergencycontactaddress,omitempty" json:"emergencycontactaddress,omitempty" validate:"required"`
	Siblinglist                 []interface{}      `bson:"siblinglist,omitempty" json:"siblinglist,omitempty" validate:"required"`
	Ttd                         string             `bson:"ttd,omitempty" json:"ttd,omitempty" validate:"required"`
	Bloodgroup                  string             `bson:"bloodgroup,omitempty" json:"bloodgroup,omitempty" validate:"required"`
	Doctorname                  string             `bson:"doctorname,omitempty" json:"doctorname,omitempty" validate:"required"`
	Doctorphone                 string             `bson:"doctorphone,omitempty" json:"doctorphone,omitempty" validate:"required"`
	Doctoraddress               string             `bson:"doctoraddress,omitempty" json:"doctoraddress,omitempty" validate:"required"`
	Medicationoption            string             `bson:"medicationoption,omitempty" json:"medicationoption,omitempty" validate:"required"`
	Isrecassmedicationoption    string             `bson:"isrecassmedicationoption,omitempty" json:"isrecassmedicationoption,omitempty" validate:"required"`
	Naturemedication            string             `bson:"naturemedication,omitempty" json:"naturemedication,omitempty" validate:"required"`
	Alergicoption               string             `bson:"alergicoption,omitempty" json:"alergicoption,omitempty" validate:"required"`
	Natureofallergy             string             `bson:"natureofallergy,omitempty" json:"natureofallergy,omitempty" validate:"required"`
	Limitationofphysical        string             `bson:"limitationofphysical,omitempty" json:"limitationofphysical,omitempty" validate:"required"`
	Limitationofphysicalexplain string             `bson:"limitationofphysicalexplain,omitempty" json:"limitationofphysicalexplain,omitempty"`
	Surgeryoperation            string             `bson:"surgeryoperation,omitempty" json:"surgeryoperation,omitempty" validate:"required"`
	Surgeryoperationexplain     string             `bson:"surgeryoperationexplain,omitempty" json:"surgeryoperationexplain,omitempty"`
	Medicalproblemoptions       []string           `bson:"medicalproblemoptions,omitempty" json:"medicalproblemoptions,omitempty" validate:"required"`
	Specificdisability          string             `bson:"specificdisability,omitempty" json:"specificdisability,omitempty" validate:"required"`
	Recommendedoption           string             `bson:"recommendedoption,omitempty" json:"recommendedoption,omitempty" validate:"required"`
	Btbparentnamerec            string             `bson:"btbparentnamerec,omitempty" json:"btbparentnamerec,omitempty" validate:"required"`
	Btbstudentnamerec           string             `bson:"btbstudentnamerec,omitempty" json:"btbstudentnamerec,omitempty" validate:"required"`
	Btbstudentgraderec          string             `bson:"btbstudentgraderec,omitempty" json:"btbstudentgraderec,omitempty" validate:"required"`
	Btbstudentphonehomerec      string             `bson:"btbstudentphonehomerec,omitempty" json:"btbstudentphonehomerec,omitempty" validate:"required"`
	Btbstudentphonemobilerec    string             `bson:"btbstudentphonemobilerec,omitempty" json:"btbstudentphonemobilerec,omitempty" validate:"required"`
	Attachment                  []interface{}      `bson:"attachment,omitempty" json:"attachment,omitempty" validate:"required"`
	RegisteredDate              primitive.DateTime `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	Updateddate                 primitive.DateTime `bson:"updateddate,omitempty" json:"updateddate,omitempty"`
	Nationality                 string             `bson:"nationality,omitempty" json:"nationality,omitempty" validate:"required"`
	Fatheroccupation            string             `bson:"fatheroccupation,omitempty" json:"fatheroccupation,omitempty" validate:"required"`
	Fathercompanyname           string             `bson:"fathercompanyname,omitempty" json:"fathercompanyname,omitempty" validate:"required"`
	FatherbusinessAddress       string             `bson:"fatherbusinessAddress,omitempty" json:"fatherbusinessAddress,omitempty" validate:"required"`
	Fathertelephone             string             `bson:"fathertelephone,omitempty" json:"fathertelephone,omitempty" validate:"required"`
	Fatherfax                   string             `bson:"fatherfax,omitempty" json:"fatherfax,omitempty" validate:"required"`
	Motheroccupation            string             `bson:"motheroccupation,omitempty" json:"motheroccupation,omitempty" validate:"required"`
	Mothercompanyname           string             `bson:"mothercompanyname,omitempty" json:"mothercompanyname,omitempty" validate:"required"`
	MotherbusinessAddress       string             `bson:"motherbusinessAddress,omitempty" json:"motherbusinessAddress,omitempty" validate:"required"`
	Mothertelephone             string             `bson:"mothertelephone,omitempty" json:"mothertelephone,omitempty" validate:"required"`
	Motherfax                   string             `bson:"motherfax,omitempty" json:"motherfax,omitempty" validate:"required"`
}

type DraftStudentRegistrationData struct {
	RegistrationCode string `bson:"registrationcode,omitempty" json:"registrationcode,omitempty"`
}
