package presenter

import (
	"encoding/json"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"maps"
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	StatusCode     int
	// ウェブでは複数のユースケース利用の結果を最終的に一つのレスポンスとしてまとめて返す必要があるため、複数のViewModelを保持できるようにする
	ViewModels []viewmodel.ViewModel
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		StatusCode:     0,
		ViewModels:     []viewmodel.ViewModel{},
	}
}

func (c *Context) AddViewModel(vm viewmodel.ViewModel) {
	c.ViewModels = append(c.ViewModels, vm)
}

func (c *Context) RespondWithJSON() {
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
