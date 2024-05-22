createdb:
	createdb --username=j2mf --owner=j2mf simple_bank

dropdb:
	dropdb simple_bank

migrate:
	migrate -path ./db/migrations -database "postgresql://j2mf:1122@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrated:
	migrate -path ./db/migrations -database "postgresql://j2mf:1122@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...


server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb simple-bank/db/sqlc Store

.PHONY: createdb dropdb migrate migrated sqlc test server mock
