package persistence

import (
	"testing"

	"github.com/lavash256/neon-auth/internal/domain/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPsqlRepositorySaveAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var returningID uint64 = 1
	account, _ := model.NewAccount("test@test.ru", "Test")
	mock.ExpectQuery("INSERT INTO *").WillReturnRows(func() *sqlmock.Rows {
		rr := sqlmock.NewRows([]string{"id"})
		rr.AddRow(1)
		return rr
	}())

	accountRepository := &PsqlAccountRepository{DB: db, NameTable: "account"}

	err = accountRepository.SaveAccount(account)
	assert.Equal(t, err, nil)
	assert.Equal(t, account.ID, returningID)
	defer func() {
		err = accountRepository.Close()
		assert.Equal(t, err.Error(), "all expectations were already fulfilled, call to database Close was not expected", "Close database connection")
	}()
}

func TestPsqlRepositoryFindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	result := []string{"id", "email", "password", "status"}
	var returningID uint64 = 1
	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(sqlmock.NewRows(result).AddRow(returningID, "test@test.ru", "Password", "Created"))
	accountRepository := &PsqlAccountRepository{DB: db, NameTable: "account"}
	account, err := accountRepository.FindByEmail("test@test.ru")
	assert.Equal(t, err, nil)
	assert.Equal(t, account.ID, returningID)
	assert.Equal(t, account.Email, "test@test.ru")
	assert.Equal(t, account.Password, "Password")
	defer func() {
		err = accountRepository.Close()
		assert.Equal(t, err.Error(), "all expectations were already fulfilled, call to database Close was not expected", "Close database connection")
	}()

}
