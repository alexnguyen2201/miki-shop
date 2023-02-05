postgres:
	docker run --name postgres12 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root miki_shop

dropdb:
	docker exec -it postgres12 dropdb miki_shop

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/miki_shop?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/miki_shop?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/miki_shop?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/miki_shop?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nguyenvanson2201/miki-shop/db/sqlc Store

# migratefile:
# 	migrate create -ext sql -dir db/migration -seq {}

.PHONY: postgres craetedb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test mock