package controllers

import (
	"net/http"

	"C214-teoria-GO/database"
	"C214-teoria-GO/models"

	"github.com/gin-gonic/gin"
)

func bd() database.DBInterface {
	return database.Conecta_BD()
}

func TodosAlunos(c *gin.Context)      { ExecutarHandler(c, TodosAlunosHandler) }
func CriarNovoAluno(c *gin.Context)   { ExecutarHandler(c, CriarNovoAlunoHandler) }
func BuscarAlunoPorID(c *gin.Context) { ExecutarHandler(c, BuscarAlunoPorIDHandler) }
func DeletarAluno(c *gin.Context)     { ExecutarHandler(c, DeletarAlunoHandler) }
func EditarAluno(c *gin.Context)      { ExecutarHandler(c, EditarAlunoHandler) }
func BuscaAlunoPorCPF(c *gin.Context) { ExecutarHandler(c, BuscaAlunoPorCPFHandler) }
func BuscaRG(c *gin.Context)          { ExecutarHandler(c, BuscaRGHandler) }

func ExecutarHandler(c *gin.Context, handler func(c *gin.Context, db database.DBInterface)) {
	handler(c, bd())
}

func TodosAlunosHandler(c *gin.Context, db database.DBInterface) {
	var alunos []models.Aluno
	db.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz: ": "E ai " + nome + ", tudo beleza?",
	})
}

func CriarNovoAlunoHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorIDHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.First(&aluno, id)
	c.JSON(http.StatusOK, aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"not found": "Aluno não encontrado"})
		return
	}
}

func DeletarAlunoHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func EditarAlunoHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPFHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	db.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func BuscaRGHandler(c *gin.Context, db database.DBInterface) {
	var aluno models.Aluno
	rg := c.Param("rg")
	db.Where(&models.Aluno{RG: rg}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
