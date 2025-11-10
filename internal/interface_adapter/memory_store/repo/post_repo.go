package repo

import (
	"context"
	"go-clean-microblog/internal/domain"
	"sync"
)

type PostRepo struct {
	posts  []*domain.Post
	nextID int64
	mu     sync.Mutex
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		posts:  make([]*domain.Post, 0),
		nextID: 1,
	}
}

func (r *PostRepo) Create(ctx context.Context, post *domain.Post) error {
	// In-memory implementation: assign incremental ID
	r.mu.Lock()
	defer r.mu.Unlock()
	post.ID = r.nextID
	r.nextID++
	r.posts = append(r.posts, post)
	return nil
}

func (r *PostRepo) FindAll(ctx context.Context) ([]*domain.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Return a copy of all posts
	result := make([]*domain.Post, len(r.posts))
	copy(result, r.posts)
	return result, nil
}
