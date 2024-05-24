package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

    "github.com/gorilla/mux"
	"github.com/CodeMonkVJ/perusenews/server/handlers"
)

func main() {
    l := log.New(os.Stdout, "perusenews-api", log.LstdFlags)
    wh := handlers.NewWebsites(l)
    
    sm := mux.NewRouter()

    getRouter := sm.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/", wh.GetWebsites)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", wh.UpdateWebsite)
    putRouter.Use(wh.MiddlewareWebsiteValidation)

    postRouter := sm.Methods(http.MethodPost).Subrouter()
    postRouter.HandleFunc("/", wh.AddWebsite)
    postRouter.Use(wh.MiddlewareWebsiteValidation)
    
    s := &http.Server{
        Addr: ":9090",
        Handler: sm,
        IdleTimeout: 120*time.Second,
        WriteTimeout: 1*time.Second,
        ReadTimeout: 1*time.Second,
    }

    go func() {
        err := s.ListenAndServe()
        if err != nil {
            l.Fatal(err)
        }
    }()

    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <-sigChan
    l.Println("Received terminate, shutting down", sig)

    tc, _ := context.WithTimeout(context.Background(), time.Second*30)
    err := s.Shutdown(tc)
    if err != nil {
        l.Println("Shutdown not successful")
        l.Fatal(err)
    }
}
