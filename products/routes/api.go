package routes

import (
    "build-microservice-with-go/products/handlers"
    "fmt"
    "net/http"
    "os"
    "time"

    "github.com/go-openapi/runtime/middleware"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// RegisterAPIRoutes is used to register the routes we need for the web application.
func RegisterAPIRoutes() *http.Server {
    docOptions := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
    sh := middleware.Redoc(docOptions, nil)
    router := mux.NewRouter()

    router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
    router.Handle("/docs", sh).Methods("GET")
    router.Path("/api/products").Handler(http.HandlerFunc(handlers.Index))
    productsPublicArea := router.PathPrefix("/api/products").Subrouter()
    productsPublicArea.Use(handlers.MiddlewareProductValidation)
    productsPublicArea.HandleFunc("", handlers.Store).Methods("POST")
    productsPublicArea.HandleFunc("/{id:[0-9]+}", handlers.Update).Methods("PUT")
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:" + os.Getenv("APP_PORT")},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowCredentials: true,
        // Enable Debugging for testing, consider disabling in production
        Debug: true,
    })
    // Insert the middleware
    handler := c.Handler(router)
    server := &http.Server{
        Addr:         ":" + os.Getenv("APP_PORT"),
        Handler:      handler,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1 * time.Second,
        WriteTimeout: 1 * time.Second,
    }
    fmt.Println(os.Getenv("APP_PORT"))
    return server
}
