package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Servidor rodando na porta 321")
	http.ListenAndServe(":321", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//temp.ExecuteTemplate(w, "index.html", nil)
	temp.ExecuteTemplate(w, "Index", nil)
}
