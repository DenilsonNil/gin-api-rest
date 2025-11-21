package controllers

import "github.com/gin-gonic/gin"

func GetAlunos(c *gin.Context) {
	alunos := []string{"Ana", "Bruno", "Carla", "Daniel"}
	c.JSON(200, gin.H{
		"alunos": alunos,
	})
}
