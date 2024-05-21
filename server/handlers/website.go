package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/CodeMonkVJ/perusenews/server/data"
)

type Websites struct {
    l *log.Logger
}

func NewWebsites(l*log.Logger) *Websites {
    return &Websites{l}
}

func (w*Websites) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.getWebsites(rw, r)
        return
    }

    if r.Method == http.MethodPost {
        w.addWebsite(rw, r)
        return
    }

    if r.Method == http.MethodPut {
        rg := regexp.MustCompile(`/([0-9]+)`);
        g := rg.FindAllStringSubmatch(r.URL.Path, -1)
        
        if len(g) != 1 {
            http.Error(rw, "Invalid URI", http.StatusBadRequest)
            return
        }

        if len(g[0]) != 2 {
            http.Error(rw, "Invalid URI", http.StatusBadRequest)
            return
        }

        idString := g[0][1]
        id, err := strconv.Atoi(idString)
        if err != nil {
            http.Error(rw, "Invalid URI", http.StatusBadRequest)            
            return
        }
        
        w.updateWebsite(id, rw, r)
    }

    rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (w *Websites) getWebsites(rw http.ResponseWriter, r *http.Request) {
    lw := data.GetWebsites()
    err := lw.ToJSON(rw)
    if err != nil {
        http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
    }
}

func (w *Websites) addWebsite(rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle POST Website")

    web := &data.Website{}
    err := web.FromJSON(r.Body)
    if err != nil {
        http.Error(rw, "Unable to marshal json data", http.StatusBadRequest)
    }
    data.AddWebsite(web) 
}

func (w *Websites) updateWebsite(id int, rw http.ResponseWriter, r *http.Request) {
    w.l.Println("Handle PUT Website")
    
    web := &data.Website{}
    err := web.FromJSON(r.Body)
    if err != nil {
        w.l.Println("Unable to marshal json data")
        http.Error(rw, "Unable to marshal json data", http.StatusBadRequest)
        return
    }

    err = data.UpdateWebsite(id, web)
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
