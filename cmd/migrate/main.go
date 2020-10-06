package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lavash256/neon-auth/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	var configPath string
	var migrateDir string
	var flagMigrate int
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
	flag.StringVar(&migrateDir, "migrate-dir", "migrations/", "path to migrate dir")
	flag.IntVar(&flagMigrate, "flag-migrate", 1, "flag migrate")
	flag.Parse()
	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	absPath, _ := filepath.Abs(migrateDir)
	m, err := migrate.New(
		"file://"+absPath,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			configFile.Database.User, configFile.Database.Password, configFile.Database.Host,
			configFile.Database.Port, configFile.Database.Name),
	)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	_ = m.Steps(flagMigrate)
	logrus.Info("Migrations were successful")

}
