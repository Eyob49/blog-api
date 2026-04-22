package services

import (
	"blog-api/models"
	"blog-api/store"
)

type PostService struct {
	store *store.PostStore
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
	return s.store.Create(p)
}

func (s *PostService) Update(id int, p models.Post) (models.Post, error) {
	return s.store.Update(id, p)
}

func (s *PostService) Delete(id int) error {
	return s.store.Delete(id)
}