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
)

var db *sqlx.DB

func getEnv(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func migration() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", "postgres://daryl:daryl@localhost:5432/daryl?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/migrations", dir),
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Info(err)
	}
}

func init() {
	migration()
	log.Info("pouet")
	host := getEnv("PG_HOST", "localhost")
	user := getEnv("PG_USER", "daryl")
	password := getEnv("PG_PASSWORD", "daryl")
	dbname := getEnv("PG_DB", "daryl")

	d, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
