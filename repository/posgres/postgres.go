package posgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "soburjon"
	password = "0226"
	dbname   = "prstore"
)

func Connect() *sql.DB {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, er := sql.Open("postgres", psql)
	if er != nil {
		log.Fatal("Fatal1: ", er)
	}
	return db
}
