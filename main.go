package main

import (
	"database/sql"
	"log"

	"github.com/k2madureira/gobank/api"
	dbbank "github.com/k2madureira/gobank/db/sqlc"
	"github.com/k2madureira/gobank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("caccot connect to db:", err)
	}

	store := dbbank.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
