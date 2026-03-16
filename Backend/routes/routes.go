package routes

import (
	"Golang_Test/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Create groups
	o := router.Group("/o")
	// r := router.Group("/r")
	// c := router.Group("/c")

	// Example routes structure
	// router.GET("/socket.io/", h.handleWhiteboardsocket)

	// // Open routes (no authentication required)
	o.POST("/target", handlers.CreateTargetHandler)
	o.POST("/targets", handlers.FetchTargetHandler)

}
