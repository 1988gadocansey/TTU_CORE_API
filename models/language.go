package models

import "gorm.io/gorm"

// User has and belongs to many languages, use `user_languages` as join table

type Language struct {
	Base
	gorm.Model
	Name     string
	Students []Student `gorm:"many2many:student_languages;"`
}

//db.Model(&language).Related(&users)
//// SELECT * FROM "users" INNER JOIN "student_languages" ON "student_languages"."student_id" = "students"."id" WHERE  ("student_languages"."language_id" IN ('111'))
