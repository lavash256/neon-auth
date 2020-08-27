package account

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	//An account is needed to access the application
	ID        uint64    `json:"account_id"`
	ProfileID uint64    `json:"profile_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//When initially creating an account, you should not know about the profile
func NewAccount(email string, password string) (*Account, error) {
	account := &Account{Email: email, Password: password}
	//Validate email
	err := account.ValidateEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	account.CreatedAt = time.Now()
	err = account.EncryptsPassword()
	if err != nil {
		return nil, err
	}

	return account, nil
}

//Encrypts the password and assigns it to the object
func (a *Account) EncryptsPassword() error {
	bytePassword := []byte(a.Password)
	//What is mean minCost?
	hashPwd, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	a.Password = string(hashPwd)
	return nil
}

func (a *Account) ValidateEmailAndPassword(email string, password string) error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)

}
