package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CodeMonkVJ/perusenews/server/handlers"
)

func main() {
    l := log.New(os.Stdout, "perusenews-api", log.LstdFlags)
    wh := handlers.NewWebsites(l)
    
    sm := http.NewServeMux()
    sm.Handle("/", wh)
    
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
