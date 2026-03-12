package routes

import (
	"Golang_Test/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Create groups
	o := router.Group("/o")
	r := router.Group("/r")
	// c := router.Group("/c")

	// Example routes structure
	// router.GET("/socket.io/", h.handleWhiteboardsocket)

	// // Open routes (no authentication required)
	r.POST("/register", handlers.RegisterHandler)
	o.POST("/login", handlers.LoginHandler)
	r.POST("/activityTypes", handlers.ActivityTypesHandler)
	r.POST("/addActivityType", handlers.AddActivityTypeHandler)
	r.POST("/bookSession", handlers.BookSessionHandler)
	r.POST("/getMyBookings", handlers.GetMyBookingsHandler)
	r.POST("/completeActivity", handlers.CompleteActivityHandler)
	r.POST("/leaderboard", handlers.GetLeaderboard)

}
