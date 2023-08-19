package database

import (
	"devstream-rest-api/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgDatabase struct {
	config *config.Config
	DB     *sqlx.DB
}

func NewPgDatabase(config *config.Config) *PgDatabase {
	return &PgDatabase{
		config: config,
	}
}

func (database *PgDatabase) Connect() {

	db, err := sqlx.Open("postgres", database.config.DB_URL)

	if err != nil {
		logger.E(err)
	}

	if err = db.Ping(); err != nil {
		logger.E(err)
	}

	logger.I("Connected to postgres")

	database.DB = db
}
