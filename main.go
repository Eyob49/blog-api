package main

import (
	"blog-api/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// For demonstration, we will return a static list of posts
			posts := []models.Post{
				{ID: 1, Title: "First Post", Content: "This is the first post.", CreatedAt: time.Now()},
				{ID: 2, Title: "Second Post", Content: "This is the second post.", CreatedAt: time.Now()},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}) // Register the handler for the /posts endpoint

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}