package repositories

import (
	database "TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Pay(payment models.Payment) (uint, error)
	Update(payment models.Payment) error
	Delete(payment models.Payment) error
	FindAll() []*models.Payment
	FindByID(Payment uint) (*models.Payment, error)
	FindByTransactionId(transactionId string) (*models.Payment, error)
	DeleteByTransactionId(transactionId string) error
	DeleteByID(payment uint) error
	FindByField(fieldName, fieldValue string) (*models.Payment, error)
	UpdateSingleField(Payment models.Payment, fieldName, fieldValue string) error
}
type paymentDatabase struct {
	connection *gorm.DB
}

// Pay accept payment from external entity say bank, ussd, local api into db /**
func (p paymentDatabase) Pay(payment models.Payment) (uint, error) {

}

func (p paymentDatabase) Update(payment models.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) Delete(payment models.Payment) error {
	record := p.connection.Delete(payment)

	return record.Error
}

func (p paymentDatabase) FindAll() []*models.Payment {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) FindByID(Payment uint) (*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) FindByTransactionId(transactionId string) (*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) DeleteByTransactionId(transactionId string) error {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) DeleteByID(payment uint) error {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) FindByField(fieldName, fieldValue string) (*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentDatabase) UpdateSingleField(Payment models.Payment, fieldName, fieldValue string) error {
	//TODO implement me
	panic("implement me")
}

func NewPaymentRepository() PaymentRepository {

	db := database.Init()
	return &paymentDatabase{
		connection: db,
	}
}
