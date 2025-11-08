package server

import (
	"go-clean-microblog/internal/interface_adapter/memory_store/repo"
	"go-clean-microblog/internal/interface_adapter/web/handler"
	"go-clean-microblog/internal/interface_adapter/web/presenter"
	"go-clean-microblog/internal/usecase/createpost"
)

var CreatePostHandler *handler.CreatePostHandler

// Inject dependencies for handlers
func init() {
	repo := repo.NewPostRepo()
	presenter := presenter.NewCreatePostPresenter()
	useCase := createpost.NewInteractor(repo, presenter)
	CreatePostHandler = handler.NewCreatePostHandler(useCase)
}
