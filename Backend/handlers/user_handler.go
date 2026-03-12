package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"Golang_Test/config" // import your config package for MongoClient

	"github.com/cespare/xxhash/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtSecret = []byte("100908070605040302010") // use env var in production

type User struct {
	UserID    string `bson:"userId" json:"userId"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	EmailId   string `bson:"emailId" json:"emailId"`
	// Password will store the hash of the phone number
	Password string `bson:"password" json:"password"`
	Phone    string `bson:"phone" json:"phone"`
	Gender   string `bson:"gender" json:"gender"`
	DOB      string `bson:"dob" json:"dob"`
	Role     string `bson:"role" json:"role"`
}

type ActivityType struct {
	ActivityTypeId string `bson:"activityTypeId" json:"activityTypeId"`
	ActivityName   string `bson:"activityName" json:"activityName"`
	CaloriesPerHr  int    `bson:"caloriesPerHr" json:"caloriesPerHr"`
	ActivityType   string `bson:"activityType" json:"activityType"`
	Icon           string `bson:"icon" json:"icon"`
	IsActive       bool   `bson:"isActive" json:"isActive"`
}

type BookingRequest struct {
	ActivityID   string `json:"activityId"`
	ActivityName string `json:"activityName"`
	UserInfo     struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"emailId"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		UserID    string `json:"userId"`
	} `json:"userInfo"`
}

// JWT Claims
type Claims struct {
	EmailId string `json:"emailId"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

func LoginHandler(c *gin.Context) {
	// 1️⃣ Bind login request
	fmt.Println("inside handler")
	var loginData struct {
		EmailId  string `json:"emailId"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Login data:", loginData)

	// 2️⃣ Connect to MongoDB collection
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("SG_Users")

	// 3️⃣ Find user by email
	var user User
	// var rawData bson.M
	err := collection.FindOne(context.TODO(), bson.M{"emailId": loginData.EmailId}).Decode(&user)
	if err != nil {
		fmt.Println("FindOne error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (User not found or DB error)"})
		return
	}

	fmt.Println("User found:", user)
	// 4️⃣ Hash incoming password with xxhash64 and compare
	inputHash := fmt.Sprintf("%d", xxhash.Sum64String(loginData.Password))
	fmt.Printf("Stored Hash:  %s\n", user.Password)
	fmt.Printf("Input Hash:   %s\n", inputHash)

	if user.Password != inputHash {
		fmt.Println("Password mismatch")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (Password mismatch)"})
		return
	}

	// // 5️⃣ Set role if missing
	userRole := user.Role
	if userRole == "" {
		userRole = "user"
	}

	// // 6️⃣ Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		EmailId: user.EmailId,
		Role:    userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Golang_Test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// // 7️⃣ Return JWT + user info
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"userId":    user.UserID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"emailId":   user.EmailId,
			"phone":     user.Phone,
			"gender":    user.Gender,
			"dob":       user.DOB,
			"role":      userRole,
		},
	})
}

func ActivityTypesHandler(c *gin.Context) {
	// 1. Access the collection (ensure your MongoClient is initialized in your config)
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("activityTypes")

	// 2. Set a context with timeout for the DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. Define an empty slice to hold the results
	var activities []ActivityType

	// 4. Find all documents (using an empty filter bson.M{} to get everything)
	// You can use bson.M{"isActive": true} if you only want active ones
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query activity types"})
		return
	}
	defer cursor.Close(ctx)

	// 5. Iterate through the cursor and decode each document
	for cursor.Next(ctx) {
		var activity ActivityType
		if err := cursor.Decode(&activity); err != nil {
			// Log the error but keep processing other records
			fmt.Printf("Error decoding activity: %v\n", err)
			continue
		}
		activities = append(activities, activity)
	}

	// 6. Check if cursor had errors during iteration
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
		return
	}

	// 7. Return the slice (even if empty, returns [])
	c.JSON(http.StatusOK, activities)
}

func RegisterHandler(c *gin.Context) {
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("SG_Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	// 1. Bind JSON from Vue (This fills 'user.Phone' from the 'phone' key)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	fmt.Println("user", user)

	// 2. Validate that we actually received a phone number
	if user.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required to generate password"})
		return
	}

	// 3. Check if user already exists (by Email)
	var existingUser User
	err := collection.FindOne(ctx, bson.M{"emailId": user.EmailId}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// 4. Set generated fields
	user.UserID = generateUserID()

	// 5. Hash the phone number to be the password
	// This uses xxhash as you requested
	user.Password = fmt.Sprintf("%d", xxhash.Sum64String(user.Phone))

	// 6. Insert the full document into MongoDB
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user to database"})
		return
	}

	// 7. Success Response
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"userId":  user.UserID,
	})
}

func generateUserID() string {
	// Generates: "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	id := uuid.New().String()
	// Returns: "f47ac10b58cc4372a5670e02b2c3d479"
	return strings.ReplaceAll(id, "-", "")
}

func AddActivityTypeHandler(c *gin.Context) {
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("activityTypes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var activity ActivityType
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Generate a unique ID for the activity
	activity.ActivityTypeId = generateUserID() // You can reuse your UUID generator

	_, err := collection.InsertOne(ctx, activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity type added successfully"})
}

func BookSessionHandler(c *gin.Context) {
	var req BookingRequest

	// 1. Bind the JSON payload to our struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload structure"})
		return
	}

	// 2. Prepare database connection
	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("SG_Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Create the document for MongoDB
	newBooking := bson.M{
		"sessionId":    req.ActivityID,
		"activityName": req.ActivityName,
		"userId":       req.UserInfo.UserID,
		"user": bson.M{
			"firstName": req.UserInfo.FirstName,
			"lastName":  req.UserInfo.LastName,
			"emailId":   req.UserInfo.Email,
			"phone":     req.UserInfo.Phone,
			"role":      req.UserInfo.Role,
		},
		"status":   "Scheduled",
		"bookedAt": time.Now(),
	}

	// 4. Insert into MongoDB
	result, err := collection.InsertOne(ctx, newBooking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Booking successful",
		"bookingId": result.InsertedID,
	})
}

func GetMyBookingsHandler(c *gin.Context) {
	var req struct {
		UserID string `json:"userId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := config.MongoClient.Database("NewLMS2021_DevDB").Collection("SG_Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter by UserID
	cursor, err := collection.Find(ctx, bson.M{"userId": req.UserID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	var bookings []bson.M
	cursor.All(ctx, &bookings)
	c.JSON(http.StatusOK, bookings)
}

func CompleteActivityHandler(c *gin.Context) {
	var req struct {
		SessionID    string `json:"sessionId"` // Now using sessionId
		UserID       string `json:"userId"`
		ActivityName string `json:"activityName"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db := config.MongoClient.Database("NewLMS2021_DevDB")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Update Booking Status based on SessionID AND UserID
	// This targets the specific record in SG_Bookings
	filter := bson.M{
		"sessionId": req.SessionID,
		"userId":    req.UserID,
		"status":    "Scheduled", // Only update if not already completed
	}

	update := bson.M{
		"$set": bson.M{
			"status":      "Completed",
			"completedAt": time.Now(),
		},
	}

	result, err := db.Collection("SG_Bookings").UpdateOne(ctx, filter, update)

	if err != nil || result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active booking found for this session"})
		return
	}

	// 2. Fetch calories from activityTypes to log them
	var activity bson.M
	db.Collection("activityTypes").FindOne(ctx, bson.M{"activityTypeId": req.SessionID}).Decode(&activity)

	calories := 300
	if val, ok := activity["caloriesPerHr"].(int32); ok {
		calories = int(val)
	}

	// 3. Insert into SG_Logs for Leaderboard
	db.Collection("SG_Logs").InsertOne(ctx, bson.M{
		"userId":       req.UserID,
		"activityName": req.ActivityName,
		"calories":     calories,
		"date":         time.Now(),
	})

	c.JSON(http.StatusOK, gin.H{"message": "Session marked as completed!"})
}

type LeaderboardUser struct {
	UserID       string `json:"userId" bson:"userId"`
	FirstName    string `json:"firstName" bson:"firstName"`
	LastName     string `json:"lastName" bson:"lastName"`
	Role         string `json:"role" bson:"role"`
	BookingCount int    `json:"bookingCount" bson:"bookingCount"`
}

func GetLeaderboard(c *gin.Context) {
	db := config.MongoClient.Database("NewLMS2021_DevDB")
	collection := db.Collection("SG_Bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		// 1️⃣ Only completed activities
		{{Key: "$match", Value: bson.M{
			"status": "Completed",
		}}},

		// 2️⃣ Group by userId and count
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$userInfo.userId"},
			{Key: "activityCount", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "firstName", Value: bson.D{{Key: "$first", Value: "$userInfo.firstName"}}},
			{Key: "lastName", Value: bson.D{{Key: "$first", Value: "$userInfo.lastName"}}},
		}}},

		// 3️⃣ Sort highest activity count first
		{{Key: "$sort", Value: bson.D{
			{Key: "activityCount", Value: -1},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Aggregation failed"})
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data parsing failed"})
		return
	}

	// 4️⃣ Add ranking manually
	for i := range results {
		results[i]["rank"] = i + 1
	}

	c.JSON(http.StatusOK, results)
}
