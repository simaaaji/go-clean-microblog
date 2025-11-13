package handler

import (
	"go-clean-microblog/internal/interface_adapter/web"
	"go-clean-microblog/internal/interface_adapter/web/responder"
	"go-clean-microblog/internal/usecase/listposts"
	"net/http"
)

type ListPostsHandler struct {
	listPosts listposts.UseCase
}

func NewListPostsHandler(listPosts listposts.UseCase) *ListPostsHandler {
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

	re := responder.NewResponderByContentType(w, r)
	re.SetStatusCode(http.StatusOK)
	if re.IsInertiaResponder() {
		re.Respond(context.ViewModel(), "PostsIndex")
	} else {
		re.Respond(context.ViewModel())
	}
}
