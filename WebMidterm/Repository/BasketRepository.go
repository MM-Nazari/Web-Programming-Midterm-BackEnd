package Repository

import (
	"WebMidterm/Model"
	"database/sql"
	"errors"
)

var (
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
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

func (r *SQLiteRepository) Create(basket Model.Basket) (*Model.Basket, error) {
	res, err := r.Db.Exec("INSERT INTO baskets(created_at, updated_at, basketData, status) values(?,?,?,?)", basket.CreatedAt, basket.UpdatedAt, basket.Data, basket.Status)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	basket.ID = int64(int(id))

	return &basket, nil
}

func (r *SQLiteRepository) All() ([]Model.Basket, error) {
	rows, err := r.Db.Query("SELECT * FROM baskets")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var all []Model.Basket
	for rows.Next() {
		var basket Model.Basket
		if err := rows.Scan(&basket.ID, &basket.CreatedAt, &basket.UpdatedAt, &basket.Data, &basket.Status); err != nil {
			return nil, err
		}
		all = append(all, basket)
	}
	return all, nil
}

func (r *SQLiteRepository) GetById(id int64) (*Model.Basket, error) {
	row := r.Db.QueryRow("SELECT * FROM baskets WHERE id = ?", id)

	var basket Model.Basket
	if err := row.Scan(&basket.ID, &basket.CreatedAt, &basket.UpdatedAt, &basket.Data, &basket.Status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &basket, nil
}

func (r *SQLiteRepository) Update(id int64, updated Model.Basket) (*Model.Basket, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}

	if updated.Status == "COMPLETED" {
		return nil, errors.New("basket status is completed, cannot be changed")
	}

	res, err := r.Db.Exec("UPDATE baskets SET created_at = ?, updated_at = ?, basketData = ?, status = ? WHERE id = ?", updated.CreatedAt, updated.UpdatedAt, updated.Data, updated.Status, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.Db.Exec("DELETE FROM baskets WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
