package models

type Aluno struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

var Alunos []Aluno
