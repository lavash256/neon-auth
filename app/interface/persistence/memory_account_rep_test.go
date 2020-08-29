package persistence

import (
	"neon-auth/app/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountMemoryRepository(t *testing.T) {
	account, err := model.NewAccount("test@test.ru", "Test")
	assert.Equal(t, err, nil, "With the correct values ​​there should be no error")
	accountRepository := MemoryAccountRepository{}
	accountRepository.SaveAccount(account)
	accountInRepo, err := accountRepository.FindByEmail("test@test.ru")
	assert.Equal(t, account, accountInRepo, "Accounts must be equal")

	newAccount, err := model.NewAccount("new_test@test.ru", "Test")
	assert.Equal(t, err, nil, "With the correct values ​​there should be no error")
	assert.NotEqual(t, newAccount, accountInRepo, "Accounts must not be equal")

	newAccountInRepo, err := accountRepository.FindByEmail("testNot@test.ru")
	assert.Equal(t, err, nil, "With the correct values ​​there should be no error")
	assert.NotEqual(t, newAccountInRepo, nil, "If you selected an email that is not in the repository, nil should be returned")
}
