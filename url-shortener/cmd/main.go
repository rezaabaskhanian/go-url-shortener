package main

import (
	"context"
	"log"

	"github.com/rezaabaskhanian/go-url-shortener/internal/delivery/httpserver"
	urlhandler "github.com/rezaabaskhanian/go-url-shortener/internal/delivery/httpserver/handler"
	"github.com/rezaabaskhanian/go-url-shortener/internal/repository/postgres"
	"github.com/rezaabaskhanian/go-url-shortener/internal/usecase"
)

func main() {
	// connStr := "postgres://user:pass@localhost:5432/url_shortener?sslmode=disable"

	// pool, err := pgxpool.New(context.Background(), connStr)
	// if err != nil {
	// 	log.Fatal("Unable to connect to database:", err)
	// }
	// defer pool.Close()

	mydbPostGress := postgres.New()

	if err := mydbPostGress.Ping(context.Background()); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	} else {
		log.Println("Connected OK!")
	}

	myUrlRepo := postgres.NewMyPostgres(mydbPostGress)

	svc := usecase.New(myUrlRepo)

	handler := urlhandler.New(svc)

	server := httpserver.New(handler)

	server.Serve()

}
