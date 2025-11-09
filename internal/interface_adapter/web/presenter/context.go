package presenter

import (
	"encoding/json"
	"log"
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	StatusCode     int
	ResponseData   map[string]any
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		StatusCode:     0,
		ResponseData:   make(map[string]any),
	}
}

func (c *Context) Set(key string, value any) {
	c.ResponseData[key] = value
}

func (c *Context) RespondWithJSON() {
	jsonResp, err := json.Marshal(c.ResponseData)
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
