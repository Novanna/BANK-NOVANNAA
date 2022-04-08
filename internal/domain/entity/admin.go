package entity

import (
	"Trial/BANK-NOVANNA/pkg/security"

	"github.com/badoux/checkmail"
)

type Admin struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ReqisterViewModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginViewModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminViewModel struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (a *Admin) EncryptPassword(password string) (string, error) {
	hashPassword, err := security.Hash(password)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func (a *Admin) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	if a.Email == "" {
		errorMessages["email_required"] = "email required"
	}
	if a.Email != "" {
		if err = checkmail.ValidateFormat(a.Email); err != nil {
			errorMessages["invalid_email"] = "email email"
		}
	}

	return errorMessages
}

func (a *LoginViewModel) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	if a.Password == "" {
		errorMessages["password_required"] = "password is required"
	}
	if a.Email == "" {
		errorMessages["email_required"] = "email is required"
	}
	if a.Email != "" {
		if err = checkmail.ValidateFormat(a.Email); err != nil {
			errorMessages["invalid_email"] = "please provide a valid email"
		}
	}

	return errorMessages
}

func (a *ReqisterViewModel) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	if a.FirstName == "" {
		errorMessages["firstname_required"] = "first name is required"
	}
	if a.LastName == "" {
		errorMessages["lastname_required"] = "last name is required"
	}
	if a.Password == "" {
		errorMessages["password_required"] = "password is required"
	}
	if a.Password != "" && len(a.Password) < 6 {
		errorMessages["invalid_password"] = "password should be at least 6 characters"
	}
	if a.Email == "" {
		errorMessages["email_required"] = "email is required"
	}
	if a.Email != "" {
		if err = checkmail.ValidateFormat(a.Email); err != nil {
			errorMessages["invalid_email"] = "please provide a valid email"
		}
	}

	return errorMessages
}