package models

import (
    "encoding/json"
    "fmt"
    "io"
    "regexp"
    "time"

    "github.com/go-playground/validator"
)

// Product type.
// swagger:model
type Product struct {
    // the id for this user.
    // required: true
    // min: 1
    ID int `json:"id"`
    // the name for this product.
    // required: true.
    Name string `json:"name" validate:"required"`
    // the description for  this product.
    // required: true.
    Description string `json:"description" validate:"required"`
    // the price for this product.
    // required: true.
    // greater than: 0
    Price float32 `json:"price" validate:"gt=0"`
    // the sku for this product.
    // it should be a valid sku format.
    // required: true.
    // example: abcdasdas-abcea-abcd
    SKU       string `json:"sku" validate:"required,sku"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at,omitempty"`
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

// Validate function is used to validate the struct of product.
func (p *Product) Validate() error {
    validate := validator.New()
    validate.RegisterValidation("sku", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
        matches := re.FindAllString(fl.Field().String(), -1)
        if len(matches) != 1 {
            return false
        }
        return true
    })
    return validate.Struct(p)
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

// DeleteProduct function is used to delete a product.
func DeleteProduct(id int) error {
    _, key, err := findProduct(id)
    if err != nil {
        return err
    }
    products = append(products[:key], products[key+1:]...)
    return nil
}
