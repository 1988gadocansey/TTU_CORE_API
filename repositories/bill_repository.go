package repositories

import "TTU_CORE_ERP_API/models"

type BillRepository interface {
	Save(bill models.Bill) (uint, error)
	Update(bill models.Bill) error
	Delete(bill models.Bill) error
	FindAll() []*models.Bill
	FindByID(Bill uint) (*models.Bill, error)
	DeleteByID(Bill uint) error
	FindByName(name string) (*models.Bill, error)
	FindByField(fieldName, fieldValue string) (*models.Bill, error)
	UpdateSingleField(bill models.Bill, fieldName, fieldValue string) error
}
