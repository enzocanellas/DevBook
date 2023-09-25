package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc { // essa função escreve informações da requisição no terminal  
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf(" %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc { // essa função verifica se o usuário fazendo a requisição está autenticado
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return 
		}
		proximaFuncao(w, r)
	}
}