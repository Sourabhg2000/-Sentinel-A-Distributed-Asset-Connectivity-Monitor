package handlers

import (
	"Golang_Test/config"
	"Golang_Test/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTargetHandler(c *gin.Context) {
	fmt.Println("Inside handler")
	var target model.Target

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Println("after binding data")

	// 2. Set Defaults
	// Note: Don'target manually set target.ID here; let MongoDB generate it
	// OR set it manually if your struct tag doesn'target use omitempty.
	target.LastStatus = "PENDING"
	if target.IntervalSeconds == 0 {
		target.IntervalSeconds = 60
	}

	// 3. Insert into MongoDB
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("targets")
	res, err := collection.InsertOne(c.Request.Context(), target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to MongoDB"})
		return
	}

	fmt.Println("after insering data")

	// 4. THE CRITICAL STEP: Map the ID back to the struct
	// This ensures your Vue frontend receives the _id to use for deletes/updates.
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		target.ID = oid
	}

	// 5. Return success
	c.JSON(http.StatusCreated, target)
}

func FetchTargetHandler(c *gin.Context) {
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("targets")
	cursor, err := collection.Find(c.Request.Context(), primitive.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from MongoDB"})
		return
	}
	var targets []model.Target
	if err := cursor.All(c.Request.Context(), &targets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode MongoDB results"})
		return
	}
	c.JSON(http.StatusOK, targets)
}
