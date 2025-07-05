package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"bobbabank/internal/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// conn, err := pgxpool.New(context.Background(), dbSource)
	// if err != nil {
	// 	log.Fatalln("can't connect to DB: ", err)
	// }
	// defer conn.Close()

	// testQueries = New(conn)
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("can't load configurations: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("can't connect to DB: ", err)
	}
	defer testDB.Close()
	testQueries = New(testDB)

	os.Exit(m.Run())
}
