package repo

import (
	"context"
	"go-clean-microblog/internal/domain"
)

type PostRepo struct {
	posts  []*domain.Post
	nextID int64
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		posts:  []*domain.Post{},
		nextID: 1,
	}
}

func (r *PostRepo) Create(ctx context.Context, post *domain.Post) error {
	// In-memory implementation: assign incremental ID
	post.ID = r.nextID
	r.nextID++
	r.posts = append(r.posts, post)
	return nil
}

func (r *PostRepo) FindAll(ctx context.Context) ([]*domain.Post, error) {
	// Return a copy of all posts
	result := make([]*domain.Post, len(r.posts))
	copy(result, r.posts)
	return result, nil
}
