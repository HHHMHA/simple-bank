package main

import (
	"database/sql"
	"log"
	"simple-bank/api"
	db "simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://j2mf:1122@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal(err)
	}

}
