package models

import "gorm.io/gorm"

type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
	gorm.Model
}

var Alunos []Aluno