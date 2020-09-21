package register

import (
	"github.com/Lavash95/neon-auth/internal/interface/persistence"
	"github.com/Lavash95/neon-auth/internal/usecase"
	"github.com/Lavash95/neon-auth/tools"
)

//AccountUsecaseBuilder is needed to hide the dependencies required to create UseCase
//Receives the following string as input: "user=postgres password=mypass dbname=productdb sslmode=disable"
//Returns an error if unable to connect to the database
func AccountUsecaseBuilder(config persistence.PostgresConfig) (*usecase.AccountUsecase, error) {
	accountRepository, err := persistence.NewPsqlAccountRepository(config)
	if err != nil {
		return nil, err
	}
	logger := tools.NewLogrusLogger()
	accountUsecase := usecase.NewAccountUsecase(accountRepository, logger)
	return accountUsecase, nil
}
