package handler

import (
	"go-clean-microblog/internal/interface_adapter/web/presenter"
	"go-clean-microblog/internal/usecase"
	"go-clean-microblog/internal/usecase/createpost"
	"net/http"
)

type CreatePostHandler struct {
	createPost createpost.Interaction
}

func NewCreatePostHandler(createPost createpost.Interaction) *CreatePostHandler {
	return &CreatePostHandler{
		createPost: createPost,
	}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	input := createpost.Input{
		Content: r.FormValue("content"),
	}
	presenterCtx := presenter.NewContext(w, r)
	context := usecase.Context{
		Ctx:          r.Context(),
		PresenterCtx: presenterCtx,
	}

	h.createPost.Execute(context, input)

	presenterCtx.RespondWithJSON()
}
