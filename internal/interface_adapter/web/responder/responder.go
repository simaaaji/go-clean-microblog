package responder

import (
	"encoding/json"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"net/http"
)

type Responder interface {
	SetStatusCode(code int)
	Respond(viewModel viewmodel.ViewModel, args ...any)
	IsInertiaResponder() bool
}

type BaseResponder struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	StatusCode     int
}

func NewResponderByContentType(w http.ResponseWriter, r *http.Request) Responder {
	contentType := r.Header.Get("Content-Type")
	accept := r.Header.Get("Accept")

	// Inertiaリクエストかチェック
	if r.Header.Get("X-Inertia") != "" {
		return NewInertiaResponder(w, r)
	}

	// Content-TypeまたはAcceptヘッダーでJSON形式が指定されている場合
	if contentType == "application/json" || accept == "application/json" {
		return NewJSONResponder(w, r)
	}

	// デフォルトはInertiaResponder
	return NewInertiaResponder(w, r)
}

func (r *BaseResponder) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r *BaseResponder) IsInertiaResponder() bool {
	return false
}

func (r *BaseResponder) vmToJSON(vm viewmodel.ViewModel) ([]byte, error) {
	data, err := json.Marshal(vm)
	if err != nil {
		log.Printf("error marshaling viewmodel to JSON. Err: %v", err)
		return nil, err
	}
	return data, nil
}

func (r *BaseResponder) vmToMap(vm viewmodel.ViewModel) map[string]any {
	vmMap := make(map[string]any)
	vmJSON, err := r.vmToJSON(vm)
	if err != nil {
		log.Printf("error marshaling viewmodel. Err: %v", err)
	} else {
		if err := json.Unmarshal(vmJSON, &vmMap); err != nil {
			log.Printf("error unmarshaling viewmodel to map. Err: %v", err)
		}
	}
	return vmMap
}
