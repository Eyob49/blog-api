package handlers

import (
	"blog-api/models"
	"blog-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)


type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Posts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		posts, err := h.service.GetAll()
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

		createdPost, err := h.service.Create(p)
		if err != nil {
			http.Error(w, "Failed to create post", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdPost)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}  
}

func(h *PostHandler)  PostByID(w http.ResponseWriter, r *http.Request) {
        
	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		post, err := h.service.GetByID(id)
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


		post, err := h.service.Update(id, updated)
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

	case http.MethodDelete:
		err := h.service.Delete(id)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Error(w, "Post not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


