package main

import (
	"blog-api/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var posts []models.Post
var nextID = 1
func main() {
    http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
        idStr := r.URL.Path[len("/posts/"):]

        switch r.Method {
        case http.MethodGet:
            // If the ID part is empty, they want all posts
            if idStr == "" {
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(posts)
                return
            }

            // Otherwise, look for the specific post
            for _, post := range posts {
                if idStr == strconv.Itoa(post.ID) {
                    w.Header().Set("Content-Type", "application/json")
                    json.NewEncoder(w).Encode(post)
                    return
                }
            }
            http.Error(w, "Post not found", http.StatusNotFound)

        case http.MethodPost:
            if idStr != "" {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                return
            }
            
            var newPost models.Post
            if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
                http.Error(w, "Invalid request payload", http.StatusBadRequest)
                return
            }
            newPost.ID = nextID
            newPost.CreatedAt = time.Now()
            nextID++
            posts = append(posts, newPost)
            
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(newPost)
        case http.MethodDelete:
			if idStr == "" {
				http.Error(w, "Post ID is required", http.StatusBadRequest)
				return
			}
			for i, post := range posts {
				if idStr == strconv.Itoa(post.ID) {
					posts = append(posts[:i], posts[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.Error(w, "Post not found", http.StatusNotFound)


        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}