package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Type             string    `json:"type,omitempty" bson:"type,omitempty"`
	Coordinates      []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	FormattedAddress string    `json:"formattedAddress,omitempty" bson:"formattedAddress,omitempty"`
	Street           string    `json:"street,omitempty" bson:"street,omitempty"`
	City             string    `json:"city,omitempty" bson:"city,omitempty"`
	State            string    `json:"state,omitempty" bson:"state,omitempty"`
	Zipcode          string    `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Country          string    `json:"country,omitempty" bson:"country,omitempty"`
}

type Bootcamp struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Slug          string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Website       string             `json:"website,omitempty" bson:"website,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty"`
	Location      Location           `json:"location,omitempty" bson:"location,omitempty"`
	Careers       []string           `json:"careers,omitempty" bson:"careers,omitempty"`
	AverageRating float64            `json:"averageRating,omitempty" bson:"averageRating,omitempty"`
	AverageCost   float64            `json:"averageCost,omitempty" bson:"averageCost,omitempty"`
	Photo         string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Housing       bool               `json:"housing,omitempty" bson:"housing,omitempty"`
	JobAssistance bool               `json:"jobAssistance,omitempty" bson:"jobAssistance,omitempty"`
	JobGuarantee  bool               `json:"jobGuarantee,omitempty" bson:"jobGuarantee,omitempty"`
	AcceptGi      bool               `json:"acceptGi,omitempty" bson:"acceptGi,omitempty"`
	CreatedAt     time.Time          `json:"-" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"-" bson:"updatedAt,omitempty"`
}
