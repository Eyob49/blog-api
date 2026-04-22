package services

import (
	"blog-api/models"
	"blog-api/store"
	"errors"
)

type PostService struct {
	store *store.PostStore
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}


func NewPostService(store *store.PostStore) *PostService {
	return &PostService{store: store}
}


// Methods 

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.store.GetAll()
}

func (s *PostService) GetByID(id int) (models.Post, error) {
	return s.store.GetByID(id)
}

func (s *PostService) Create(p models.Post) (models.Post, error) {
	if p.Title == "" {
		return models.Post{}, &ValidationError{Message: "title is required"}
	}
	if p.Content == "" {
		return models.Post{}, &ValidationError{Message: "content is required"}
	}


	return s.store.Create(p)
}

func (s *PostService) Update(id int, p models.Post) (models.Post, error) {

	if p.Title == "" {
		return models.Post{}, errors.New("title is required")
	}
	if p.Content == "" {
		return models.Post{}, errors.New("content is required")
	}

	return s.store.Update(id, p)
}

func (s *PostService) Delete(id int) error {
	return s.store.Delete(id)
}