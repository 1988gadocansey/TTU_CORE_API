package repositories

import (
	"TTU_CORE_PORTAL_GO/configs"
	"TTU_CORE_PORTAL_GO/models"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentRepository interface {
	Save(student models.Student) (uint, error)
	Update(student models.Student) error
	Delete(student models.Student) error
	FindAll() []*models.Student
	FindByID(StudentID uint) (*models.Student, error)
	FindByUUID(UUID uuid.UUID) (*models.Student, error)
	DeleteByID(StudentID uint) error
	FindByName(name string) (*models.Student, error)
	FindByField(fieldName, fieldValue string) (*models.Student, error)
	UpdateSingleField(student models.Student, fieldName, fieldValue string) error
}
type studentDatabase struct {
	connection *gorm.DB
}

func NewStudentRepository() StudentRepository {

	db := database.Init()
	return &studentDatabase{
		connection: db,
	}
}
func (db studentDatabase) DeleteByID(studentID uint) error {
	student := models.Student{}
	student.ID = studentID
	result := db.connection.Delete(&student)
	return result.Error
}
func (db studentDatabase) Save(student models.Student) (uint, error) {
	result := db.connection.Create(&student)
	if result.Error != nil {
		return 0, result.Error
	}
	return student.ID, nil
}
func (db studentDatabase) Update(student models.Student) error {
	result := db.connection.Save(&student)
	return result.Error
}

func (db studentDatabase) Delete(student models.Student) error {
	result := db.connection.Delete(&student)
	return result.Error
}

func (db studentDatabase) FindAll() []*models.Student {
	var student []*models.Student
	db.connection.Preload(clause.Associations).Find(&student)
	return student
}

func (db studentDatabase) FindByID(studentID uint) (*models.Student, error) {
	var student models.Student
	result := db.connection.Preload(clause.Associations).Find(&student, "id = ?", studentID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &student, nil
	}
	return nil, nil
}

func (db studentDatabase) FindByUUID(UUID uuid.UUID) (*models.Student, error) {
	var student models.Student
	result := db.connection.Preload(clause.Associations).Find(&student, "uuid = ?", UUID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &student, nil
	}
	return nil, nil
}

func (db studentDatabase) FindByName(name string) (*models.Student, error) {
	var student models.Student
	result := db.connection.Preload(clause.Associations).Find(&student, "Name LIKE  ?", "%"+name+"%")

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &student, nil
	}
	return nil, nil
}

func (db studentDatabase) FindByField(fieldName, fieldValue string) (*models.Student, error) {
	var student models.Student
	result := db.connection.Preload(clause.Associations).Find(&student, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &student, nil
	}
	return nil, nil
}

func (db studentDatabase) UpdateSingleField(student models.Student, fieldName, fieldValue string) error {
	result := db.connection.Model(&student).Update(fieldName, fieldValue)
	return result.Error
}
