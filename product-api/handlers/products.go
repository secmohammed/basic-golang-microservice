package handlers

import (
    "build-microservice-with-go/product-api/data"
    "log"
    "net/http"
)

// Products type.
type Products struct {
    logger *log.Logger
}

//NewProducts function is used create a new function.
func NewProducts(logger *log.Logger) *Products {
    return &Products{logger}
}
func (products *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    lp := data.GetProducts()
    err := lp.ToJSON(w)
    if err != nil {
        http.Error(w, "Unable to marshal json", http.StatusBadRequest)
        return
    }

}
