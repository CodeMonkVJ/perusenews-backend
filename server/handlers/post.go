package handlers

import (
	"net/http"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

// swagger:route POST /websites websites addWebsite
// Create a new websites
//
// responses:
//	200: websiteResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new websites
func (w *Websites) Add(rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle POST Website")
    
    web := r.Context().Value(KeyWebsite{}).(data.Website)

    w.l.Printf("[DEBUG] Inserting website: %#v\n", web)
    data.AddWebsite(web) 
}
