package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/agolosnichenko/golang-simplebank/simplebank/api"
	db "github.com/agolosnichenko/golang-simplebank/simplebank/db/sqlc"
	_ "github.com/agolosnichenko/golang-simplebank/simplebank/doc/statik"
	"github.com/agolosnichenko/golang-simplebank/simplebank/gapi"
	"github.com/agolosnichenko/golang-simplebank/simplebank/pb"
	"github.com/agolosnichenko/golang-simplebank/simplebank/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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
	go runGatewayServer(config, store)
	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatalf("Cannot create listener: %v", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Cannot start gRPC server: %v", err)
	}
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatalf("cannot register handler server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf("cannot create statik fs: %v", err)
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)

	if err != nil {
		log.Fatalf("Cannot create listener: %v", err)
	}

	log.Printf("start HTTP Gateway server on %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("Cannot start HTTP Gateway server: %v", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
