package responder

import (
	"encoding/json"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"maps"
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
			ViewModels:     []viewmodel.ViewModel{},
		},
	}
}

func (c *BaseResponder) Respond() {
	merged := make(map[string]any)
	for _, vm := range c.ViewModels {
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

	jsonResp, err := json.Marshal(merged)
	if err != nil {
		log.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(c.ResponseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if c.StatusCode == 0 {
		log.Fatal("StatusCode is not set")
	}

	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(c.StatusCode)
	_, err = c.ResponseWriter.Write(jsonResp)
	if err != nil {
		log.Fatalf("error writing response. Err: %v", err)
	}
}
