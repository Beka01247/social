package main

import (
	"log"

	"github.com/Beka01247/social/internal/db"
	"github.com/Beka01247/social/internal/store"
)

func main() {
	//addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New("postgres://admin:adminpassword@localhost/social?sslmode=disable", 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
