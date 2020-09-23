package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product defines the structure for an API product
//Since encoding/json is a package residing outside our package we need to uppercase the first character of the fields inside the structure
//To get nice json field names we can add struct tags though. This will output the key name as the tag name
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a type defining slice of struct Product
type Products []*Product

// ToJSON is a Method on type Products (slice of Product)
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

//GetProducts - Return the product list
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Description: "Latte",
		Name:        "Milky coffee",
		SKU:         "abc323",
		Price:       200,
		UpdatedOn:   time.Now().UTC().String(),
		CreatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Description: "Expresso",
		Name:        "Strong coffee",
		SKU:         "errfer",
		Price:       150,
		UpdatedOn:   time.Now().UTC().String(),
		CreatedOn:   time.Now().UTC().String(),
	},
}
