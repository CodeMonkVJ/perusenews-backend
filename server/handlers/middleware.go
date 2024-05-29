package handlers

import (
	"context"
	"net/http"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

// MiddlewareWebsiteValidation validates the product in the request and calls next if ok
func (w *Websites) MiddlewareWebsiteValidation(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        rw.Header().Add("Content-Type", "application/json")
        web := &data.Website{}
        err := data.FromJSON(web, r.Body)
        if err != nil {
            w.l.Println("[ERROR] deserializing website data", err)
            rw.WriteHeader(http.StatusBadRequest)
            data.ToJSON(&GenericError{Message: err.Error()}, rw) 
            return
        }

        errs := w.v.Validate(web)

        if len(errs) != 0 {
            w.l.Println("[ERROR] validating website", errs)
            rw.WriteHeader(http.StatusUnprocessableEntity)
            data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
            return
        }

        ctx := context.WithValue(r.Context(), KeyWebsite{}, *web)
        r = r.WithContext(ctx)

        next.ServeHTTP(rw, r)
    })
}
