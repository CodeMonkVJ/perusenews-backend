package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CodeMonkVJ/perusenews/server/data"
	"github.com/gorilla/mux"
)

type Websites struct {
    l *log.Logger
}

func NewWebsites(l*log.Logger) *Websites {
    return &Websites{l}
}

func (w *Websites) GetWebsites(rw http.ResponseWriter, r *http.Request) {
    lw := data.GetWebsites()
    err := lw.ToJSON(rw)
    if err != nil {
        http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
    }
}

func (w *Websites) AddWebsite(rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle POST Website")
    
    web := r.Context().Value(KeyWebsite{}).(data.Website)
    data.AddWebsite(&web) 
}

func (w *Websites) UpdateWebsite(rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle PUT Website")
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(rw, "Unable to convert id", http.StatusBadRequest)
        return 
    }

    web := r.Context().Value(KeyWebsite{}).(data.Website)
    
    err = data.UpdateWebsite(id, &web)
    if err == data.ErrorWebsiteNotFound {
        w.l.Println("Website not found")
        http.Error(rw, "Website not found", http.StatusNotFound)
        return
    }

    if err != nil {
        w.l.Println("internal error ", err)
        http.Error(rw, "Website not found", http.StatusInternalServerError)
        return
    }
}

type KeyWebsite struct {}

func (w *Websites) MiddlewareWebsiteValidation(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        web := data.Website{}
        err := web.FromJSON(r.Body)
        if err != nil {
            w.l.Println("Unable to marshal json data")
            http.Error(rw, "Unable to marshal json data", http.StatusBadRequest)
            return
        }

        err = web.Validate()
        if err != nil {
            w.l.Println("[ERROR] validating website", err)
            http.Error(
                rw,
                fmt.Sprintf("Error validating product: %s", err),
                http.StatusBadRequest,
            )
            return
        }

        ctx := context.WithValue(r.Context(), KeyWebsite{}, web)
        req := r.WithContext(ctx)

        next.ServeHTTP(rw, req)
    })
}

