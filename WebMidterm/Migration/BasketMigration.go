package Migration

import (
	"database/sql"
)


type SQLiteRepository struct {
	Db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS baskets(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL,
        basketData TEXT NOT NULL,
        status TEXT NOT NULL
    );
    `

	_, err := r.Db.Exec(query)
	return err
}