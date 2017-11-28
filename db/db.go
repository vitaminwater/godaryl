package daryl_db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/config"
)

var db *sqlx.DB

func getEnv(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func migration(db *sql.DB) {
	migrationDir := config.AppContext.String("migration-dir")
	if migrationDir == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		migrationDir = fmt.Sprintf("%s/migrations", dir)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Info(err)
	}
}

func Init() {
	d, err := sqlx.Connect("postgres", config.AppContext.String("postgres-url"))
	if err != nil {
		log.Fatal(err)
	}
	db = d
	migration(db.DB)
}
