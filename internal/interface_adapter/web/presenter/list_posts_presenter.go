package presenter

import (
	"context"
	"go-clean-microblog/internal/interface_adapter/web"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"go-clean-microblog/internal/usecase/listposts"
)

type ListPostsPresenter interface {
	listposts.Presenter
}

type listPostsPresenter struct{}

func NewListPostsPresenter() *listPostsPresenter {
	return &listPostsPresenter{}
}

func (p *listPostsPresenter) Present(ctx context.Context, output *listposts.Output) error {
	c := ctx.(web.Context)

	posts := make([]viewmodel.Post, len(output.Posts))
	for idx, post := range output.Posts {
		posts[idx] = viewmodel.Post{
			ID:      post.ID,
			Content: post.Content,
		}
	}

	c.Responder().AddViewModel(&viewmodel.ListPosts{
		Posts: posts,
	})

	return nil
}
