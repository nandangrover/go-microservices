package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler
type Hello struct {
	l *log.Logger
}

// NewHello Function which gives reference to Hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
