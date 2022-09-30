package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Queries interface {
	UserQueries
}

type repository struct {
	db *sql.DB
}

func New() Queries {
	db, err := sql.Open("postgres", "postgresql://gustavolimam:gustavolimam@localhost/zombieland?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("Error trying to open connection with postgres")
	}

	return &repository{
		db: db,
	}
}
