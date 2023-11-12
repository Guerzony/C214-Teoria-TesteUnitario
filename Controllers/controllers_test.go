package controllers

import (
	"C214-teoria-GO/database"
	"C214-teoria-GO/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	teste "github.com/stretchr/testify/assert"
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
	teste.JSONEq(t, mockResposta, resposta.Body.String())
}

func TestListaNdo(t *testing.T) {
	database.Conecta_BD()
	CriaAluno()
	defer deleteAluno()
	r := SetupRotaTest()
	r.GET("/alunos", TodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func CriaAluno() {
	aluno := models.Aluno{Nome: "Nome do aluno", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func deleteAluno() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestBuscaCpf(t *testing.T) {
	database.Conecta_BD()
	CriaAluno()
	defer deleteAluno()
	r := SetupRotaTest()
	r.GET("/alunos/cpf/:cpf", BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
