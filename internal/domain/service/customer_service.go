package service

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"
	"Trial/BANK-NOVANNA/internal/domain/repository"
)

type CustomerService struct {
	customerRepo repository.ICustomerRepository
}

type ICustomerService interface {
	SaveCustomer(*entity.CustomerViewModel) (*entity.CustomerViewModel, error)
	GetAllCustomer() ([]entity.CustomerViewModel, error)
	GetDetailCustomer(int) (*entity.CustomerViewModel, error)
	UpdateCustomer(*entity.CustomerViewModel) (*entity.CustomerViewModel, error)
	DeleteCustomer(int) error
}

func NewCustomerService(customerRepo repository.ICustomerRepository) *CustomerService {
	var customerService = CustomerService{}
	customerService.customerRepo = customerRepo
	return &customerService
}

func (s *CustomerService) SaveCustomer(customerVM *entity.CustomerViewModel) (*entity.CustomerViewModel, error) {

	var customer = entity.Customer{
		ID:           customerVM.ID,
		NamaLengkap:  customerVM.NamaLengkap,
		Alamat:       customerVM.Alamat,
		TempatLahir:  customerVM.TempatLahir,
		TanggalLahir: customerVM.TanggalLahir,
		JenisKelamin: customerVM.JenisKelamin,
		NoKTP:        customerVM.NoKTP,
		NoHP:         customerVM.NoHP,
	}

	result, err := s.customerRepo.SaveCustomer(&customer)

	if err != nil {
		return nil, err
	}

	if result != nil {
		customerVM = &entity.CustomerViewModel{
			ID:           result.ID,
			NamaLengkap:  result.NamaLengkap,
			Alamat:       result.Alamat,
			TempatLahir:  result.TempatLahir,
			TanggalLahir: result.TanggalLahir,
			JenisKelamin: result.JenisKelamin,
			NoKTP:        result.NoKTP,
			NoHP:         result.NoHP,
		}
	}

	return customerVM, nil
}

func (s *CustomerService) GetAllCustomer() ([]entity.CustomerViewModel, error) {
	result, err := s.customerRepo.GetAllCustomer()
	if err != nil {
		return nil, err
	}

	var customerListVM []entity.CustomerViewModel

	if result != nil {
		for _, item := range result {
			customerVM := entity.CustomerViewModel{
				ID:           item.ID,
				NamaLengkap:  item.NamaLengkap,
				Alamat:       item.Alamat,
				TempatLahir:  item.TempatLahir,
				TanggalLahir: item.TanggalLahir,
				JenisKelamin: item.JenisKelamin,
				NoKTP:        item.NoKTP,
				NoHP:         item.NoHP,
			}

			customerListVM = append(customerListVM, customerVM)
		}
	}

	return customerListVM, nil
}

func (s *CustomerService) GetDetailCustomer(no_ktp int) (*entity.CustomerViewModel, error) {
	result, err := s.customerRepo.GetDetailCustomer(no_ktp)

	if err != nil {
		return nil, err
	}

	var customerVM entity.CustomerViewModel

	if result != nil {
		customerVM = entity.CustomerViewModel{
			ID:           result.ID,
			NamaLengkap:  result.NamaLengkap,
			Alamat:       result.Alamat,
			TempatLahir:  result.TempatLahir,
			TanggalLahir: result.TanggalLahir,
			JenisKelamin: result.JenisKelamin,
			NoKTP:        result.NoKTP,
			NoHP:         result.NoHP,
		}
	}

	return &customerVM, nil
}
func (s *CustomerService) UpdateCustomer(customerVM *entity.CustomerViewModel) (*entity.CustomerViewModel, error) {

	var customer = entity.Customer{
		ID:           customerVM.ID,
		NamaLengkap:  customerVM.NamaLengkap,
		Alamat:       customerVM.Alamat,
		TempatLahir:  customerVM.TempatLahir,
		TanggalLahir: customerVM.TanggalLahir,
		JenisKelamin: customerVM.JenisKelamin,
		NoHP:         customerVM.NoHP,
		NoKTP:        customerVM.NoKTP,
	}

	_, err := s.customerRepo.UpdateCustomer(&customer)

	if err != nil {
		return nil, err
	}

	return customerVM, nil
}

func (s *CustomerService) DeleteCustomer(no_ktp int) error {
	err := s.customerRepo.DeleteCustomer(no_ktp)
	if err != nil {
		return err
	}

	return nil
}
