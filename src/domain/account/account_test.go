package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
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
			email:    "Test",
			err:      "email: must be a valid email address.",
			message:  "Incorrectly validates invalid email value",
		},
		{
			password: "Test",
			email:    "test@yandex.ru",
			err:      "",
			message:  "If the data is valid, there should be no error",
		},
	}
	for _, testCase := range tests {

		_, err := NewAccount(testCase.email, testCase.password)
		if err != nil {
			assert.Equal(t, err.Error(), testCase.err, testCase.message)
			continue
		}
		assert.Equal(t, "", testCase.err, testCase.message)

	}
}

//Password check. That it is hashed right
func TestAccountPassword(t *testing.T) {
	password := "Test"
	email := "valid@valid.com"
	account, err := NewAccount(email, password)
	assert.Equal(t, err, nil, "If the data is valid, there should be no error")
	assert.NotEqual(t, account.Password, password, "The password must not be equal to the original")
	assert.Equal(t, true, account.ComparePwd(password), "Passwords in this case should be equal")
	assert.Equal(t, false, account.ComparePwd("NotTest"), "Hashes of different passwords should not be equal")

}
