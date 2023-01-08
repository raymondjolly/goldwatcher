package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed    = errors.New("update failed")
	errDeleteFailed    = errors.New("delete failed")
	errNewRecordFailed = errors.New("new record failed")
)

type Repository interface {
	Migrate() error
	InsertHolding(h Holdings) (*Holdings, error)
	AllHoldings() ([]Holdings, error)
	GetHoldingByID(id int) (*Holdings, error)
	UpdateHoldingByID(id int64, updated Holdings) error
	DeleteHolding(id int64) error
}

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
