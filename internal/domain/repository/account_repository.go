package repository

import "github.com/lavash256/neon-auth/internal/domain/model"

//AccountRepository it is an interface for operations with account storage
type AccountRepository interface {
	SaveAccount(*model.Account) error
	FindByEmail(string) (*model.Account, error)
}
