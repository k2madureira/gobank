postgres:
		docker run --name estudoPostgreSql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres 

createdb:
		docker exec -it estudoPostgreSql createdb --username=postgres --owner=postgres simple_bank

dropdb:
		docker exec -it estudoPostgreSql dropdb --username=postgres simple_bank

migrateup:
		migrate -path migrations -database "postgres://password:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationdown:
		migrate -path migrations -database "postgres://password:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down		


sqlc:
		sqlc generate

.PHONY: postgres createdb dropdb migrateup migrationdown sqlc