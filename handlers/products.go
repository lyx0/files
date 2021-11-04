package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("product handler")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "product")
}
