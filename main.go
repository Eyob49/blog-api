package main

import (
	"blog-api/models"
	"blog-api/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var postStore = store.NewPostStore()

func main() {

    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/posts/", postsHandlerByID)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(postStore.GetAll())

    case http.MethodPost:
        var post models.Post
        json.NewDecoder(r.Body).Decode(&post)

        post.CreatedAt = time.Now()
        created := postStore.Create(post)

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(created)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func postsHandlerByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    switch r.Method {
    case http.MethodGet:
        post, found := postStore.GetByID(id)
        if !found {
            http.Error(w, "Post not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(post)
    case http.MethodPut:
        var post models.Post
        json.NewDecoder(r.Body).Decode(&post)

        updated, found := postStore.Update(id, post)
        if !found {
            http.Error(w, "Post not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(updated)
    case http.MethodDelete:
        ok := postStore.Delete(id)
        if !ok {
            http.Error(w, "Post not found", http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}