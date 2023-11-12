package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
)

func TestValidaDadosDeAluno(t *testing.T) {
	testCases := []struct {
		name     string
		aluno    Aluno
		expected error
	}{
		{
			name: "Dados Validos",
			aluno: Aluno{
				Nome: "Chris",
				RG:   "123456789",
				CPF:  "12345678911",
			},
			expected: nil,
		},
		{
			name: "Nome em Branco",
			aluno: Aluno{
				Nome: "",
				RG:   "123456789",
				CPF:  "12345678911",
			},
			expected: validator.ErrorMap{}, // erro específico nome
		},
		{
			name: "RG Invalidos",
			aluno: Aluno{
				Nome: "Chris",
				RG:   "123", // RG com menos de 11 digitos
				CPF:  "12345678911",
			},
			expected: validator.ErrorMap{}, // erro específico RG
		},
		{
			name: "CPF Invalidos",
			aluno: Aluno{
				Nome: "Chris",
				RG:   "123456789",
				CPF:  "145", // CPF com menos de 11 Digitos
			},
			expected: validator.ErrorMap{}, // erro específico CPF
		},
		{
			name: "Faltando campo",
			aluno: Aluno{
				Nome: "Chris",
				CPF:  "145", // CPF com menos de 11 Digitos
			},
			expected: validator.ErrorMap{}, // erro específico sem campo
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidaDadosDeAluno(&tc.aluno)
			if tc.expected == nil {
				// Se o esperado é nil, o erro não deve ocorrer
				assert.Nil(t, err)
			} else {
				assert.IsType(t, validator.ErrorMap{}, err)
			}
		})
	}

}
