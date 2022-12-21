package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	cre "test/credentials"
)

func main() {
	var config cre.CredStructure
	config.DbName = "ahmed"
	config.User = "ahmed"
	config.Password = "1609"
	config.Host = "localhost"
	config.Port = 5432
	config.SslMOde = "disable"

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", config.User, config.Password, config.Host, config.Port, config.DbName, config.SslMOde)
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
