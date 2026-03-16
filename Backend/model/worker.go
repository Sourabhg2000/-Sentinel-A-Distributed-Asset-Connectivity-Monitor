package model

import (
	"Golang_Test/config"
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func StartMonitorWorker() {
	ticker := time.NewTicker(60 * time.Second) // Check every 60 seconds

	go func() {
		for range ticker.C {
			runHealthChecks()
		}
	}()
}

func runHealthChecks() {
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("targets")

	// 1. Get all targets
	cursor, _ := collection.Find(context.TODO(), bson.M{})
	var targets []Target
	cursor.All(context.TODO(), &targets)

	// 2. Ping each target
	for _, t := range targets {
		resp, err := http.Get(t.URL)
		status := "DOWN"
		if err == nil && resp.StatusCode == 200 {
			status = "UP"
		}

		// 3. Update status in MongoDB
		collection.UpdateOne(context.TODO(), bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"last_status": status}})
	}
}
