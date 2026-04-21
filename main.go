package main

import (
	"blog-api/store"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)



func main() {
    // Connect to DB
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@49@localhost:5432/blogdb")
    if err != nil {
    log.Fatal(err)
    }
    defer conn.Close(context.Background())

    // Initialize store with DB
    postStore := store.NewPostStore(conn)

    http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
        posts, err := postStore.GetAll()
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(posts)
        
    })
    //http.HandleFunc("/posts/", postsHandlerByID)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}



// func postsHandlerByID(w http.ResponseWriter, r *http.Request) {
    
// }