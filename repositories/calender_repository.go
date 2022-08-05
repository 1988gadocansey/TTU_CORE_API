package repositories

import "TTU_CORE_ERP_API/models"

type CalenderRepository interface {
	Save(calender models.Calender) (uint, error)
	Update(bull models.Calender) error
	Delete(calender models.Calender) error
	FindAll() []*models.Calender
	FindByID(calender uint) (*models.Calender, error)
	DeleteByID(calender uint) error
	FindByName(name string) (*models.Calender, error)
	FindByField(fieldName, fieldValue string) (*models.Calender, error)
	UpdateSingleField(calender models.Calender, fieldName, fieldValue string) error
}
