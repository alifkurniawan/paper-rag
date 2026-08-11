package main

import (
	"context"
	"fmt"
	"log"

	"core/config"
	"core/internal/db"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Initialize the database connection pool
	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	fmt.Println("connected to database successfully")
}
