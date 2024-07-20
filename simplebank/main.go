package main

import (
	"context"
	"log"
	"time"

	"github.com/agolosnichenko/golang-simplebank/simplebank/api"
	db "github.com/agolosnichenko/golang-simplebank/simplebank/db/sqlc"
	"github.com/agolosnichenko/golang-simplebank/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
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

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
