package handlers

import (
    "build-microservice-with-go/products/models"
    "build-microservice-with-go/products/utils"
    "context"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Products type.
type Products struct {
    logger *log.Logger
}

//NewProducts function is used create a new function.
func NewProducts(logger *log.Logger) *Products {
    return &Products{logger}
}

// Index function is used to index the current products.
func Index(w http.ResponseWriter, r *http.Request) {
    lp := models.GetProducts()
    err := lp.ToJSON(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}

//Store function is used to create a new product.
func Store(response http.ResponseWriter, request *http.Request) {
    prod := request.Context().Value(utils.KeyProduct{}).(models.Product)
    models.AddProduct(&prod)
    if err := prod.ToJSON(response); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return

    }
}

// Update function is used to update a product.
func Update(response http.ResponseWriter, request *http.Request) {
    id, err := strconv.Atoi(mux.Vars(request)["id"])
    if err != nil {
        http.Error(response, "Unable to parse int from passed parameter", http.StatusBadRequest)
        return
    }
    prod := request.Context().Value(utils.KeyProduct{}).(*models.Product)
    if err := models.UpdateProduct(id, prod); err != nil {
        http.Error(response, err.Error(), http.StatusBadRequest)
        return

    }
    if err := prod.ToJSON(response); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return

    }

}

// MiddlewareProductValidation is used to parse the product from body and insert it into context.
var MiddlewareProductValidation = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        prod := &models.Product{}

        if err := prod.FromJSON(r.Body); err != nil {
            http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
            return
        }
        ctx := context.WithValue(r.Context(), utils.KeyProduct{}, prod)
        req := r.WithContext(ctx)
        next.ServeHTTP(rw, req)
    })

}
