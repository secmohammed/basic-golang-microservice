package models

import (
    "encoding/json"
    "fmt"
    "io"
    "time"
)

// Product type.
type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    SKU         string  `json:"sku"`
    CreatedAt   string  `json:"created_at"`
    UpdatedAt   string  `json:"updated_at"`
    DeletedAt   string  `json:"deleted_at,omitempty"`
}

var errProductNotFound = fmt.Errorf("Product not found")
var products = []*Product{
    {
        ID:          1,
        Name:        "Latte",
        Description: "Frothy milky coffee",
        Price:       2.45,
        SKU:         "abc3232",
        CreatedAt:   time.Now().UTC().String(),
        UpdatedAt:   time.Now().UTC().String(),
    },
    {
        ID:          2,
        Name:        "Espresso",
        Description: "Short nd strong coffee without milk.",
        Price:       1.99,
        SKU:         "fjdka",
        CreatedAt:   time.Now().UTC().String(),
        UpdatedAt:   time.Now().UTC().String(),
    },
}

// Products are product colleciton.
type Products []*Product

// ToJSON method is used to transform the product to a json response.
func (p *Product) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}

// ToJSON method is used to transform the products to a json response.
func (p *Products) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}

// FromJSON method is used to decode a json.
func (p *Product) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(p)

}

func getNextID() int {
    product := products[len(products)-1]
    return product.ID + 1
}

// GetProducts function is used to get the products.
func GetProducts() Products {
    return products
}

// AddProduct function is used to append a new product to the collection.
func AddProduct(product *Product) {
    product.ID = getNextID()
    products = append(products, product)
}

func findProduct(id int) (*Product, int, error) {
    for key, product := range products {
        if product.ID == id {
            return product, key, nil
        }
    }
    return nil, -1, errProductNotFound
}

// UpdateProduct function is used to update a product.
func UpdateProduct(id int, product *Product) error {
    _, key, err := findProduct(id)
    if err != nil {
        return err
    }
    product.ID = id
    products[key] = product
    return nil
}
