package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"work-acc/wildberries-L0/internal/config"
)

func NewPostgresDB(cfg *config.PostgreSQL) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password='%s' sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Name,
		cfg.Password,
		cfg.SslMode))

	if err != nil {
		log.Fatalln("error connection to DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("ping DB failed")
		return nil, err
	}

	return db, nil
}
