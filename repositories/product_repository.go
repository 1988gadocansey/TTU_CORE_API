package repositories

import (
	database "TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ProductRepository interface {
	Save(product models.PaymentProduct) (uint, error)
	Update(product models.PaymentProduct) error
	Delete(product models.PaymentProduct) error
	FindAll() []*models.PaymentProduct
	FindByCode(code string) (models.PaymentProduct, error)
}
type productDatabase struct {
	connection *gorm.DB
}

func (p productDatabase) Save(product models.PaymentProduct) (uint, error) {
	data := p.connection.Save(&product)
	if data.Error != nil {
		return 0, data.Error
	}
	return product.ID, nil
}

func (p productDatabase) Update(product models.PaymentProduct) error {
	data := p.connection.Save(&product)
	return data.Error
}

func (p productDatabase) Delete(product models.PaymentProduct) error {
	record := p.connection.Delete(&product)
	return record.Error
}

func (p productDatabase) FindAll() []*models.PaymentProduct {
	currentTime := time.Now().Local()

	var record []*models.PaymentProduct
	p.connection.
		Preload(clause.Associations).
		Where("Status = ?", true).
		Where("StartDate = ?", currentTime.String()).
		Where("EndDate > ?", currentTime.String()).
		Order("Name asc").
		Find(&record)

	return record
}

func (p productDatabase) FindByCode(code string) (models.PaymentProduct, error) {
	var product models.PaymentProduct
	result := p.connection.Preload(clause.Associations).Find(&product, "Code = ?", code)

	if result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func NewProductRepository() ProductRepository {

	db := database.Init()
	return &productDatabase{
		connection: db,
	}
}
