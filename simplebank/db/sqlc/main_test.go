package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	// Configure connection pool settings
	config.MaxConns = 10                   // Set the maximum number of connections
	config.MinConns = 1                    // Set the minimum number of connections
	config.MaxConnLifetime = 0             // No maximum connection lifetime
	config.MaxConnIdleTime = 0             // No maximum idle time
	config.HealthCheckPeriod = time.Minute // Set the health check period

	testDB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
