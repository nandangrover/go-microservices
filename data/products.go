package data

import (
	"encoding/json"
	"fmt"
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

// ToJSON is a Method on type Products (slice of Product), used to covert structure to JSON
func (p *Products) ToJSON(w io.Writer) error {
	// NewEncoder requires an io.Reader. http.ResponseWriter is the same thing
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// FromJSON is a Method on type Products (slice of Product)
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

//GetProducts - Return the product list
func GetProducts() Products {
	return productList
}

//AddProduct - Add the product to our struct Product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//UpdateProduct - Updates the product to our struct Product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// ErrProductNotFound is the Standard Product not found error structure
var ErrProductNotFound = fmt.Errorf("Product not found")

// Increments the Product ID by one
func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
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
