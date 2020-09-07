package repository

import "neon-auth/src/app/domain/model"

//AccountRepository it is an interface for operations with account storage
type AccountRepository interface {
	SaveAccount(*model.Account) error
	FindByEmail(string) (*model.Account, error)
}
