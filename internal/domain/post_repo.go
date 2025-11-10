package domain

import "context"

type PostRepo interface {
	Create(ctx context.Context, post *Post) error
}
