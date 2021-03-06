// Package swagger of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package swagger

import (
    "build-microservice-with-go/products/models"
    "build-microservice-with-go/products/utils"
)

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
    // Description of the error
    // in: body
    Body utils.GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
    // Collection of the errors
    // in: body
    Body utils.ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
    // All current products
    // in: body
    Body []models.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
    // Newly created product
    // in: body
    Body models.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
    // Product data structure to Update or Create.
    // Note: the id field is ignored by update and create operations
    // in: body
    // required: true
    Body models.Product
}

// swagger:parameters updateProduct listSingleProduct deleteProduct
type productIDParamsWrapper struct {
    // The id of the product for which the operation relates
    // in: path
    // required: true
    ID int `json:"id"`
}
