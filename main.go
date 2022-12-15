package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ahmed"
	password = "1609"
	dbname   = "ahmed"
	sslmode  = "disable"
)

func main() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", user, password, host, port, dbname, sslmode)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	defer db.Close()
	//heckError(e)
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	fmt.Println("connected")
}
