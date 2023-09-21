package respostas

import (
	"encoding/json"
	"log"
	"net/http"


)


func JSON(w http.ResponseWriter, statusCode int, dados interface{}) { // retorna uma resposta em JSON para a requisição
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil { // retorna um erro em formato JSON
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	} {
		Erro: erro.Error(),
	}) 
}