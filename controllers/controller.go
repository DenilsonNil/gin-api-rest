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
		})
		return
	}

	database.DB.Create(&aluno)

	c.JSON(201, aluno)
}

func GetAlunos(c *gin.Context) {
	// alunos := []string{"Ana", "Bruno", "Carla", "Daniel"}
	c.JSON(200, gin.H{
		"alunos": models.Alunos,
	})
}

func GetId(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{
		"alunoId": id,
	})
}
