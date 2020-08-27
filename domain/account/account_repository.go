package account

type AccountRepository interface {
	SaveAccount(*Account) error
}
