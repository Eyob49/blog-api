package store

import (
	"blog-api/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type PostStore struct {
	db *pgx.Conn
}

func NewPostStore(db *pgx.Conn) *PostStore {
	return &PostStore{db: db}
}

// GET All 
func (s *PostStore) GetAll() ([]models.Post, error) {
	rows, err := s.db.Query(context.Background(), "SELECT id, title, content, created_at FROM posts")

	if err != err {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err

		}
		posts = append(posts, p)
	}
	return posts, nil
} 

// CREATE
func (s *PostStore) Create(post models.Post) (models.Post, error) {
	query := `
	    INSERT INTO posts (title, content, created_at)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	err := s.db.QueryRow(
		context.Background(),
		query,
		post.Title,
		post.Content,
		post.CreatedAt,
	).Scan(&post.ID)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
