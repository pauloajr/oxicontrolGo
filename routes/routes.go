package routes

import (
	"net/http"

	"../controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
