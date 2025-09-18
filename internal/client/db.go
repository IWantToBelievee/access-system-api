package client

import (
	"database/sql"
	"fmt"

	"access-system-api/internal/cfg"

	_ "github.com/lib/pq"
)

// ConnectDB establishes a connection to the PostgreSQL database using the provided configuration.
func ConnectDB(dbCfg *cfg.DbCfg) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName)
	db, err := sql.Open(`postgres`, connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
