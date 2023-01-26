package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/simple-bank?sslmode=disable"
)

var testQueries *Queries

// TestMain is main entry point for all test inside one golang package (db in our case)
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	// sqlc created New function to us
	testQueries = New(conn)

	os.Exit(m.Run())
}
