package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists holdings(
	    id integer primary key autoincrement,
	    amount real not null,
	    purchase_date integer not null,
	    purchase_price integer not null
	    );
`
	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertHolding(holdings Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount, purchase_date, purchase_price) values (?, ?, ?)"

	result, err := repo.Conn.Exec(stmt, holdings.Amount, holdings.PurchaseDate.Unix(), holdings.PurchasePrice)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	holdings.ID = id
	return &holdings, nil
}

func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings order by purchase_date"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []Holdings
	for rows.Next() {
		var h Holdings
		var unixTime int64
		err := rows.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		result = append(result, h)
	}
	return result, nil
}

func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings where id = ?"
	row := repo.Conn.QueryRow(query, id)

	var h Holdings
	var unixTime int64
	err := row.Scan(
		&h.ID, &h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}
	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

func (repo *SQLiteRepository) UpdateHoldingByID(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid updated Id")
	}

	stmt := "update holdings set amount = ?, purchase_date=?, purchase_price = ? where id = ?"
	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errUpdateFailed
	}
	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {

	stmt := "delete from holdings where id = ?"
	res, err := repo.Conn.Exec(stmt, id)
	if err != nil {
		return err
	}
	rowDeleted, err := res.RowsAffected()
	if rowDeleted == 0 {
		return errDeleteFailed
	}
	return nil
}
