package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Target struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL             string             `json:"url" bson:"url"`
	IntervalSeconds int                `json:"interval_seconds" bson:"interval_seconds"`
	LastStatus      string             `json:"status" bson:"last_status"`
}
