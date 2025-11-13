package responder

import (
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"net/http"

	inertia "github.com/petaki/inertia-go"
)

type InertiaResponder struct {
	BaseResponder
}

type InertiaManagerCtxKey struct{}

func NewInertiaResponder(w http.ResponseWriter, r *http.Request) *InertiaResponder {
	return &InertiaResponder{
		BaseResponder{
			ResponseWriter: w,
			Request:        r,
			StatusCode:     0,
		},
	}
}

func (r *InertiaResponder) IsInertiaResponder() bool {
	return true
}

func (r *InertiaResponder) Respond(viewModel viewmodel.ViewModel, args ...any) {
	if len(args) == 0 {
		log.Fatal("InertiaResponder.Respond requires at least one argument: ComponentName")
	}
	componentName, ok := args[0].(string)
	if !ok {
		log.Fatal("First argument to InertiaResponder.Respond must be a string representing ComponentName")
	}

	vmMap := r.vmToMap(viewModel)

	if r.StatusCode == 0 {
		log.Fatal("StatusCode is not set")
	}

	inertiaManager := r.Request.Context().Value(InertiaManagerCtxKey{}).(*inertia.Inertia)
	if inertiaManager == nil {
		log.Fatal("InertiaManager is not set in context")
	}

	err := inertiaManager.Render(r.ResponseWriter, r.Request, componentName, vmMap)

	if err != nil {
		log.Fatalf("error writing response. Err: %v", err)
	}
}
