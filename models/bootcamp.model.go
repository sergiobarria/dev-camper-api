package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bootcamp struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

// ====== ADD MODEL VALIDATION HERE 👇🏼 ======
