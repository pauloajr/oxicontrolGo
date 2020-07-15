package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func ConectarBD() *sql.DB {
	conexao := "user=postgres dbname=oxidb password=yourpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic("Nao conseguiu conectar com o banco de dados." + err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Medicao struct {
	Id              int
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
	db := ConectarBD()
	defer db.Close()
	/*medicoes := []Medicao{
		{DataHora: "14/07/2020 18:06",
			Batimentos:      69,
			NivelOxigenacao: 97},
		{"14/07/2020 18:08", 69, 98},
	}*/

	selecionarMedicoes, err := db.Query("SELECT * FROM medicoes")
	if err != nil {
		log.Fatal(err.Error())
	}

	m := Medicao{}
	medicoes := []Medicao{}

	for selecionarMedicoes.Next() {
		var id int
		var batimentos, nivelOxigenacao int
		var dataHora string

		err = selecionarMedicoes.Scan(&id, &dataHora, &batimentos, &nivelOxigenacao)
		if err != nil {
			panic(err.Error())
		}

		m.Id = id
		m.DataHora = dataHora
		m.Batimentos = batimentos
		m.NivelOxigenacao = nivelOxigenacao

		medicoes = append(medicoes, m)
	}

	temp.ExecuteTemplate(w, "Index", medicoes)

	defer db.Close()
}
