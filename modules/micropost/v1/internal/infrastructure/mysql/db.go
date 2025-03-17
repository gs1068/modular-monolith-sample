package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(connection string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", connection)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, err
}
