package controllers

import (
	"net/http"
	"text/template"

	"../model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	medicoes := model.BuscaTodasMedicao()
	temp.ExecuteTemplate(w, "Index", medicoes)

}
