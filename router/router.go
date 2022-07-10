package router

import "github.com/gin-gonic/gin"

// Get new engine
func GetRouter() *gin.Engine {
	return gin.Default()
}