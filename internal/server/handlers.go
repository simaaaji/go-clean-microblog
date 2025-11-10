package server

import (
	"go-clean-microblog/internal/interface_adapter/memory_store/repo"
	"go-clean-microblog/internal/interface_adapter/web/handler"
	"go-clean-microblog/internal/interface_adapter/web/presenter"
	"go-clean-microblog/internal/usecase/createpost"
	"go-clean-microblog/internal/usecase/listposts"
)

// Inject dependencies for handlers
var (
	CreatePostHandler *handler.CreatePostHandler
	ListPostsHandler  *handler.ListPostsHandler
)

func init() {
	r := repo.NewPostRepo()
	CreatePostHandler = newCreatePostHander(r)
	ListPostsHandler = newListPostHander(r)
}

func newCreatePostHander(r *repo.PostRepo) *handler.CreatePostHandler {
	p := presenter.NewCreatePostPresenter()
	uc := createpost.NewInteractor(r, p)
	return handler.NewCreatePostHandler(uc)
}

func newListPostHander(r *repo.PostRepo) *handler.ListPostsHandler {
	p := presenter.NewListPostsPresenter()
	uc := listposts.NewInteractor(r, p)
	return handler.NewListPostsHandler(uc)
}
