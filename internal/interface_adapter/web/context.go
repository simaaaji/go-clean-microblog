package web

import (
	ctx "context"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"net/http"
)

type Context interface {
	ctx.Context
	ViewModel() viewmodel.ViewModel
	SetViewModel(vm viewmodel.ViewModel)
}

type context struct {
	ctx.Context
	// ウェブの場合、ハンドラでレスポンスを生成する際にViewModelを参照するため、ContextにViewModelを保持させる
	viewModel viewmodel.ViewModel
}

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return &context{
		Context: r.Context(),
	}
}

func (c *context) ViewModel() viewmodel.ViewModel {
	return c.viewModel
}

func (c *context) SetViewModel(vm viewmodel.ViewModel) {
	c.viewModel = vm
}
