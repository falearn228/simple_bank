package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
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
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("can't connect to DB: ", err)
	}
	defer testDB.Close()
	testQueries = New(testDB)

	os.Exit(m.Run())
}
