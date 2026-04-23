package main

import (
	"blog-api/handlers"
	"blog-api/services"
	"blog-api/store"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)



func main() {


    err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
    }
    // Connect to DB
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
    log.Fatal(err)
    }
    defer conn.Close(context.Background())

    // Initialize store with DB
    postStore := store.NewPostStore(conn)

    service := services.NewPostService(postStore)
    handler := handlers.NewPostHandler(service) 

    

    http.HandleFunc("/posts", handler.Posts)
    http.HandleFunc("/posts/", handler.PostByID)
    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}




