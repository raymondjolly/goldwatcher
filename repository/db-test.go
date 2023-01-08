package repository

import (
	"time"
)

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (repo *TestRepository) Migrate() error {

	return nil
}

func (repo *TestRepository) InsertHolding(holdings Holdings) (*Holdings, error) {

	return &holdings, nil
}

func (repo *TestRepository) AllHoldings() ([]Holdings, error) {

	var result []Holdings

	return result, nil
}

func (repo *TestRepository) GetHoldingByID(id int) (*Holdings, error) {
	h := Holdings{
		Amount:        1,
		PurchasePrice: 1500,
		PurchaseDate:  time.Now(),
	}
	return &h, nil
}

func (repo *TestRepository) UpdateHoldingByID(id int64, updated Holdings) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {

	return nil
}
