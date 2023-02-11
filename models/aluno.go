package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	Nome string `json:"nome" validate:"nonzero, min=6"`
	CPF  string `json:"cpf"  validade:"len=9, regexp=^[0-9]*$"`
	RG   string `json:"rg"   validade:"len=11, regexp=^[0-9]*$"`
	gorm.Model
}

func ValidarDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
