package usecase

import (
	"fmt"

	"github.com/lavash256/neon-auth/internal/domain/model"
	"github.com/lavash256/neon-auth/internal/domain/repository"
	"github.com/lavash256/neon-auth/internal/domain/service"
	"github.com/lavash256/neon-auth/internal/tools"
)

//AccountUsecaseInterface is Use cases are unit of the one operation for application.
type AccountUsecaseInterface interface {
	CreateAccount(email string, password string) error
}

//AccountUsecase is Interface implementation AccountUsecaseInterface
type AccountUsecase struct {
	logger      tools.LoggerInterface
	accountRepo repository.AccountRepository
	service     *service.AccountService
}

//NewAccountUsecase is fabric to create new account usecase
func NewAccountUsecase(repo repository.AccountRepository, logger tools.LoggerInterface) *AccountUsecase {
	accountServ := service.NewAccountService(repo)
	return &AccountUsecase{accountRepo: repo, service: accountServ, logger: logger}
}

// CreateAccount ...
func (auc *AccountUsecase) CreateAccount(email string, password string) error {
	err := auc.service.Dublicated(email)
	if err != nil {
		auc.logger.Error(fmt.Sprintf("Error %s in place check dublicate", err.Error()))
		return err
	}
	newAccount, err := model.NewAccount(email, password)
	if err != nil {
		auc.logger.Error(fmt.Sprintf("Error %s in place create new account", err.Error()))
		return err
	}
	err = auc.accountRepo.SaveAccount(newAccount)
	if err != nil {
		auc.logger.Error(fmt.Sprintf("Error %s in place SaveAccount", err.Error()))
		return err
	}
	return nil
}
