package responder

import (
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"net/http"
)

type Responder interface {
	SetStatusCode(code int)
	AddViewModel(vm viewmodel.ViewModel)
	Respond()
}

type BaseResponder struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	StatusCode     int
	// ウェブでは複数のユースケース利用の結果を最終的に一つのレスポンスとしてまとめて返す必要があるため、複数のViewModelを保持できるようにする
	ViewModels []viewmodel.ViewModel
}

func (r *BaseResponder) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r *BaseResponder) AddViewModel(vm viewmodel.ViewModel) {
	r.ViewModels = append(r.ViewModels, vm)
}
