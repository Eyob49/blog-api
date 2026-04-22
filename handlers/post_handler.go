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
		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	writeJSON(w, http.StatusOK, posts)
	 
	case http.MethodPost:
		var p models.Post

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		p.CreatedAt = time.Now()

		createdPost, err := h.service.Create(p)
		if err != nil {
			if _ ,ok := err.(*services.ValidationError); ok {
				writeError(w, http.StatusBadRequest, err.Error())
			} else {
				writeError(w, http.StatusInternalServerError, "Failed to create post")
			}
			
			return
		}
		writeJSON(w, http.StatusCreated, createdPost)

	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}  
}

func(h *PostHandler)  PostByID(w http.ResponseWriter, r *http.Request) {
        
	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	switch r.Method {

	case http.MethodGet:
		post, err := h.service.GetByID(id)
		if err != nil {
			if err == pgx.ErrNoRows{
				writeError(w, http.StatusNotFound, "Post not found")
				return
			} else {
				writeError(w, http.StatusInternalServerError, "Internal server error")
			}
			return 
		}
		writeJSON(w, http.StatusOK, post)

	case http.MethodPut:
		var updated models.Post

		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid request")
			return
		}


		post, err := h.service.Update(id, updated)
		if err != nil {
			if err == pgx.ErrNoRows{
				writeError(w, http.StatusNotFound, "Post not found")
			} else {
				writeError(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		writeJSON(w, http.StatusOK, post)

	case http.MethodDelete:
		err := h.service.Delete(id)
		if err != nil {
			if err == pgx.ErrNoRows {
				writeError(w, http.StatusNotFound, "Post not found")
			} else {
				writeError(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}


