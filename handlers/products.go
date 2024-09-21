package handlers

import (
	"log"
	"net/http"

	"github.com/harshgupta9473/microservicINGO/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodGet{
		p.getProducts(w,r)
	}
	// handle 

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products)getProducts(w http.ResponseWriter,r *http.Request){
	lp := data.GetProducts()
	// d,err:=json.Marshal(lp)
	err :=lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	// w.Write(d)
}