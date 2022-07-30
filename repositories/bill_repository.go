package repositories

import "TTU_CORE_ERP_API/models"

type BillRepository interface {
	Save(bill models.Department) (uint, error)
	Update(faculty models.Department) error
	Delete(faculty models.Department) error
	FindAll() []*models.Department
	FindByID(department uint) (*models.Department, error)
	DeleteByID(department uint) error
	FindByName(name string) (*models.Department, error)
	FindByField(fieldName, fieldValue string) (*models.Department, error)
	UpdateSingleField(department models.Department, fieldName, fieldValue string) error
}
