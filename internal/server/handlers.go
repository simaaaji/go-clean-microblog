package server

import (
	"go-clean-microblog/internal/interface_adapter/memory_store/repo"
	"go-clean-microblog/internal/interface_adapter/web/handler"
	"go-clean-microblog/internal/interface_adapter/web/presenter"
	"go-clean-microblog/internal/usecase"
	"go-clean-microblog/internal/usecase/createpost"
	"go-clean-microblog/internal/usecase/listposts"
	"net/http"
)

// Inject dependencies for handlers
var (
	CreatePostHandler *handler.CreatePostHandler
	ListPostsHandler  *handler.ListPostsHandler
	sharedPostRepo    *repo.PostRepo
)

func init() {
	sharedPostRepo = repo.NewPostRepo()
	CreatePostHandler = newCreatePostHandler(sharedPostRepo)
	ListPostsHandler = newListPostsHandler(sharedPostRepo)
}

func newCreatePostHandler(r *repo.PostRepo) *handler.CreatePostHandler {
	p := presenter.NewCreatePostPresenter()
	uc := createpost.NewInteractor(r, p)
	return handler.NewCreatePostHandler(uc)
}

func newListPostsHandler(r *repo.PostRepo) *handler.ListPostsHandler {
	p := presenter.NewListPostsPresenter()
	uc := listposts.NewInteractor(r, p)
	return handler.NewListPostsHandler(uc)
}

// IndexHandler handles the index page with Inertia
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	listPostsPresenter := presenter.NewListPostsPresenter()
	listPostsUsecase := listposts.NewInteractor(sharedPostRepo, listPostsPresenter)

	presenterCtx := presenter.NewContextWithInertia(w, r, s.inertia)
	context := usecase.Context{
		Ctx:          r.Context(),
		PresenterCtx: presenterCtx,
	}

	if err := listPostsUsecase.Execute(context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	presenterCtx.RespondWithInertia("Index")
}
