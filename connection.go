package seki

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(s *Seki) (db *sql.DB, err error) {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	connectString := fmt.Sprintf("%s:%s@(%s:%s)/%s", username, password, hostname, port, name)

	db, err = sql.Open("mysql", connectString)

	s.Log.Debug("Database connect attempt to " + connectString)

	if err != nil {
		return nil, err
	}

	s.Log.Debug("Database connected")

	return db, nil
}
