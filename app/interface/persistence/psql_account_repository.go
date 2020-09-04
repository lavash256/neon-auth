package persistence

import (
	"database/sql"
	"fmt"
	"neon-auth/app/domain/model"

	_ "github.com/lib/pq" //Is needed to create postgres connection
)

//PostgresConfig ...
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
}

//PsqlAccountRepository is postgres db based storage implementation
type PsqlAccountRepository struct {
	DB        *sql.DB
	NameTable string
}

//NewPsqlAccountRepository is fabric to create PsqlAccountRepository
func NewPsqlAccountRepository(config PostgresConfig) (*PsqlAccountRepository, error) {
	table := "accounts"
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	fmt.Printf("DATABASE string %s", databaseURL)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	rep := &PsqlAccountRepository{DB: db, NameTable: table}
	if err := rep.DB.Ping(); err != nil {
		return nil, err
	}
	return &PsqlAccountRepository{DB: db, NameTable: table}, nil
}

//SaveAccount saves account in postgres db
func (repo *PsqlAccountRepository) SaveAccount(account *model.Account) error {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (email,password,created_at,status) VALUES ($1,$2,$3,$4) RETURNING id", repo.NameTable)
	err := repo.DB.QueryRow(sqlQuery,
		account.Email,
		account.Password,
		account.CreatedAt,
		account.Status).Scan(&account.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail finds an account at the given email address
func (repo *PsqlAccountRepository) FindByEmail(email string) (*model.Account, error) {
	row := repo.DB.QueryRow("SELECT id,email,password,status FROM accounts WHERE email=$1", email)
	var ID uint64
	var Email string
	var Password string
	var Status string
	err := row.Scan(&ID, &Email, &Password, &Status)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	account := model.BuildAccount(ID, Email, Password, Status)
	return account, nil

}
