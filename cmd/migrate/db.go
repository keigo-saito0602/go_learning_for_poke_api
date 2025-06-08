package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/keigo-saito0602/go_learning_for_poke_api/config"
)

func newMigrator() (*migrate.Migrate, error) {
	config.LoadConfig()

	db := config.AppConfig
	dbURL := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		db.DBUser, db.DBPass, db.DBHost, db.DBPort, db.DBName,
	)
	return migrate.New("file://assets/migrations", dbURL)
}
