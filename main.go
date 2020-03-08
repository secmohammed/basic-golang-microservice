package main

import (
    "build-microservice-with-go/product-api/handlers"
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    // router := mux.NewRouter()
    hh := handlers.NewProducts(l)
    sm := http.NewServeMux()
    sm.Handle("/", hh)
    server := &http.Server{
        Addr:         ":8000",
        Handler:      sm,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1 * time.Second,
        WriteTimeout: 1 * time.Second,
    }
    // router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

    // })
    go func() {
        err := server.ListenAndServe()
        if err != nil {
            l.Fatal(err)
        }
    }()
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)
    sig := <-sigChan
    l.Println("received terminate, graceful shutdown", sig)
    tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    // whenever we shutdown the server, it won't accept any other requests, but it will process the requests it received and once finished, it will shutdown.will
    server.Shutdown(tc)
}
