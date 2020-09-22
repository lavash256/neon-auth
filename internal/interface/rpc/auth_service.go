package rpc

import (
	"context"

	"github.com/lavash256/neon-auth/internal/usecase"
)

//AccountService Grpc service implementation
type AccountService struct {
	accountUsecase usecase.AccountUsecaseInterface
}

//NewAccountService is fabric to create accountService object
func NewAccountService(accountUsecase usecase.AccountUsecaseInterface) *AccountService {
	return &AccountService{
		accountUsecase: accountUsecase,
	}
}

// CreateAccount ...
func (a *AccountService) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	if err := a.accountUsecase.CreateAccount(req.GetEmail(), req.GetPassword()); err != nil {
		return &CreateAccountResponse{}, err
	}
	return &CreateAccountResponse{}, nil

}
