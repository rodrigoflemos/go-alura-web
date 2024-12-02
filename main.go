package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Malha", Preco: 90, Quantidade: 5},
		{"Tenis", "Nike", 890, 3},
		{"Fone", "JBL", 590, 2},
		{"Celular", "Samsung", 1999.99, 100},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
