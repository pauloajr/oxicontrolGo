package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"../model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	medicoes := model.BuscaTodasMedicao()
	temp.ExecuteTemplate(w, "Index", medicoes)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bati := r.FormValue("bati-min")
		sp02 := r.FormValue("sp02")

		batiConvertInt, err := strconv.Atoi(bati)
		if err != nil {
			log.Println("bati nao conseguiu converter", err.Error())
		}
		sp02ConvertInt, err2 := strconv.Atoi(sp02)
		if err2 != nil {
			log.Println("sp02 nao conseguiu converter", err.Error())
		}

		data := time.Now()

		dataFinal := string(data.Format(("2006-01-02 15:04:05 ")))

		model.CriarNovaMedicao(dataFinal, batiConvertInt, sp02ConvertInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idMedicao := r.URL.Query().Get("id")
	model.DeletarMedicao(idMedicao)
	http.Redirect(w, r, "/", 301)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bati := r.FormValue("bati-min")
		sp02 := r.FormValue("sp02")
		id := r.FormValue("id")
		dthr := r.FormValue("dthr")

		batiConvertInt, err := strconv.Atoi(bati)
		if err != nil {
			log.Println("bati nao conseguiu converter", err.Error())
		}
		sp02ConvertInt, err2 := strconv.Atoi(sp02)
		if err2 != nil {
			log.Println("sp02 nao conseguiu converter", err.Error())
		}
		idConvertInt, err3 := strconv.Atoi(id)
		if err3 != nil {
			log.Println("id nao conseguiu converter", err.Error())
		}

		model.AtualizaProduto(idConvertInt, batiConvertInt, sp02ConvertInt, dthr)
	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idMedicao := r.URL.Query().Get("id")
	produto := model.BuscarProduto(idMedicao)
	temp.ExecuteTemplate(w, "Edit", produto)
}
