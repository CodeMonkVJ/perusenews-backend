package handlers

import (
	"net/http"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

// swagger:route DELETE /websites/{id} websites deleteWebsite
// Returns a list of websites
// responses:
//  201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteWebsites handles DELETE requests and removes the websites from the data store
func (w *Websites) DeleteWebsite(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")
    id := getWebsiteID(r)
    w.l.Println("[DEBUG] deleting record id", id)
    
    err := data.DeleteWebsite(id)
    if err == data.ErrorWebsiteNotFound {
        w.l.Println("[ERROR] deleting record id does not exist")

        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }

    if err != nil {
        w.l.Println("[ERROR] deleting record", err)

        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }

    rw.WriteHeader(http.StatusNoContent)
}
