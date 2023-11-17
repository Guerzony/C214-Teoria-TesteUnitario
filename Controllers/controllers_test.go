package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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

func TestStatusCode(t *testing.T) {

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
	gin.SetMode(gin.TestMode)
	mockDB.On("Find", mock.Anything).Return(mockDB)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	TodosAlunosHandler(c, mockDB)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func performRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestTodosAlunosEmpty(t *testing.T) {
	mockDB := new(MockDB)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB.On("Find", mock.Anything).Return(&MockDB{})

	router.GET("/alunos", func(c *gin.Context) {
		TodosAlunosHandler(c, mockDB)
	})

	w := performRequest(router, "GET", "/alunos", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert2.Equal(t, w.Body.String(), "null")
	mockDB.AssertExpectations(t)
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
