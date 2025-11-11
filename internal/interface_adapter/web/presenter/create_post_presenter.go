package presenter

import (
	"context"
	"go-clean-microblog/internal/interface_adapter/web"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"go-clean-microblog/internal/usecase/createpost"
)

type createPostPresenter struct{}

func NewCreatePostPresenter() *createPostPresenter {
	return &createPostPresenter{}
}

func (p *createPostPresenter) Present(ctx context.Context, output *createpost.Output) error {
	c := ctx.(web.Context)
	responder := *c.Responder()
	responder.AddViewModel(&viewmodel.CreatePost{
		Post: &viewmodel.Post{
			ID:      output.ID,
			Content: output.Content,
		},
	})

	return nil
}
