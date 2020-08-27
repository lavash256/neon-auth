package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountEmptyVar(t *testing.T) {
	password := ""
	email := ""

	_, err := NewAccount(email, password)
	assert.Equal(t, err.Error(), "email: cannot be blank; password: cannot be blank.", "Incorrectly validates empty password and email values")
}

func TestAccountInvalidEmail(t *testing.T) {
	password := "Test"
	email := "Test"

	_, err := NewAccount(email, password)
	assert.Equal(t, err.Error(), "email: must be a valid email address.", "Incorrectly validates invalid email value")
}

func TestAccountEncryptPassword(t *testing.T) {
	password := "Test"
	email := "Test@test.ru"
	account, err := NewAccount(email, password)
	assert.Equal(t, err, nil, "If the password and email values ​​are correct, it should not give an error")
	assert.NotEqual(t, account.Password, password, "Password should not be stored in original format")
}
