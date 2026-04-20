package main

import (
	"blog-api/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var posts []models.Post
var nextID = 1

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// By return if the func NewPost is working
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
			return
		} 

		if r.Method == http.MethodPost {
			var newPost models.Post
			err := json.NewDecoder(r.Body).Decode(&newPost)
			if err != nil {
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
            return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}) // Register the handler for the /posts endpoint
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}