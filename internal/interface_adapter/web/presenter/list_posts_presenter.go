package presenter

import (
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"go-clean-microblog/internal/usecase"
	"go-clean-microblog/internal/usecase/listposts"
)

type ListPostsPresenter struct{}

func NewListPostsPresenter() *ListPostsPresenter {
	return &ListPostsPresenter{}
}

func (p *ListPostsPresenter) Present(ctx usecase.PresenterContext, output *listposts.Output) error {
	c := ctx.(*Context)

	posts := make([]viewmodel.Post, len(output.Posts))
	for idx, post := range output.Posts {
		posts[idx] = viewmodel.Post{
			ID:      post.ID,
			Content: post.Content,
		}
	}

	c.AddViewModel(&viewmodel.ListPosts{
		Posts: posts,
	})

	return nil
}
