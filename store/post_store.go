package store

import (
	"blog-api/models"
	"sync"
)

type PostStore struct {
	mu sync.RWMutex
	posts []models.Post
	nextID int
}

func NewPostStore() *PostStore {
	return &PostStore{
		posts: []models.Post{},
		nextID: 1,
	}
}

// ReadAll 
func (s *PostStore) GetAll() []models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]models.Post(nil), s.posts...)
} 

// GET by ID 
func (s *PostStore) GetByID(id int) (models.Post, bool){
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _,post := range s.posts {
		if post.ID == id {
			return post, true
		}
	}
	return models.Post{}, false
}

// CREATE
func (s *PostStore) Create(post models.Post) models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	post.ID = s.nextID
	s.nextID++
	s.posts = append(s.posts, post)
	return post
}

// UPDATE
func (s *PostStore) Update(id int, updated models.Post) (models.Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
    
	for i, post := range s.posts {
		if post.ID == id {
			updated.ID = id
			s.posts[i] = updated
			return updated, true
		}
	}
	return models.Post{}, false
}

// DELETE
func (s *PostStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, post := range s.posts {
		if post.ID == id {
			s.posts = append(s.posts[:i], s.posts[i+1:]...)
			return true
		}
	}
	return false
}