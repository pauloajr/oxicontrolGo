package model

import (
	"log"

	bd "../db"
)

type Medicao struct {
	Id              int
	DataHora        string
	Batimentos      int
	NivelOxigenacao int
}

func BuscaTodasMedicao() []Medicao {
	db := bd.ConectarBD()
	defer db.Close()

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

	return medicoes
}
