package createpost

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

func (i *Interactor) Execute(ctx usecase.Context, input Input) error {
	post := &domain.Post{
		Content: input.Content,
	}

	if err := i.repo.Create(ctx.Ctx, post); err != nil {
		return err
	}

	return i.presenter.Present(ctx.PresenterCtx, &Output{
		ID:      post.ID,
		Content: post.Content,
	})
}
