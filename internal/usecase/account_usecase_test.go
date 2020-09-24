package usecase

import (
	"testing"

	"github.com/lavash256/neon-auth/internal/interface/persistence"
	"github.com/lavash256/neon-auth/internal/tools"

	"github.com/stretchr/testify/assert"
)

func TestAccountUseCaseCreate(t *testing.T) {
	accountRepo := persistence.MemoryAccountRepository{}
	logger := &tools.LoggerStub{}
	accountUsecase := NewAccountUsecase(&accountRepo, logger)
	tests := []struct {
		password string
		email    string
		err      string
		message  string
	}{
		{
			password: "",
			email:    "",
			err:      "email: cannot be blank; password: cannot be blank.",
			message:  "Incorrectly validates empty password and email values",
		},
		{
			password: "Test",
			email:    "Email",
			err:      "email: must be a valid email address.",
			message:  "Incorrectly validates invalid email values",
		},
		{
			password: "Test",
			email:    "test@test.ru",
			err:      "",
			message:  "Valid data. Should persist",
		},
		{
			password: "test",
			email:    "test@test.ru",
			err:      "test@test.ru already exists",
			message:  "Should not save with the same email",
		},
	}
	for _, testCase := range tests {
		err := accountUsecase.CreateAccount(testCase.email, testCase.password)
		if err != nil {
			assert.Equal(t, testCase.err, err.Error(), testCase.message)
			continue
		}
		assert.Equal(t, testCase.err, "", testCase.message)
	}

}
