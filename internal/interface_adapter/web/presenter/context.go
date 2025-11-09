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

func (c *Context) RespondWithJSON() error {
	jsonResp, err := json.Marshal(c.ResponseData)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	if c.StatusCode == 0 {
		log.Fatal("StatusCode is not set")
	}

	c.ResponseWriter.WriteHeader(c.StatusCode)
	_, err = c.ResponseWriter.Write(jsonResp)

	return err
}
