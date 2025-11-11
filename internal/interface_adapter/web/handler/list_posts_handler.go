package handler

import (
	"go-clean-microblog/internal/interface_adapter/web"
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
	context := web.NewContext(w, r)

	if err := h.listPosts.Execute(context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	context.Responder().SetStatusCode(http.StatusOK)
	context.Responder().Respond()
}
