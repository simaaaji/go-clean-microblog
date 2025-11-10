package handler

import (
	"go-clean-microblog/internal/interface_adapter/web/presenter"
	"go-clean-microblog/internal/usecase"
	"go-clean-microblog/internal/usecase/listposts"
	"net/http"
)

type ListPostsHandler struct {
	listPosts listposts.Interaction
}

func NewListPostsHandler(listPosts listposts.Interaction) *ListPostsHandler {
	return &ListPostsHandler{
		listPosts: listPosts,
	}
}

func (h *ListPostsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	presenterCtx := presenter.NewContext(w, r)
	context := usecase.Context{
		Ctx:          r.Context(),
		PresenterCtx: presenterCtx,
	}

	if err := h.listPosts.Execute(context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	presenterCtx.StatusCode = http.StatusOK
	presenterCtx.RespondWithJSON()
}
