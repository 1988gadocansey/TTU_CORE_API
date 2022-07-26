package repositories

import (
	"TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProgrammeRepository interface {
	Save(programme models.Programme) (uint, error)
	Update(programme models.Programme) error
	Delete(programme models.Programme) error
	FindAll() []*models.Programme
	FindByID(programme uint) (*models.Programme, error)
	DeleteByID(programme uint) error
	FindByName(name string) (*models.Programme, error)
	FindByField(fieldName, fieldValue string) (*models.Programme, error)
	UpdateSingleField(programme models.Programme, fieldName, fieldValue string) error
}
type programmeDatabase struct {
	connection *gorm.DB
}

func NewProgrammeRepository() ProgrammeRepository {

	db := database.Init()
	return &programmeDatabase{
		connection: db,
	}

}

func (db programmeDatabase) DeleteByID(programmeId uint) error {
	programme := models.Programme{}
	programme.ID = programmeId
	result := db.connection.Delete(&programme)
	return result.Error
}
func (db programmeDatabase) Save(programme models.Programme) (uint, error) {
	result := db.connection.Create(&programme)
	if result.Error != nil {
		return 0, result.Error
	}
	return programme.ID, nil
}
func (db programmeDatabase) Update(programme models.Programme) error {
	result := db.connection.Save(&programme)
	return result.Error
}

func (db programmeDatabase) Delete(programme models.Programme) error {
	result := db.connection.Delete(&programme)
	return result.Error
}

func (db programmeDatabase) FindAll() []*models.Programme {
	var programme []*models.Programme
	db.connection.Preload(clause.Associations).Find(&programme)
	return programme
}

func (db programmeDatabase) FindByID(programmeId uint) (*models.Programme, error) {
	var programme models.Programme
	result := db.connection.Preload(clause.Associations).Find(&programme, "id = ?", programmeId)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &programme, nil
	}
	return nil, nil
}

func (db programmeDatabase) FindByName(name string) (*models.Programme, error) {
	var programme models.Programme
	result := db.connection.Preload(clause.Associations).Find(&programme, "Name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &programme, nil
	}
	return nil, nil
}

func (db programmeDatabase) FindByField(fieldName, fieldValue string) (*models.Programme, error) {
	var programme models.Programme
	result := db.connection.Preload(clause.Associations).Find(&programme, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &programme, nil
	}
	return nil, nil
}

func (db programmeDatabase) UpdateSingleField(programme models.Programme, fieldName, fieldValue string) error {
	result := db.connection.Model(&programme).Update(fieldName, fieldValue)
	return result.Error
}
