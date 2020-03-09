package main

import (
    "build-microservice-with-go/products/routes"
    "context"
    "log"
    "os"

    _ "github.com/joho/godotenv/autoload"

    "os/signal"
    "time"
)

func main() {
    server := routes.RegisterAPIRoutes()
    go func() {
        err := server.ListenAndServe()
        if err != nil {
            log.Fatal(err)
        }
    }()
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)
    sig := <-sigChan
    log.Println("received terminate, graceful shutdown", sig)
    tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    server.Shutdown(tc)

}
