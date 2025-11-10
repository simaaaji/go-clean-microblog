package listposts

import (
	"go-clean-microblog/internal/domain"
	"go-clean-microblog/internal/usecase"
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

func (i *Interactor) Execute(ctx usecase.Context) error {
	posts, err := i.repo.FindAll(ctx.Ctx)
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

	return i.presenter.Present(ctx.PresenterCtx, output)
}
