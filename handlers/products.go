package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/harshgupta9473/microservicINGO/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect the id in the URI
		reg:= regexp.MustCompile(`/([0-9]+)`)
		g:=reg.FindAllStringSubmatch(r.URL.Path,-1)
		if len(g)!=1{
			http.Error(w,"Invalid URI",http.StatusBadRequest)
			return
		}
		if len(g[0])!=2{
			http.Error(w,"Invalid URI",http.StatusBadRequest)
			return
		}

		idString:=g[0][1]
		id,err:=strconv.Atoi(idString)
		if err!=nil{
			http.Error(w,"can't convert id string to int id",http.StatusInternalServerError)
			return
		}
		p.uodateProducts(id , w,r)
		return

		return

	}
	// handle

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	// d,err:=json.Marshal(lp)
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	// w.Write(d)
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	// p.l.Printf("Prod %#v",prod)
	data.AddProduct(prod)
}

func (p *Products)uodateProducts(id int,w http.ResponseWriter,r *http.Request){
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err=data.UpdateProduct(id, prod)
	if err==data.ErrProductNotFound{
		http.Error(w,"Product not found",http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(w,"Product not found",http.StatusInternalServerError)
	}
	return
}