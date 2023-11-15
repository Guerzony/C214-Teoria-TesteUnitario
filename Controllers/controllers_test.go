package controllers

import (
	"C214-teoria-GO/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ID int

func SetupRotaTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCde(t *testing.T) {

	r := SetupRotaTest()
	r.GET("/:nome", Saudacoes)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
	mockResposta := `{"API diz: ": "E ai gui, tudo beleza?"}`
	assert2.JSONEq(t, mockResposta, resposta.Body.String())
}

func TestTodosAlunos(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Find", mock.Anything).Return(mockDB)
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	TodosAlunosHandler(c, mockDB)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestCriarAluno(t *testing.T) {
	mockDB := &MockDB{}
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	alunoJSON := `{"nome": "John Doe", "rg": "123456789", "cpf": "12345678901"}`

	req, _ := http.NewRequest("POST", "/criar-novo-aluno", bytes.NewBufferString(alunoJSON))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockDB.On("Create", mock.AnythingOfType("*models.Aluno")).Return(mockDB)

	CriarNovoAlunoHandler(c, mockDB)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
	mockDB.AssertExpectations(t)
}

func TestDeletarAluno(t *testing.T) {

	mockDB := &MockDB{}
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	mockDB.On("Delete", mock.AnythingOfType("*models.Aluno"), mock.Anything).Return(mockDB)

	DeletarAlunoHandler(c, mockDB)

	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

	expectedJSON := `{"data":"Aluno deletado com sucesso"}`
	assert2.JSONEq(t, expectedJSON, w.Body.String())
}

func TestEditarAluno(t *testing.T) {
	mockDB := &MockDB{}
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Simula a busca do aluno no banco de dados
	alunoSimulado := &models.Aluno{
		Nome: "John Doe",
		RG:   "123456789",
		CPF:  "12345678901",
	}
	mockDB.On("First", mock.AnythingOfType("*models.Aluno"), mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
		aluno := args.Get(0).(*models.Aluno)
		*aluno = *alunoSimulado
	})

	// Configuração do corpo da requisição (JSON)
	jsonBody := `{"Nome":"Jane",
		"RG":"987654321",
		"CPF":"98765432109"
	}`
	c.Request = httptest.NewRequest(http.MethodPut, "/alunos/1", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Configuração do método Model e UpdateColumns para retornar o próprio mockDB
	mockDB.On("Model", mock.AnythingOfType("*models.Aluno")).Return(mockDB)
	mockDB.On("UpdateColumns", mock.AnythingOfType("*models.Aluno")).Return(mockDB)

	// Chamada da função
	EditarAlunoHandler(c, mockDB)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	// Configuração do banco de dados mock
	mockDB := &MockDB{}

	// Configuração do contexto
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simula o CPF do aluno na URL
	c.Params = append(c.Params, gin.Param{Key: "cpf", Value: "12345678901"})

	// Simula a busca do aluno no banco de dados
	alunoSimulado := &models.Aluno{

		Nome: "John Doe",
		RG:   "123456789",
		CPF:  "12345678901",
	}
	mockDB.On("Where", mock.AnythingOfType("*models.Aluno"), "cpf = ?", "12345678901").Return(mockDB).Run(func(args mock.Arguments) {
		aluno := args.Get(0).(*models.Aluno)
		*aluno = *alunoSimulado
	})

	// Configuração do método First para retornar o próprio mockDB
	mockDB.On("First", mock.AnythingOfType("*models.Aluno")).Return(mockDB)

	// Chamada da função
	BuscaAlunoPorCPFHandler(c, mockDB)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

	// Verifica se a resposta JSON contém os dados do aluno
	expectedJSON := `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Nome":"John Doe","RG":"123456789","CPF":"12345678901"}`
	assert2.JSONEq(t, expectedJSON, w.Body.String())
}

func TestBuscaRGHandler(t *testing.T) {
	// Configuração do banco de dados mock
	mockDB := &MockDB{}

	// Configuração do contexto
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simula o RG do aluno na URL
	c.Params = append(c.Params, gin.Param{Key: "rg", Value: "123456789"})

	// Simula a busca do aluno no banco de dados
	alunoSimulado := &models.Aluno{
		Nome: "John Doe",
		RG:   "123456789",
		CPF:  "12345678901",
	}
	mockDB.On("Where", mock.AnythingOfType("*models.Aluno"), "rg = ?", "123456789").Return(mockDB).Run(func(args mock.Arguments) {
		aluno := args.Get(0).(*models.Aluno)
		*aluno = *alunoSimulado
	})

	// Configuração do método First para retornar o próprio mockDB
	mockDB.On("First", mock.AnythingOfType("*models.Aluno")).Return(mockDB)

	// Chamada da função
	BuscaRGHandler(c, mockDB)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

	// Verifica se a resposta JSON contém os dados do aluno
	expectedJSON := `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Nome":"John Doe","RG":"123456789","CPF":"12345678901"}`
	assert2.JSONEq(t, expectedJSON, w.Body.String())
}
