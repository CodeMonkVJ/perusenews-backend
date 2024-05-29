package handlers

import (
	"net/http"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

// swagger:route GET /websites websites listWebsites
// Returns a list of websites
// responses:
//  200: websitesResponse

// ListAll handles GET requests and returns all current websites
func (w *Websites) ListAll(rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle GET Websites")
    rw.Header().Add("Content-Type", "application/json")

    lw := data.GetWebsites()

    err := data.ToJSON(lw, rw)
    if err != nil {
        w.l.Println("[ERROR] serializing website", err)
    }
}

// swagger:route GET /websites/[id] websites listSingleWebsite
// Return a list of websites from the db
// responses:
//   200: websitesResponse
//   404: errorResponse

// ListSingle handles GET requests
func (w *Websites) ListSingle(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")

    id := getWebsiteID(r)

    w.l.Println("[DEBUG] get record id", id)

    web, err := data.GetWebsiteByID(id)

    switch err {
        case nil:

        case data.ErrorWebsiteNotFound:
            w.l.Println("[ERROR] fetching website", err)

            rw.WriteHeader(http.StatusNotFound)
            data.ToJSON(&GenericError{Message: err.Error()}, rw)
            return

        default:
            w.l.Println("[ERROR] fetching website", err)

            rw.WriteHeader(http.StatusInternalServerError)
            data.ToJSON(&GenericError{Message: err.Error()}, rw)
            return
    }

    err = data.ToJSON(web, rw)

    if err != nil {
        w.l.Println("[ERROR] serializing website", err)
    }
}
