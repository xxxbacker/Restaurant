package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golangRestaurantManagement/api"
	db "golangRestaurantManagement/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "user=postgres password=postgres dbname=restaurantDB sslmode=disable" //"postgresql://root:postgres@locallhost:5432/restaurantDB?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	//connStr := "user=postgres password=postgres dbname=restaurantDB sslmode=disable"
	//conn, err := sql.Open("postgres", connStr)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connection db", err)
	}

	store := db.Newstore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
