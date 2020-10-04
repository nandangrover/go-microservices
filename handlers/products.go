package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/nandangrover/go-microservices/data"
)

//Products structure that holds a logger
type Products struct {
	l *log.Logger
}

// NewProducts function return the pointer to Products structure
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		// expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		// Ignore the error for now
		id, _ := strconv.Atoi(idString)

		p.updateProducts(id, rw, r)
	}
	// catch all other http verb with 405
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	listOfProducts := data.GetProducts()
	// Use encoder as it is marginally faster than json.marshal. It's important when we use multiple threads
	// d, err := json.Marshal(listOfProducts)
	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	// The reason why we use a buffer reader is so that we don't have to allocate all the memory instantly to a slice or something like that,
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	// p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put product")

	prod := &data.Product{}
	// The reason why we use a buffer reader is so that we don't have to allocate all the memory instantly to a slice or something like that,
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
