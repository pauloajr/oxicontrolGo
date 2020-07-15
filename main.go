package main

import (
	"fmt"
	"net/http"

	"./routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	fmt.Println("Servidor rodando na porta 321")
	http.ListenAndServe(":321", nil)
}
