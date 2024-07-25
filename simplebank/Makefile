migrateup:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" up

migrateup1:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" up 1

migratedown:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" down

migratedown1:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb --build_flags=--mod=mod -destination db/mock/store.go github.com/agolosnichenko/golang-simplebank/simplebank/db/sqlc Store

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc test server mock