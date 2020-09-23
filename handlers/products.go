package handlers

import (
	"log"
	"net/http"

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
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	listOfProducts := data.GetProducts()
	// Use encoder as it is marginally faster than json.marshal. It's important when we use multiple threads
	// d, err := json.Marshal(listOfProducts)
	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
