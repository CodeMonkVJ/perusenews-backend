// Package classification of Website API
//
// Documentation of Website API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CodeMonkVJ/perusenews/server/data"
	"github.com/gorilla/mux"
)

type KeyWebsite struct{}

type Websites struct {
    l *log.Logger
    v *data.Validation
}

func NewWebsites(l*log.Logger, v *data.Validation) *Websites {
    return &Websites{l, v}
}

var ErrInvalidWebsitePath = fmt.Errorf("Invalid path, path should be /websites/[id]")

type GenericError struct {
    Message string `json:"message"`
}

type ValidationError struct {
    Messages []string `json:"messages"`
}

// getWebsiteID returns the website ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getWebsiteID(r *http.Request) int {
    vars := mux.Vars(r)

    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        panic(err)
    }

    return id
}

