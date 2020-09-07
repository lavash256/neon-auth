package persistence

import "neon-auth/src/app/domain/model"

//MemoryAccountRepository is a stub for account repository
type MemoryAccountRepository struct {
	accounts []*model.Account
}

//SaveAccount  saves account in memory storage
func (s *MemoryAccountRepository) SaveAccount(account *model.Account) error {
	s.accounts = append(s.accounts, account)
	return nil
}

//FindByEmail finds an account by the specified email
func (s *MemoryAccountRepository) FindByEmail(email string) (*model.Account, error) {
	for _, value := range s.accounts {
		currentEmail := value.Email
		if currentEmail == email {
			return value, nil
		}
	}
	return nil, nil
}
