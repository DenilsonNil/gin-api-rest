package controllers

import (
	"github.com/DenilsonNil/gin-api-rest/database"
	"github.com/DenilsonNil/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"cause": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&aluno).Error; err != nil { // 👈 AQUI
		c.JSON(500, gin.H{
			"error": "Erro ao salvar no banco",
			"cause": err.Error(),
		})
		return
	}

	c.JSON(201, aluno)
}

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func GetById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"NOT FOUND": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}
