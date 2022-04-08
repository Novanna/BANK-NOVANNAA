package service

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"
	"Trial/BANK-NOVANNA/internal/domain/repository"
	"Trial/BANK-NOVANNA/pkg/security"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	adminRepo repository.IAdminRepository
}

type IAdminService interface {
	SaveAdmin(*entity.ReqisterViewModel) (*entity.AdminViewModel, error)
	GetAdminByEmailPassword(loginVM entity.LoginViewModel) (*entity.Admin, error)
}

func NewAdminService(adminRepo repository.IAdminRepository) *AdminService {
	var adminService = AdminService{}
	adminService.adminRepo = adminRepo
	return &adminService
}

func (s *AdminService) SaveAdmin(adminVM *entity.ReqisterViewModel) (*entity.AdminViewModel, error) {
	var admin = entity.Admin{
		FirstName: adminVM.FirstName,
		LastName:  adminVM.LastName,
		Email:     adminVM.Email,
	}

	password, err := admin.EncryptPassword(adminVM.Password)
	if err != nil {
		return nil, err
	}

	admin.Password = password

	result, err := s.adminRepo.SaveAdmin(&admin)
	if err != nil {
		return nil, err
	}

	var afterRegVM entity.AdminViewModel

	if result != nil {
		afterRegVM = entity.AdminViewModel{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}

	return &afterRegVM, nil
}

func (s *AdminService) GetAdminByEmailPassword(loginVM entity.LoginViewModel) (*entity.Admin, error) {
	result, err := s.adminRepo.GetAdminByEmailPassword(loginVM)
	if err != nil {
		return nil, err
	}

	// Verify Password
	err = security.VerifyPassword(result.Password, loginVM.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("Incorrect Password. Error %s", err.Error())
	}

	return result, nil
}
