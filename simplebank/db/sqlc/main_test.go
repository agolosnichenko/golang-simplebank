package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/agolosnichenko/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	dbConfig, err := pgxpool.ParseConfig(config.DbSource)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	// Configure connection pool settings
	dbConfig.MaxConns = 10                   // Set the maximum number of connections
	dbConfig.MinConns = 1                    // Set the minimum number of connections
	dbConfig.MaxConnLifetime = 0             // No maximum connection lifetime
	dbConfig.MaxConnIdleTime = 0             // No maximum idle time
	dbConfig.HealthCheckPeriod = time.Minute // Set the health check period

	testDB, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
