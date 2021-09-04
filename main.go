package main

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"ent-validation-bug/ent"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	err = client.Schema.Create(context.TODO())
	if err != nil {
		log.Panic(err)
	}

	var sb strings.Builder
	for i := 0; i < 255; i++ {
		sb.WriteRune('Й')
	}

	s := sb.String()
	log.Println(len(s), len([]rune(s))) // 510 bytes and 255 chars

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	row := db.QueryRow(`INSERT INTO works (name) VALUES($1) RETURNING id, name`, s)

	var id int64
	var name string
	err = row.Scan(&id, &name)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Inserted into postgres", id, name)
	// Error occurs here
	client.Work.Create().SetName(s).ExecX(context.TODO())
}
