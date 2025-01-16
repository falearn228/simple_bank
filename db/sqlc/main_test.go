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

func TestMain(m *testing.M) {
	// conn, err := pgxpool.New(context.Background(), dbSource)
	// if err != nil {
	// 	log.Fatalln("can't connect to DB: ", err)
	// }
	// defer conn.Close()

	// testQueries = New(conn)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("can't connect to DB: ", err)
	}
	defer conn.Close()
	testQueries = New(conn)

	os.Exit(m.Run())
}
