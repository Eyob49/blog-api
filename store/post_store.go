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

// GET by ID
func (s *PostStore) GetByID(id int) (models.Post, error) {
	var p models.Post
	err := s.db.QueryRow(context.Background(), "SELECT id, title, content, created_at FROM posts WHERE id=$1", id).Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)

	if err != nil {
		return models.Post{}, err
	}
	return p, nil
	
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

// UPDATE 
func (s *PostStore) Update(id int, updated models.Post) (models.Post, error) {

	var p models.Post
	
	query := `

	    UPDATE posts
		SET title = $1, content = $2
		WHERE id = $3
		RETURNING id, title, content, created_at;
	`
	err := s.db.QueryRow(
		context.Background(),
		query,
		updated.Title,
		updated.Content,
		id,
	).Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)

	if err != nil {
		return models.Post{}, err
	}

	return p, nil

}

