postgres:
		docker run --name estudoPostgreSql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres 

createdb:
		docker exec -it estudoPostgreSql createdb --username=postgres --owner=postgres simple_bank

dropdb:
		docker exec -it estudoPostgreSql dropdb --username=postgres simple_bank

migrationup:
		migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationuplast:
		migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up 1		

migrationdown:
		migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down		

migrationdownlast:
		migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
		sqlc generate

test:
		go test -v -cover ./...		

server:
		go run main.go		

mock:
		 mockgen -package mockdb -destination db/mock/store.go github.com/k2madureira/gobank/db/sqlc Store

.PHONY: postgres createdb dropdb migrationup migrationuplast migrationdown migrationdownlast sqlc server mock