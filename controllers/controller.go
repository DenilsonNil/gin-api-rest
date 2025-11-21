package controllers

import (
	"github.com/DenilsonNil/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

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
