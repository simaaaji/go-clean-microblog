package createpost

import (
	"context"
	"go-clean-microblog/internal/domain"
)

type Interactor struct {
	repo      domain.PostRepo
	presenter Presenter
}

func NewInteractor(repo domain.PostRepo, presenter Presenter) *Interactor {
	return &Interactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *Interactor) Execute(ctx context.Context, input Input) error {
	post := &domain.Post{
		Content: input.Content,
	}

	if err := i.repo.Create(ctx, post); err != nil {
		return err
	}

	return i.presenter.Present(ctx, &Output{
		ID:      post.ID,
		Content: post.Content,
	})
}
