package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CodeMonkVJ/perusenews/server/data"
	"github.com/CodeMonkVJ/perusenews/server/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
    l := log.New(os.Stdout, "perusenews-api", log.LstdFlags)
    v := data.NewValidation()

    wh := handlers.NewWebsites(l,v)
    
    sm := mux.NewRouter()

    getRouter := sm.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/websites", wh.ListAll)
    getRouter.HandleFunc("/websites/{id:[0-9]+}", wh.ListSingle)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/websites", wh.UpdateWebsite)
    putRouter.Use(wh.MiddlewareWebsiteValidation)

    postRouter := sm.Methods(http.MethodPost).Subrouter()
    postRouter.HandleFunc("/websites", wh.Add)
    postRouter.Use(wh.MiddlewareWebsiteValidation)

    deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouter.HandleFunc("/websites/{id:[0-9]+}", wh.DeleteWebsite)

    opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
    sh := middleware.Redoc(opts, nil)

    getRouter.Handle("/docs", sh)
    getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))
    
    ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

    s := &http.Server{
        Addr: ":9090",
        Handler: ch(sm),
        IdleTimeout: 120*time.Second,
        WriteTimeout: 15*time.Second,
        ReadTimeout: 20*time.Second,
    }

    go func() {
        l.Println("Starting server on port 9090")

        err := s.ListenAndServe()
        if err != nil {
            l.Printf("Error starting server: %s\n", err)
            os.Exit(1)
        }
    }()

    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <-sigChan
    log.Println("Got signal:", sig)

    tc, _ := context.WithTimeout(context.Background(), time.Second*30)
    s.Shutdown(tc)
}
