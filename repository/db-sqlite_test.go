package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed")
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchasePrice: 1500,
		PurchaseDate:  time.Now(),
	}
	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed", err)
	}
	if result.ID <= 0 {
		t.Error("invalid id sent back", result.ID)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed", err)
	}
	if len(h) != 1 {
		t.Error("wrong number of rows returned; expected 1. got", len(h))
	}
}

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("Expecting id: 1 but got ", h.ID)
	}
	if h.PurchasePrice != 1500 {
		t.Error("Incorrect Purchase Price. Wanted 1500 but got ", h.PurchasePrice)
	}
	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("a record was returned for a none existent id")
	}
}

func TestSQLiteRepository_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("No record is being returned")
	}
	h.Amount = 2
	h.PurchasePrice = 2000

	err = testRepo.UpdateHoldingByID(1, *h)
	if err != nil {
		t.Error("update failed", err)
		if err != errUpdateFailed {
			t.Error("incorrect error returned")
		}
	}
}

func TestSQLiteRepository_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding", err)
		if err != errDeleteFailed {
			t.Error("incorrect error returned")
		}
	}

	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete a non existent id")
	}

}
