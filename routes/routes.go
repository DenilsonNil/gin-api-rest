package routes

import (
	"github.com/DenilsonNil/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.POST("/alunos", controllers.CreateAluno)
	r.GET("/alunos/:id", controllers.GetById)
	r.DELETE("/alunos/:id", controllers.DeleteById)
	r.Run(":5001")
}
