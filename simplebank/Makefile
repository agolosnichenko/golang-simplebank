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

.PHONY: migrateup migratedown sqlc test server