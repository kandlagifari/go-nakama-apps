package main

import (
	"log"

	"github.com/kandlagifari/go-nakama-apps/internal/db"
	"github.com/kandlagifari/go-nakama-apps/internal/env"
	"github.com/kandlagifari/go-nakama-apps/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/nakama_db?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
