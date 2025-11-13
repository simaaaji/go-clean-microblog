package responder

import (
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"net/http"
)

type JSONResponder struct {
	BaseResponder
}

func NewJSONResponder(w http.ResponseWriter, r *http.Request) *JSONResponder {
	return &JSONResponder{
		BaseResponder{
			ResponseWriter: w,
			Request:        r,
			StatusCode:     0,
		},
	}
}

func (r *JSONResponder) Respond(viewModel viewmodel.ViewModel, args ...any) {
	jsonResp, err := r.vmToJSON(viewModel)
	if err != nil {
		log.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(r.ResponseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.StatusCode == 0 {
		log.Fatal("StatusCode is not set")
	}

	r.ResponseWriter.Header().Set("Content-Type", "application/json")
	r.ResponseWriter.WriteHeader(r.StatusCode)
	_, err = r.ResponseWriter.Write(jsonResp)
	if err != nil {
		log.Fatalf("error writing response. Err: %v", err)
	}
}
