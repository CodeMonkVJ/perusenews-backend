package handlers

import (
	"net/http"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

// swagger:route PUT /websites websites updateWebsite
// Update a websites details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update websites
func (w *Websites) UpdateWebsite(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")

    // fetch website from context
    web := r.Context().Value(KeyWebsite{}).(data.Website)
    w.l.Println("[DEBUG] updating record id", web.ID)
    
    err := data.UpdateWebsite(web)
    if err == data.ErrorWebsiteNotFound {
        w.l.Println("Website not found")
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: "Website not found in db"}, rw)
        return
    }
    
    // write the no content success Header
    rw.WriteHeader(http.StatusNoContent)
}
