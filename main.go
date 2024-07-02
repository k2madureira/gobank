package main

import (
	"database/sql"
	"log"

	"github.com/k2madureira/gobank/api"
	dbbank "github.com/k2madureira/gobank/db/sqlc"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
const serverAddress = "0.0.0.0:3000"

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("caccot connect to db:", err)
	}

	store := dbbank.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
