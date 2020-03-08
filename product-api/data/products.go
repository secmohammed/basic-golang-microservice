package data

import (
    "encoding/json"
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

// ToJSON method is used to transform the data to a json response.
func (p *Products) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}

// GetProducts function is used to get the products.
func GetProducts() Products {
    return products
}
