package main

import (
	"github.com/gin-gonic/gin"
)

func getAlunos(c *gin.Context) {
	alunos := []string{"Ana", "Bruno", "Carla", "Daniel"}
	c.JSON(200, gin.H{
		"alunos": alunos,
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", getAlunos)
	r.Run(":5001")
}
