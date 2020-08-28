package account

//Repository it is an interface for operations with account storage
type Repository interface {
	SaveAccount(*Account) error
}
