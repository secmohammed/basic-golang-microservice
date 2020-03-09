package handlers

import (
    "build-microservice-with-go/products/models"
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
        http.Error(w, "Unable to marshal json", http.StatusBadRequest)
        return
    }
}

//Store function is used to create a new product.
func Store(response http.ResponseWriter, request *http.Request) {
    prod := &models.Product{}

    if err := prod.FromJSON(request.Body); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return
    }
    models.AddProduct(prod)
    if err := prod.ToJSON(response); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return

    }
}

// Update function is used to update a product.
func Update(response http.ResponseWriter, request *http.Request) {
    prod := &models.Product{}
    id, err := strconv.Atoi(mux.Vars(request)["id"])
    if err != nil {
        http.Error(response, "Unable to parse int from passed parameter", http.StatusBadRequest)
        return
    }

    if err := prod.FromJSON(request.Body); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return
    }
    if err := models.UpdateProduct(id, prod); err != nil {
        http.Error(response, err.Error(), http.StatusBadRequest)
        return

    }
    if err := prod.ToJSON(response); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return

    }

}
