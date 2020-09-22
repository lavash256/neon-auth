package service

import (
	"testing"

	"github.com/lavash256/neon-auth/internal/domain/model"
	"github.com/lavash256/neon-auth/internal/interface/persistence"

	"github.com/stretchr/testify/assert"
)

func TestAccountService(t *testing.T) {
	accountRepository := persistence.MemoryAccountRepository{}
	account, err := model.NewAccount("test@test.ru", "Test")
	_ = accountRepository.SaveAccount(account)
	accountService := AccountService{repository: &accountRepository}
	err = accountService.Dublicated("new_test@test.ru")
	assert.Equal(t, nil, err, "Since this account is not in the repository, it must return nil")

	err = accountService.Dublicated("test@test.ru")
	assert.Equal(t, "test@test.ru already exists", err.Error(), "Since a user with such an account exists, he should return an error")
}
