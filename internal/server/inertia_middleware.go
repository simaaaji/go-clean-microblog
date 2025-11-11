package server

import (
	"context"
	"go-clean-microblog/internal/interface_adapter/web/responder"
	"net/http"

	inertia "github.com/petaki/inertia-go"
)

func InertiaMiddleware(next http.Handler) http.Handler {
	// ミドルウェア登録時にコールされる初期化処理
	// TODO: 置き場所はここで良いか？
	url := "http://localhost:8080"         // Application URL for redirect
	rootTemplate := "./inertia/app.gohtml" // Root template, see the example below
	version := ""                          // Asset version

	inertiaManager := inertia.New(url, rootTemplate, version)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), responder.InertiaManagerCtxKey{}, inertiaManager)

		nnext := inertiaManager.Middleware(next)
		nnext.ServeHTTP(w, r.WithContext(ctx))
	})
}
