package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

//Account is a method to identify a user
type Account struct {
	//An account is needed to access the application
	ID        uint64    `json:"account_id"`
	ProfileID uint64    `json:"profile_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Status    string    `json:"status"`
}

//EncryptsPassword ...
func (a *Account) EncryptsPassword() error {
	bytePassword := []byte(a.Password)
	//What is mean minCost?
	hashPwd, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashPwd)
	return nil
}

//ValidateEmailAndPassword checks email and password
func (a *Account) ValidateEmailAndPassword(email string, password string) error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)

}

//ComparePwd ...
func (a *Account) ComparePwd(password string) bool {
	byteHash := []byte(a.Password)
	bytePwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

//NewAccount is fabric for Account
func NewAccount(email string, password string) (*Account, error) {
	account := &Account{Email: email, Password: password}
	//Validate email
	err := account.ValidateEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	account.CreatedAt = time.Now()
	account.Status = "Create"
	err = account.EncryptsPassword()
	if err != nil {
		return nil, err
	}

	return account, nil
}

//BuildAccount Creates an Account object from the data
func BuildAccount(ID uint64, email string, password string, status string) *Account {
	return &Account{ID: ID, Email: email, Password: password, Status: status}
}
