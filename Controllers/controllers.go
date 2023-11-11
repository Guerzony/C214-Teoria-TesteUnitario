package controllers

import (
	"net/http"

)

func ExibeAlunos (c * gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao (c * gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: " : "E ai " + nome + ", tudo beleza?"
	})
}

func NovoAluno (c * gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(https.StatusBadRequest, gin.H{"error" : err.error()})
		return
	}
	if err := models.ValidaDados (&aluno); err != nil{
		c.JSON(https.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func Buscar(c * gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	c.Json(http.StatusOK, aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"not found" : "Aluno não encontrado"})
		return
	}	
}

func Delete(c * gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data" : "Aluno deletado com sucesso",
	})
}

func Editar(c * gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)
	if err != nill {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error()
		})
		return
	}

	if err := models.ValidaDados(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaCPF(c * gin.Context){
	var aluno models.Aluno
	cpf :=  c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H["not found" : "Aluno não encontrado"])
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func BuscaRG(c * gin.Context){
	var aluno models.Aluno 
	rg := c.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(https.StatusNotFound, gin.H{"not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}