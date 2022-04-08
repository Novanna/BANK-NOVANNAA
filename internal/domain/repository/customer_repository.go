package repository

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type ICustomerRepository interface {
	SaveCustomer(*entity.Customer) (*entity.Customer, error)
	GetAllCustomer() ([]entity.Customer, error)
	GetDetailCustomer(int) (*entity.Customer, error)
	UpdateCustomer(*entity.Customer) (*entity.Customer, error)
	DeleteCustomer(int) error
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	var customerRepo = CustomerRepository{}
	customerRepo.db = db
	return &customerRepo
}

func (r *CustomerRepository) SaveCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepository) GetAllCustomer() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Order("id asc").Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) GetDetailCustomer(no_ktp int) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("no_ktp = ?", no_ktp).Take(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) UpdateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepository) DeleteCustomer(no_ktp int) error {
	var customer entity.Customer
	err := r.db.Where("no_ktp = ?", no_ktp).Delete(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
