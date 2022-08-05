package repositories

import "TTU_CORE_ERP_API/models"

type PaymentRepository interface {
	Save(payment models.Payment) (uint, error)
	Update(payment models.Payment) error
	Delete(payment models.Payment) error
	FindAll() []*models.Payment
	FindByID(Payment uint) (*models.Payment, error)
	DeleteByID(payment uint) error
	FindByName(name string) (*models.Payment, error)
	FindByField(fieldName, fieldValue string) (*models.Payment, error)
	UpdateSingleField(Payment models.Payment, fieldName, fieldValue string) error
}
