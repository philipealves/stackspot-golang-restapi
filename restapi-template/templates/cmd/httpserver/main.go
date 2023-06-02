package main

import (
	healthCheck "{{project-name}}/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/health", healthCheck.Get)

	router.Run(":8080")
}
