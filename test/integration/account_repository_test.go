// +build integration

package integration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/lavash256/neon-auth/internal/domain/model"
	"github.com/lavash256/neon-auth/internal/interface/persistence"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

var (
	user     = "user"
	password = "password"
	db       = "db"
	port     = "7777"
)

var (
	accountRepoPostgres *persistence.PsqlAccountRepository
)

//Start postgres and create repository
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + db,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	if err = pool.Retry(func() error {
		absPath, _ := filepath.Abs("../../migrations/")
		m, err := migrate.New(
			"file://"+absPath,
			fmt.Sprintf("postgres://%s:%s@localhost:7777/%s?sslmode=disable", user, password, db),
		)
		if err != nil {
			return err
		}
		m.Steps(2)
		config := persistence.PostgresConfig{Host: "localhost", Port: port, User: user, Password: password, Name: db}

		accountRepoPostgres, err = persistence.NewPsqlAccountRepository(config)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		accountRepoPostgres.Close()
	}()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
func TestAccountRepositoryIntegrationSaveAccount(t *testing.T) {
	account, err := model.NewAccount("test@test.ru", "test")
	assert.Equal(t, err, nil)
	err = accountRepoPostgres.SaveAccount(account)
	if err != nil {
		assert.Equal(t, err, nil, "There should be no errors when adding to a clean base")
		os.Exit(1)
	}
	//When you add the same account again, the code should return an error
	err = accountRepoPostgres.SaveAccount(account)
	assert.Equal(t, err.Error(), "pq: duplicate key value violates unique constraint \"accounts_email_key\"")

}
func TestAccountRepositoryIntegrationFindAccount(t *testing.T) {

	account, err := model.NewAccount("find@test.ru", "test")
	assert.Equal(t, err, nil)
	err = accountRepoPostgres.SaveAccount(account)
	if err != nil {
		assert.Equal(t, err, nil, "There should be no errors when adding to a clean base")
		os.Exit(1)
	}

	newAccount, err := accountRepoPostgres.FindByEmail("find@test.ru")
	if err != nil {
		assert.Equal(t, err, nil, "When receiving from the database there should be no errors")
		os.Exit(1)
	}

	assert.Equal(t, newAccount.Email, account.Email, "Emails must match when received from the repository")
	assert.Equal(t, newAccount.Password, account.Password, "Passwords must match when received from the repository")
	assert.Equal(t, newAccount.Status, account.Status, "Status must match when received from the repository")

}
