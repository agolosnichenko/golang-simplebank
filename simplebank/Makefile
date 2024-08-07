DB_URL=postgresql://golang:golang@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path ./db/migration -database "$(DB_URL)" up

migrateup1:
	migrate -path ./db/migration -database "$(DB_URL)" up 1

migratedown:
	migrate -path ./db/migration -database "$(DB_URL)" down

migratedown1:
	migrate -path ./db/migration -database "$(DB_URL)" down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb --build_flags=--mod=mod -destination db/mock/store.go github.com/agolosnichenko/golang-simplebank/simplebank/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/openapi/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc test server mock db_docs db_schema proto evans