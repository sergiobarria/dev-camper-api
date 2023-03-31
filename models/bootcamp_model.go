package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bootcamp struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

// ===== Validate model here 👇🏼 =====
