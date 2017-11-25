package daryl_db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

var db *sqlx.DB

func getEnv(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func Init() {
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
