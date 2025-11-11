package listposts

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

func (i *Interactor) Execute(ctx context.Context) error {
	posts, err := i.repo.FindAll(ctx)
	if err != nil {
		return err
	}

	output := &Output{
		Posts: make([]PostOutput, len(posts)),
	}

	for idx, post := range posts {
		output.Posts[idx] = PostOutput{
			ID:      post.ID,
			Content: post.Content,
		}
	}

	return i.presenter.Present(ctx, output)
}
