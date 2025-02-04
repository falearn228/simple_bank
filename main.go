package main

import (
	"database/sql"
	"log"

	"github.com/simple_bank_course/api"
	db "github.com/simple_bank_course/db/sqlc"
	"github.com/simple_bank_course/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load configurations: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("can't connect to DB: ", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Address)
	if err != nil {
		log.Fatalln("cant't start a server: ", err)
	}
}
