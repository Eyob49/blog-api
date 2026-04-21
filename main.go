package main

import (
	"blog-api/models"
	"blog-api/store"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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

        switch r.Method {

        case http.MethodGet:
            posts, err := postStore.GetAll()
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(posts)
         
        case http.MethodPost:
            var p models.Post

            err := json.NewDecoder(r.Body).Decode(&p)
            if err != nil {
                http.Error(w, "Invalid request", http.StatusBadRequest)
                return
            }

            p.CreatedAt = time.Now()

            createdPost, err := postStore.Create(p)
            if err != nil {
                http.Error(w, "Failed to create post", http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(createdPost)

        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }  
    })

    http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
        
        idStr := strings.TrimPrefix(r.URL.Path, "/posts/")

        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        switch r.Method {

        case http.MethodGet:
            post, err := postStore.GetByID(id)

            if err != nil {
                if err == pgx.ErrNoRows{
                    http.Error(w, "Post not found", http.StatusNotFound)
                    return
                } else {
                    http.Error(w, "Internal server error", http.StatusInternalServerError)
                }
                return 
            }
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(post)

        case http.MethodPut:
            var updated models.Post

            err := json.NewDecoder(r.Body).Decode(&updated)
            if err != nil {
                http.Error(w, "Invalid request", http.StatusBadRequest)
                return
            }


            post, err := postStore.Update(id, updated)
            if err != nil {
                if err == pgx.ErrNoRows{
                    http.Error(w, "Post Not Foud", http.StatusNotFound)
                } else {
                    http.Error(w, "Internal server error", http.StatusInternalServerError)
                }
                return
            }

            w.Header().Set("Content-type", "application/json")
            json.NewEncoder(w).Encode(post)



        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
   

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}



