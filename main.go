package main

import (
	"blog-api/handlers"
	"blog-api/services"
	"blog-api/store"
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)



func main() {
    // Connect to DB
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres%4049@localhost:5432/blogdb")
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



