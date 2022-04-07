package repositories

import (
	"TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LevelRepository interface {
	Save(company models.Level) (uint, error)
	Update(level models.Level) error
	Delete(level models.Level) error
	FindAll() []*models.Level
	FindByID(LevelID uint) (*models.Level, error)
	DeleteByID(LevelID uint) error
	FindByName(name string) (*models.Level, error)
	FindByField(fieldName, fieldValue string) (*models.Level, error)
	UpdateSingleField(level models.Level, fieldName, fieldValue string) error
}
type levelDatabase struct {
	connection *gorm.DB
}

func NewLevelRepository() LevelRepository {

	db := database.Init()
	return &levelDatabase{
		connection: db,
	}
}
func (db levelDatabase) DeleteByID(levelID uint) error {
	level := models.Level{}
	level.ID = levelID
	result := db.connection.Delete(&level)
	return result.Error
}
func (db levelDatabase) Save(level models.Level) (uint, error) {
	result := db.connection.Create(&level)
	if result.Error != nil {
		return 0, result.Error
	}
	return level.ID, nil
}
func (db levelDatabase) Update(level models.Level) error {
	result := db.connection.Save(&level)
	return result.Error
}

func (db levelDatabase) Delete(level models.Level) error {
	result := db.connection.Delete(&level)
	return result.Error
}

func (db levelDatabase) FindAll() []*models.Level {
	var levels []*models.Level
	db.connection.Preload(clause.Associations).Find(&levels)
	return levels
}

func (db levelDatabase) FindByID(levelID uint) (*models.Level, error) {
	var level models.Level
	result := db.connection.Preload(clause.Associations).Find(&level, "id = ?", levelID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &level, nil
	}
	return nil, nil
}

func (db levelDatabase) FindByName(name string) (*models.Level, error) {
	var level models.Level
	result := db.connection.Preload(clause.Associations).Find(&level, "Name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &level, nil
	}
	return nil, nil
}

func (db levelDatabase) FindByField(fieldName, fieldValue string) (*models.Level, error) {
	var level models.Level
	result := db.connection.Preload(clause.Associations).Find(&level, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &level, nil
	}
	return nil, nil
}

func (db levelDatabase) UpdateSingleField(level models.Level, fieldName, fieldValue string) error {
	result := db.connection.Model(&level).Update(fieldName, fieldValue)
	return result.Error
}
