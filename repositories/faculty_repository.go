package repositories

import (
	"TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FacultyRepository interface {
	Save(faculty models.Faculty) (uint, error)
	Update(level models.Faculty) error
	Delete(level models.Faculty) error
	FindAll() []*models.Faculty
	FindByID(LevelID uint) (*models.Faculty, error)
	DeleteByID(LevelID uint) error
	FindByName(name string) (*models.Faculty, error)
	FindByField(fieldName, fieldValue string) (*models.Faculty, error)
	UpdateSingleField(level models.Faculty, fieldName, fieldValue string) error
}
type facultyDatabase struct {
	connection *gorm.DB
}

func NewFacultyRepository() FacultyRepository {

	db := database.Init()
	return &facultyDatabase{
		connection: db,
	}

}

func (db facultyDatabase) DeleteByID(levelID uint) error {
	level := models.Level{}
	level.ID = levelID
	result := db.connection.Delete(&level)
	return result.Error
}
func (db facultyDatabase) Save(level models.Faculty) (uint, error) {
	result := db.connection.Create(&level)
	if result.Error != nil {
		return 0, result.Error
	}
	return level.ID, nil
}
func (db facultyDatabase) Update(level models.Faculty) error {
	result := db.connection.Save(&level)
	return result.Error
}

func (db facultyDatabase) Delete(level models.Faculty) error {
	result := db.connection.Delete(&level)
	return result.Error
}

func (db facultyDatabase) FindAll() []*models.Faculty {
	var faculties []*models.Faculty
	db.connection.Preload(clause.Associations).Find(&faculties)
	return faculties
}

func (db facultyDatabase) FindByID(facultyID uint) (*models.Faculty, error) {
	var faculty models.Faculty
	result := db.connection.Preload(clause.Associations).Find(&faculty, "id = ?", facultyID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &faculty, nil
	}
	return nil, nil
}

func (db facultyDatabase) FindByName(name string) (*models.Faculty, error) {
	var faculties models.Faculty
	result := db.connection.Preload(clause.Associations).Find(&faculties, "Name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &faculties, nil
	}
	return nil, nil
}

func (db facultyDatabase) FindByField(fieldName, fieldValue string) (*models.Faculty, error) {
	var faculties models.Faculty
	result := db.connection.Preload(clause.Associations).Find(&faculties, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &faculties, nil
	}
	return nil, nil
}

func (db facultyDatabase) UpdateSingleField(faculties models.Faculty, fieldName, fieldValue string) error {
	result := db.connection.Model(&faculties).Update(fieldName, fieldValue)
	return result.Error
}
