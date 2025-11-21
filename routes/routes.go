package routes

import (
	"github.com/DenilsonNil/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/alunos/:id", controllers.GetId)
	r.Run(":5001")
}
