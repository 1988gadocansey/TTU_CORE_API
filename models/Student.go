package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	Base
	gorm.Model
	Stno                 string          `db:"Stno" json:"Stno" gorm:"uniqueIndex" validate:"required,Stno"`
	Indexno              string          `db:"Indexno" json:"Indexno" gorm:"uniqueIndex" validate:"required,Indexno"`
	CustomId             string          `db:"CustomId" json:"CustomId" gorm:"uniqueIndex"`
	LevelID              uint8           ` json:"LevelID" validate:"required,LastName"`
	Level                Level           ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Level"`
	Title                AppTitle        `db:"Title" json:"Title"  validate:"required"`
	FirstName            string          `db:"FirstName" json:"FirstName"  validate:"required,FirstName"`
	LastName             string          `db:"LastName" json:"LastName"  validate:"required,LastName"`
	OtherNames           string          `db:"OtherName" json:"OtherName"`
	PreviousName         string          `db:"PreviousName" json:"PreviousName"`
	Gender               GenderData      `db:"Gender" json:"Gender"  validate:"required"`
	DateOfBirth          time.Time       `db:"DateOfBirth" json:"DateOfBirth"  validate:"required"`
	MaritalStatus        MaritalStatuses `db:"MaritalStatus" json:"MaritalStatus"  validate:"required"`
	PhoneNumber          string          `db:"PhoneNumber" json:"PhoneNumber"  validate:"required"`
	AltPhoneNumber       string          `db:"AltPhoneNumber" json:"AltPhoneNumber" `
	Hometown             string          `db:"Hometown" json:"Hometown"  validate:"required"`
	DistrictID           uint            ` json:"DistrictID"`
	District             District        ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"District"`
	NationalityID        uint8
	Nationality          Nationality ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Nationality"`
	DenominationID       uint8
	Denomination         Denomination ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Denomination"`
	FormerSchoolID       uint8
	FormerSchool         FormerSchool ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"FormerSchool"`
	Emails               []Email      `json:"emails"`
	Section              Sections     `db:"Section" json:"Section"  validate:"required"`
	GraduationGroup      string       `db:"GraduationGroup" json:"GraduationGroup"  validate:"required"`
	GuardianName         string       `db:"GuardianName" json:"GuardianName"  validate:"required"`
	GuardianPhone        string       `db:"GuardianPhone" json:"GuardianPhone"  validate:"required"`
	GuardianAddress      string       `db:"GuardianAddress" json:"GuardianAddress"  validate:"required"`
	GuardianOccupation   string       `db:"GuardianOccupation" json:"GuardianOccupation"  validate:"required"`
	GuardianLocation     string       `db:"GuardianLocation" json:"GuardianLocation"  validate:"required"`
	GuardianRelationship string       `db:"GuardianRelationship" json:"GuardianRelationship"  validate:"required"`
	DateAdmitted         time.Time    `db:"DateAdmitted" json:"DateAdmitted"  validate:"required"`
	LevelAdmittedID      uint8
	LevelAdmitted        Level   ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"LevelAdmitted"`
	CGPA                 float32 `db:"CGPA" json:"CGPA"`
	ClassingID           uint8
	Classing             Classing ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"ClassObtained"`
	Matriculated         bool     `db:"Matriculated" json:"Matriculated"`
	QualityAssurance     bool     `db:"QualityAssurance" json:"QualityAssurance"`
	RoomNumber           string   `db:"RoomNumber" json:"RoomNumber"`
	Hostel               string   `db:"Hostel" json:"Hostel"`
	GhanaCard            string   `db:"GhanaCard" json:"GhanaCard"`
	HallID               uint8
	Hall                 Hall             ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Hall"`
	ProgrammeID          uint8            ` json:"ProgrammeID"`
	Languages            []Language       `gorm:"many2many:student_languages;"` //student has and belongs to many languages, use `student_languages` as join table
	Programme            Programme        ` gorm:"many2many,constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Programme"`
	Address              []Address        `gorm:"many2many,ForeignKey:StudentID" json:"Address" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
	Results              []AcademicRecord `gorm:"many2many,ForeignKey:StudentID" json:"Results" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
	Disabilities         []Disability     `gorm:"many2many,ForeignKey:StudentID" json:"Disabilities" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
	Payments             []Payment        `gorm:"many2many,ForeignKey:StudentID" json:"Payments" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}
