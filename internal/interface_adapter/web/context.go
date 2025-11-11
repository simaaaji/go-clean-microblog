package web

import (
	ctx "context"
	"go-clean-microblog/internal/interface_adapter/web/responder"
	"net/http"
)

type Context interface {
	ctx.Context
	Responder() *responder.Responder
}

type context struct {
	ctx.Context
	responder responder.Responder
}

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	responder := responderByContentType(w, r)

	return &context{
		Context:   r.Context(),
		responder: responder,
	}
}

func responderByContentType(w http.ResponseWriter, r *http.Request) responder.Responder {
	// 今後Content-Typeに応じて異なるResponderを返すようにする場合はここで実装する
	// return responder.NewJSONResponder(w, r)
	return responder.NewInertiaResponder(w, r)
}

func (c *context) Responder() *responder.Responder {
	return &c.responder
}
