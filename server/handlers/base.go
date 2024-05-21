package handlers

import (
    "net/http"
    "fmt"
	"io"
	"log"
)

type Hello struct {
    l *log.Logger    
}

func NewHello(l *log.Logger) *Hello {
    return &Hello{l}
}

func (h*Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    h.l.Println("Received request")
    d, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(rw, "Error", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(rw, "hello %s", d)
}
