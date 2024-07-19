package main

import (
	"context"
	"log"
	"time"

	"github.com/agolosnichenko/simplebank/api"
	db "github.com/agolosnichenko/simplebank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
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

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
