package service

import (
	"fmt"
	"neon-auth/src/app/domain/repository"
)

//AccountService contains business rules that apply to Account
type AccountService struct {
	repository repository.AccountRepository
}

//NewAccountService is fabric
func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repository: repo}
}

//Dublicated Checks the existence of an account with this email
func (service *AccountService) Dublicated(email string) error {
	account, err := service.repository.FindByEmail(email)
	if account != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}
