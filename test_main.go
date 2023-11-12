package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusRibas/api-go-gin/controllers"
	"github.com/viniciusRibas/api-go-gin/database"
	"github.com/viniciusRibas/api-go-gin/models"
)

var ID int

func SetupRotaTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCde(t *testing.T) {
	r := SetupRotaTest()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Erro devia ser igual")
	mockResposta := `{"API diz: ": "E ai gui, tudo beleza"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, respostaBody)
}
func TestListaNdo(t *testing.T) {
	database.Conceta_BD()
	CriaAluno()
	defer deleteAluno()
	r := SetupRotaTest()
	r.GET("/alunos", controllers.ExibeAlunos)
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
	database.Conceta_BD()
	CriaAluno()
	defer deleteAluno()
	r := SetupRotaTest()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
