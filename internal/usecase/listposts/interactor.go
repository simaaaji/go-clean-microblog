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

	// Limit to most recent 20 posts
	maxPosts := 20
	if len(posts) > maxPosts {
		posts = posts[len(posts)-maxPosts:]
	}

	// Reverse to show most recent first
	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
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
