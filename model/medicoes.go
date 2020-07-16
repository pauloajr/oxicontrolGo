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

func CriarNovaMedicao(dataHora string, bati, sp02 int) {
	db := bd.ConectarBD()

	insertDadosNoBanco, err := db.Prepare("INSERT INTO medicoes (datahora, batimentos, niveloxigenacao) VALUES ($1,$2,$3)")
	if err != nil {
		log.Println("Houve um erro ao inserir o dado. " + err.Error())
	}

	insertDadosNoBanco.Exec(dataHora, bati, sp02)
	defer db.Close()
}

func DeletarMedicao(idMedicao string) {
	db := bd.ConectarBD()

	deletarDados, err := db.Prepare("DELETE FROM medicoes WHERE id = $1")
	if err != nil {
		log.Println("Houve um erro ao deletar o dado. " + err.Error())
	}
	deletarDados.Exec(idMedicao)

	defer db.Close()
}

func BuscarProduto(idMedicao string) Medicao {
	db := bd.ConectarBD()

	selecionarMedicoes, err := db.Query("SELECT * FROM medicoes WHERE id = $1", idMedicao)
	if err != nil {
		log.Fatal(err.Error())
	}

	medicaoSelecionada := Medicao{}

	for selecionarMedicoes.Next() {
		var id int
		var batimentos, nivelOxigenacao int
		var dataHora string

		err = selecionarMedicoes.Scan(&id, &dataHora, &batimentos, &nivelOxigenacao)
		if err != nil {
			panic(err.Error())
		}

		medicaoSelecionada.Id = id
		medicaoSelecionada.DataHora = dataHora
		medicaoSelecionada.Batimentos = batimentos
		medicaoSelecionada.NivelOxigenacao = nivelOxigenacao

	}

	defer db.Close()

	return medicaoSelecionada
}

func AtualizaProduto(idConvertInt, batiConvertInt, sp02ConvertInt int, dthr string) {
	db := bd.ConectarBD()

	AtualizarProduto, err := db.Prepare("UPDATE medicoes SET datahora=$1, batimentos=$2, niveloxigenacao=$3 WHERE id = $4")
	if err != nil {
		log.Fatal(err.Error())
	}

	AtualizarProduto.Exec(dthr, batiConvertInt, sp02ConvertInt, idConvertInt)

	defer db.Close()
}
