package routes

import (
	"Teste/Alura/ApiRest/controllers"
	"Teste/Alura/ApiRest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleWare) //  todas requisições retornarão o content type no formato json
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidas", controllers.Todaspersonalidas).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.CriarNovaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades", controllers.UmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")

	// CORS se trata de portas,protocolos e/ou host diferentes entre as integrações
	// AllowedOrigins permite integrações de origens difentes, com o front-end
	// * é utilizado pra permitir que qualquer aplicação consiga acessar e usar o dados da minha API
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
