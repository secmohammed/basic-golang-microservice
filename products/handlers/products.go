// Package handlers of Product API
//
// Documntation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - appication/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
    "build-microservice-with-go/products/models"
    "build-microservice-with-go/products/utils"
    "context"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// A list of products return in the response.
// swagger:response productsResponse
type productsResponse struct {
    // All products in the system.
    // in: body
    Body []models.Product
}

// swagger:parameters deleteProduct  updateProduct
type productIDParameterWrapper struct {
    // the id of the product to delete/update from the database.
    //in: path
    //required: true
    ID int `json:"id"`
}

// returned product in the response.
// swagger:response productResponse
type productResponse struct {
    // All products in the system.
    // in: body
    Body models.Product
}

// Products type.
type Products struct {
    logger *log.Logger
}

//NewProducts function is used create a new function.
func NewProducts(logger *log.Logger) *Products {
    return &Products{logger}
}

// swagger:route GET /products products listProducts
// Returns a list of products.
//  responses:
//      200: productsResponse

// Index function is used to index the current products.
func Index(w http.ResponseWriter, r *http.Request) {
    lp := models.GetProducts()
    err := lp.ToJSON(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}

// swagger:route POST /products products storeProduct
// store product.
//  responses:
//      200: productResponse

//Store function is used to create a new product.
func Store(response http.ResponseWriter, request *http.Request) {
    prod := request.Context().Value(utils.KeyProduct{}).(models.Product)
    models.AddProduct(&prod)
    if err := prod.ToJSON(response); err != nil {
        http.Error(response, "Unable to marshal json", http.StatusBadRequest)
        return

    }
}

// swagger:route PUT /products/{id} products updateProduct
// update product.
//  responses:
//      200: productResponse

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

// swagger:route DELETE /products/{id} products deleteProduct
// delete product.

// DeleteProduct function is used to delete a product from the database.
func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    err := models.DeleteProduct(id)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusNotFound)
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
        // validate the product.
        if err := prod.Validate(); err != nil {
            http.Error(
                rw,
                fmt.Sprintf("Error validating product: %s", err),
                http.StatusUnprocessableEntity,
            )
            return
        }
        ctx := context.WithValue(r.Context(), utils.KeyProduct{}, prod)
        req := r.WithContext(ctx)
        next.ServeHTTP(rw, req)
    })

}
