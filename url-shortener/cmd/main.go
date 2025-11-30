package main

import (
	"log"

	"github.com/rezaabaskhanian/go-url-shortener/internal/repository/postgres"
)

func main() {
	// connStr := "postgres://user:pass@localhost:5432/url_shortener?sslmode=disable"

	// pool, err := pgxpool.New(context.Background(), connStr)
	// if err != nil {
	// 	log.Fatal("Unable to connect to database:", err)
	// }
	// defer pool.Close()

	mydbPostGress := postgres.New()

	myurlRepo := postgres.NewMyPostgres(mydbPostGress)

	log.Println("Connected to Postgres successfully!")

}
