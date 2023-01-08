package repository

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"os"
	"testing"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}
	testRepo = NewSQLiteRepository(db)
	os.Exit(m.Run())
}
