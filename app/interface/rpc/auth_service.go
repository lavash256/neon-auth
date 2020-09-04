package rpc

import (
	"context"
	rpc "neon-auth/app/interface/rpc/protocol"
	"neon-auth/app/usecase"
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
func (a *AccountService) CreateAccount(ctx context.Context, req *rpc.CreateAccountRequest) (*rpc.CreateAccountResponse, error) {
	if err := a.accountUsecase.CreateAccount(req.GetEmail(), req.GetPassword()); err != nil {
		return &rpc.CreateAccountResponse{}, err
	}
	return &rpc.CreateAccountResponse{}, nil

}
