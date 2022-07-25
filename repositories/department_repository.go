package repositories

import (
	"TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DepartmentRepository interface {
	Save(faculty models.Department) (uint, error)
	Update(faculty models.Department) error
	Delete(faculty models.Department) error
	FindAll() []*models.Department
	FindByID(department uint) (*models.Department, error)
	DeleteByID(department uint) error
	FindByName(name string) (*models.Department, error)
	FindByField(fieldName, fieldValue string) (*models.Department, error)
	UpdateSingleField(department models.Department, fieldName, fieldValue string) error
}
type departmentDatabase struct {
	connection *gorm.DB
}

func NewDepartmentRepository() DepartmentRepository {
	db := database.Init()
	return &departmentDatabase{
		connection: db,
	}

}

func (db departmentDatabase) DeleteByID(departmentId uint) error {
	department := models.Department{}
	department.ID = departmentId
	result := db.connection.Delete(&department)
	return result.Error
}
func (db departmentDatabase) Save(department models.Department) (uint, error) {
	result := db.connection.Create(&department)
	if result.Error != nil {
		return 0, result.Error
	}
	return department.ID, nil
}
func (db departmentDatabase) Update(department models.Department) error {
	result := db.connection.Save(&department)
	return result.Error
}

func (db departmentDatabase) Delete(department models.Department) error {
	result := db.connection.Delete(&department)
	return result.Error
}

func (db departmentDatabase) FindAll() []*models.Department {
	var department []*models.Department
	db.connection.Preload(clause.Associations).Find(&department)
	return department
}

func (db departmentDatabase) FindByID(departmentId uint) (*models.Department, error) {
	var department models.Department
	result := db.connection.Preload(clause.Associations).Find(&department, "id = ?", departmentId)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &department, nil
	}
	return nil, nil
}

func (db departmentDatabase) FindByName(name string) (*models.Department, error) {
	var department models.Department
	result := db.connection.Preload(clause.Associations).Find(&department, "Name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &department, nil
	}
	return nil, nil
}

func (db departmentDatabase) FindByField(fieldName, fieldValue string) (*models.Department, error) {
	var department models.Department
	result := db.connection.Preload(clause.Associations).Find(&department, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &department, nil
	}
	return nil, nil
}

func (db departmentDatabase) UpdateSingleField(department models.Department, fieldName, fieldValue string) error {
	result := db.connection.Model(&department).Update(fieldName, fieldValue)
	return result.Error
}
