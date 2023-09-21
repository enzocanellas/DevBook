package modelos

import "time"

//representa um usuário usando a rede social
type Usuario struct {
	ID       uint64 `json:"id, omitempty"`
	Nome     string `json:"nome, omitempty"`
	Nickname string `json:"nick, omitempty"`
	Email    string `json:"email, omitempty"`
	senha    string `json:"senha, omitempty"`
	CriadoEm time.Time `json:"criadoEm, omitempty"`
} 