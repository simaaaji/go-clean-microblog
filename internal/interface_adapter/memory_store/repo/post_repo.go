package repo

import (
	"context"
	"go-clean-microblog/internal/domain"
)

type PostRepo struct {
	posts []*domain.Post
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		posts: []*domain.Post{},
	}
}

func (r *PostRepo) Create(ctx context.Context, post *domain.Post) error {
	// In-memory implementation: just assign a dummy ID
	post.ID = 1
	r.posts = append(r.posts, post)
	return nil
}
