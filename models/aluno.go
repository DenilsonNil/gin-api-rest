package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

var Alunos []Aluno
