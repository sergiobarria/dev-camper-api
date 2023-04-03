package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type BootcampModel struct {
	db *mongo.Database
}

var coll string = "bootcamps"

// ===== Validate model here 👇🏼 =====

// ===== Model Funcs here 👇🏼 =====
func (m *BootcampModel) InsertOne(b *Bootcamp) error {
	// b.Slug = helpers.Slugify(b.Name)
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	result, err := m.db.Collection(coll).InsertOne(context.Background(), b)
	if err != nil {
		return err
	}

	b.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *BootcampModel) FindAll() ([]Bootcamp, error) {
	filter := bson.D{{}}

	cursor, err := m.db.Collection(coll).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var bootcamps []Bootcamp
	if err = cursor.All(context.Background(), &bootcamps); err != nil {
		return nil, err
	}

	return bootcamps, nil
}

func (m *BootcampModel) FindByID(id string) (*Bootcamp, error) {
	var bootcamp Bootcamp
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}

	err := m.db.Collection(coll).FindOne(context.Background(), filter).Decode(&bootcamp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("bootcamp with ID %s not found", id)
		}
		return nil, err
	}

	return &bootcamp, nil
}

func (m *BootcampModel) UpdateOne(id string, b *Bootcamp) error {
	b.UpdatedAt = time.Now()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	newDoc := bson.D{{Key: "$set", Value: b}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result, err := m.db.Collection(coll).FindOneAndUpdate(context.Background(), filter, newDoc, opts).DecodeBytes()
	if err != nil {
		return err
	}

	if err = bson.Unmarshal(result, b); err != nil {
		return err
	}

	return nil
}

func (m *BootcampModel) DeleteOne(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}

	result, err := m.db.Collection(coll).DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("bootcamp with ID %s not found", id)
	}

	return nil
}
