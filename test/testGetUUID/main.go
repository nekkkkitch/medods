package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var (
	login    = "123"
	host     = "localhost"
	port     = 5434
	user     = "user"
	password = "123"
	dbname   = "medodsdb"
)

// get user uuid here
func main() {
	connection := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, password, host, port, dbname)
	d, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		log.Fatalln(err)
	}
	var userID string
	err = d.QueryRow(context.Background(), `select id from public.users where login = $1`, login).Scan(&userID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(userID)
}
