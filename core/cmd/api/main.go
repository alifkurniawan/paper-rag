package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"core/config"
	"core/internal/db"
	"core/internal/db/sqlc"
	"core/internal/user"

	"github.com/go-chi/chi/v5"
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

	queries := sqlc.New(pool)
	userRepo := user.NewRepository(queries)
	userService := user.NewService(userRepo)
	userHandler := *user.NewHandler(userService)

	r := chi.NewRouter()
	userHandler.RegisterRoutes(r)

	fmt.Println("API Server is running on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
