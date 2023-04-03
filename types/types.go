package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bootcamp struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Slug          string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Website       string             `json:"website,omitempty" bson:"website,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Address       string             `json:"address,omitempty" bson:"-"`
	Careers       []string           `json:"careers,omitempty" bson:"careers,omitempty"`
	Housing       bool               `json:"housing,omitempty" bson:"housing,omitempty"`
	JobAssistance bool               `json:"jobAssistance,omitempty" bson:"jobAssistance,omitempty"`
	JobGuarantee  bool               `json:"jobGuarantee,omitempty" bson:"jobGuarantee,omitempty"`
	AcceptGi      bool               `json:"acceptGi,omitempty" bson:"acceptGi,omitempty"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Location struct {
}

// TODO: Add Validation here 👇🏼
