migrateup:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" up

migratedown:
	migrate -path ./db/migration -database "postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb --build_flags=--mod=mod -destination db/mock/store.go github.com/agolosnichenko/golang-simplebank/simplebank/db/sqlc Store

.PHONY: migrateup migratedown sqlc test server mock