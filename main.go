package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nguyenvanson2201/miki-shop/api"
	db "github.com/nguyenvanson2201/miki-shop/db/sqlc"
	"github.com/nguyenvanson2201/miki-shop/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load cofig")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
