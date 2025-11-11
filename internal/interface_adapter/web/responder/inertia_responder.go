package responder

import (
	"encoding/json"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"maps"
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
			ViewModels:     []viewmodel.ViewModel{},
		},
	}
}

func (r *InertiaResponder) Respond(args ...any) {
	if len(args) == 0 {
		log.Fatal("InertiaResponder.Respond requires at least one argument: ComponentName")
	}
	componentName, ok := args[0].(string)
	if !ok {
		log.Fatal("First argument to InertiaResponder.Respond must be a string representing ComponentName")
	}

	merged := make(map[string]any)
	for _, vm := range r.ViewModels {
		vmJSON, err := json.Marshal(vm)
		if err != nil {
			log.Printf("error marshaling viewmodel. Err: %v", err)
			continue
		}
		var vmMap map[string]any
		if err := json.Unmarshal(vmJSON, &vmMap); err != nil {
			log.Printf("error unmarshaling viewmodel to map. Err: %v", err)
			continue
		}
		maps.Copy(merged, vmMap)
	}

	if r.StatusCode == 0 {
		log.Fatal("StatusCode is not set")
	}

	inertiaManager := r.Request.Context().Value(InertiaManagerCtxKey{}).(*inertia.Inertia)
	if inertiaManager == nil {
		log.Fatal("InertiaManager is not set in context")
	}

	err := inertiaManager.Render(r.ResponseWriter, r.Request, componentName, merged)

	if err != nil {
		log.Fatalf("error writing response. Err: %v", err)
	}
}
