package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	// gin.Run()
	gin.Run(":5001")
}
