package models

import "testing"

func TestChecksValidation(t *testing.T) {
    p := &Product{
        Name:        "latte",
        Price:       1.00,
        Description: "Hello",
        SKU:         "abs-abc-def",
    }
    err := p.Validate()
    if err != nil {
        t.Fatal(err.Error())
    }
}
