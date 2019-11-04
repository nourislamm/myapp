package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const(
DbUser = "uh7iwqw1lvyesf9ojgvs"
DbPassword = "dfX89QUJcSO8kRXEkpv1"
DbName = "bzx9lkigjqdafipgmv2a"
DbHost = "bzx9lkigjqdafipgmv2a-postgresql.services.clever-cloud.com"
DbPort = "5432"
)

func main() {
    Info := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DbUser, DbPassword, DbName, DbHost, DbPort)

	db, err := sql.Open("postgres", Info)
	if err != nil {
		fmt.Print(err)
		fmt.Print("exit")
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)

		rows.Scan(&id, &name)

		log.Printf("%d: %s", id, name)
	}
}
