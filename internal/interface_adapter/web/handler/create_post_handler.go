package handler

import (
	"go-clean-microblog/internal/interface_adapter/web"
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
	context := web.NewContext(w, r)

	if err := h.createPost.Execute(context, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	context.Responder().SetStatusCode(http.StatusCreated)
	context.Responder().Respond()
}
