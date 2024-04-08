package controllers

import (
	"Teste/Alura/ApiRest/database"
	"Teste/Alura/ApiRest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func Todaspersonalidas(w http.ResponseWriter, r *http.Request) {

	// setando o tipo da resposta a ser JSON
	// w.Header().Set("Content-type", "application/json")
	//pasta "middleware" executando a função pra todas requisições

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func UmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var NovaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&NovaPersonalidade)
	database.DB.Create(&NovaPersonalidade)
	json.NewEncoder(w).Encode(NovaPersonalidade)

}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.First(&personalidade, id)
	// se temos um dado que recebemos em json, e precisamos converter esse dado pra que o GO entenda, usamos Decoder
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)

	// se temos um dado que é json, e quer somente exibir, usamos Encode
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
