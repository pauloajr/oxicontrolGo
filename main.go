package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Medicao struct {
	DataHora        string
	Batimentos      int
	NivelOxigenacao int
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Servidor rodando na porta 321")
	http.ListenAndServe(":321", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	medicoes := []Medicao{
		{DataHora: "14/07/2020 18:06",
			Batimentos:      69,
			NivelOxigenacao: 97},
		{"14/07/2020 18:08", 69, 98},
	}
	temp.ExecuteTemplate(w, "Index", medicoes)
}
