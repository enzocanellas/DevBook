package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct { // representa uma publicacao feita por um usuario
	ID        uint64 `json:"id, omitempty"`
	Titulo    string `json: "titulo, omitempty"`
	Conteudo  string `json: "conteudo, omitempty"`
	AutorID   uint64 `json: "autorId, omitempty"`
	AutorNick string `json: "autorNick, omitempty"`
	Curtidas  uint64 `json: "curtidas"`
	CriadaEm  time.Time `json: "criadaEm, omitempty"`
}

func (publicacao *Publicacao) Preparar() error { // essa função vai chamar os métodos para validar e formatar a publicção recebida
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}