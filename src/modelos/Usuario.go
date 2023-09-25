package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//representa um usuário usando a rede social
type Usuario struct {
	ID       uint64 `json:"id, omitempty"`
	Nome     string `json:"nome, omitempty"`
	Nick string `json:"nick, omitempty"`
	Email    string `json:"email, omitempty"`
	Senha    string `json:"senha, omitempty"`
	CriadoEm time.Time `json:"criadoEm, omitempty"`
} 

func (usuario *Usuario) validar(etapa string) error { // valida a condição para o usuario ser cadastrado no banco de dados
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}
	
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")}
	
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) Preparar(etapa string) error { // vai chamar os metodos para validar e formatar o usuario recebidp
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.Formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) Formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
 