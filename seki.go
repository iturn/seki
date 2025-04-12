package seki

import (
	"database/sql"
	"fmt"
	"log/slog"

	"net/http"
	"os"
	"strings"
)

type Seki struct {
	Log    *slog.Logger
	Db     *sql.DB
	Server *http.Server
	Mux    *http.ServeMux
}

func New() (s *Seki) {
	err := LoadEnvFile()

	log := NewLogger()

	log.Debug("Seki app bootstrap")

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Debug("Did not load .env file")
		} else {
			log.Error("Failed to load .env file")
			os.Exit(1)
		}

	} else {
		log.Debug("Did load .env file")
	}

	s = &Seki{
		Log: log,
		Mux: http.NewServeMux(),
	}

	// setup server
	hostString := fmt.Sprintf("%s:%s", os.Getenv("API_HOSTNAME"), os.Getenv("API_PORT"))
	s.Server = &http.Server{
		Addr:    hostString,
		Handler: s.Mux,
	}

	// setup db
	db, databaseErr := Connect(s)
	if databaseErr != nil {
		log.Error("Failed to connect to database " + databaseErr.Error())
		os.Exit(1)
	} else {
		s.Db = db
	}

	return
}
