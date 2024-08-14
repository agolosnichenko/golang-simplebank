package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/agolosnichenko/golang-simplebank/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer connPool.Close()

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
