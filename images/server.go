package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

    gohandlers "github.com/gorilla/handlers"
	"github.com/CodeMonkVJ/perusenews/images/files"
	"github.com/CodeMonkVJ/perusenews/images/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
    l := hclog.New(
        &hclog.LoggerOptions{
            Name: "product-images",
            Level: hclog.LevelFromString("debug"),
        },
    )

    sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

    // max file size - 5MB
    stor, err := files.NewLocal("./imagestore", 1024*1000*5)
    if err != nil {
        l.Error("Unable to create storage", "error", err)
        os.Exit(1)
    }

    fh := handlers.NewFiles(stor, l)

    sm := mux.NewRouter()

    ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

    ph := sm.Methods(http.MethodPost).Subrouter()
    ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
    ph.HandleFunc("/", fh.UploadMultipart)

    gh := sm.Methods(http.MethodGet).Subrouter()
    gh.Handle(
        "/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
        http.StripPrefix("/images/", http.FileServer(http.Dir("./imagestore"))),
    )

    s := http.Server{
        Addr: ":9090",
        Handler: ch(sm),
        ErrorLog: sl,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout: 120 * time.Second,
    }

    go func() {
        l.Info("Starting server", "bind_address", ":9090")

        err := s.ListenAndServe()
        if err != nil {
            l.Error("Unable to start server", "error", err)
            os.Exit(1)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)

    sig := <-c
    l.Info("Shutting down server with", "signal", sig)

    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    s.Shutdown(ctx)
}


