package presenter

import (
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"go-clean-microblog/internal/usecase"
	"go-clean-microblog/internal/usecase/createpost"
)

type CreatePostPresenter struct{}

func NewCreatePostPresenter() *CreatePostPresenter {
	return &CreatePostPresenter{}
}

func (p *CreatePostPresenter) Present(ctx usecase.PresenterContext, output *createpost.Output) error {
	c := ctx.(*Context)

	c.AddViewModel(&viewmodel.CreatePost{
		Post: &viewmodel.Post{
			ID:      output.ID,
			Content: output.Content,
		},
	})

	return nil
}
