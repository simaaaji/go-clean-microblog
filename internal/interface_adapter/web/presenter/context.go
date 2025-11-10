package presenter

import (
	"encoding/json"
	"go-clean-microblog/internal/interface_adapter/web/viewmodel"
	"log"
	"net/http"

	"github.com/romsar/gonertia"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	StatusCode     int
	// ウェブでは複数のユースケース利用の結果を最終的に一つのレスポンスとしてまとめて返す必要があるため、複数のViewModelを保持できるようにする
	ViewModels []viewmodel.ViewModel
	// Inertia instance for rendering Inertia.js responses
	Inertia *gonertia.Inertia
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		StatusCode:     0,
		ViewModels:     []viewmodel.ViewModel{},
		Inertia:        nil,
	}
}

func NewContextWithInertia(w http.ResponseWriter, r *http.Request, inertia *gonertia.Inertia) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		StatusCode:     0,
		ViewModels:     []viewmodel.ViewModel{},
		Inertia:        inertia,
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
		var vmMap map[string]interface{}
		if err := json.Unmarshal(vmJSON, &vmMap); err != nil {
			log.Printf("error unmarshaling viewmodel to map. Err: %v", err)
			continue
		}
		for k, v := range vmMap {
			merged[k] = v
		}
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

func (c *Context) RespondWithInertia(component string) {
	if c.Inertia == nil {
		log.Printf("Inertia is not initialized")
		http.Error(c.ResponseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Merge all ViewModels into a single props map
	props := make(gonertia.Props)
	for _, vm := range c.ViewModels {
		vmJSON, err := json.Marshal(vm)
		if err != nil {
			log.Printf("error marshaling viewmodel. Err: %v", err)
			continue
		}
		var vmMap map[string]interface{}
		if err := json.Unmarshal(vmJSON, &vmMap); err != nil {
			log.Printf("error unmarshaling viewmodel to map. Err: %v", err)
			continue
		}
		for k, v := range vmMap {
			props[k] = v
		}
	}

	err := c.Inertia.Render(c.ResponseWriter, c.Request, component, props)
	if err != nil {
		log.Printf("error rendering with Inertia. Err: %v", err)
		http.Error(c.ResponseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
