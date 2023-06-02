package handler

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {

	c.JSON(200, gin.H{"status": "Ok"})
}
