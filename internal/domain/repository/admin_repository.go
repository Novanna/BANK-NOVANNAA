package repository

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

type IAdminRepository interface {
	SaveAdmin(*entity.Admin) (*entity.Admin, error)
	GetAdminByEmailPassword(loginVM entity.LoginViewModel) (*entity.Admin, error)
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	var adminRepo = AdminRepository{}
	adminRepo.db = db
	return &adminRepo
}

func (r *AdminRepository) SaveAdmin(admin *entity.Admin) (*entity.Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepository) GetAdminByEmailPassword(loginVM entity.LoginViewModel) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("email = ?", loginVM.Email).Take(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
