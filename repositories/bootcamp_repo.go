package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/sergiobarria/dev-camper-api/helpers"
	"github.com/sergiobarria/dev-camper-api/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BootcampRepo interface {
	InsertOne(bootcamp *models.Bootcamp) error
	FindAll() (*[]models.Bootcamp, error)
	FindByID(id string) (*models.Bootcamp, error)
	UpdateOne(id string, bootcamp *models.Bootcamp) error
	DeleteOne(id string) error
}

type BootcampImpl struct {
	coll *mongo.Collection
}

func NewBootcampRepo(db *mongo.Client) *BootcampImpl {
	dbName := viper.GetString("DATABASE")

	coll := db.Database(dbName).Collection("bootcamps")
	return &BootcampImpl{coll: coll}
}

func (r *BootcampImpl) InsertOne(bootcamp *models.Bootcamp) error {
	bootcamp.Slug = helpers.Slugify(bootcamp.Name)
	bootcamp.CreatedAt = time.Now()
	bootcamp.UpdatedAt = time.Now()

	result, err := r.coll.InsertOne(context.Background(), bootcamp)
	if err != nil {
		return err
	}

	bootcamp.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *BootcampImpl) FindAll() (*[]models.Bootcamp, error) {
	filter := bson.D{{}}

	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var bootcamps []models.Bootcamp
	if err = cursor.All(context.Background(), &bootcamps); err != nil {
		return nil, err
	}

	return &bootcamps, nil
}

func (r *BootcampImpl) FindByID(id string) (*models.Bootcamp, error) {
	var bootcamp models.Bootcamp
	objId, _ := primitive.ObjectIDFromHex(id) // needs to convert to ObjectID
	filter := bson.D{{Key: "_id", Value: objId}}

	err := r.coll.FindOne(context.Background(), filter).Decode(&bootcamp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("bootcamp with id %s not found", id)
		}
		return nil, err
	}

	return &bootcamp, nil
}

func (r *BootcampImpl) UpdateOne(id string, bootcamp *models.Bootcamp) error {
	bootcamp.UpdatedAt = time.Now()

	objId, _ := primitive.ObjectIDFromHex(id) // needs to convert to ObjectID
	filter := bson.D{{Key: "_id", Value: objId}}
	newDoc := bson.D{{Key: "$set", Value: bootcamp}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result, err := r.coll.FindOneAndUpdate(context.Background(), filter, newDoc, opts).DecodeBytes()
	if err != nil {
		return err
	}

	if err = bson.Unmarshal(result, bootcamp); err != nil {
		return err
	}

	return nil
}

func (r *BootcampImpl) DeleteOne(id string) error {
	objId, _ := primitive.ObjectIDFromHex(id) // needs to convert to ObjectID
	filter := bson.D{{Key: "_id", Value: objId}}

	result, err := r.coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("bootcamp with id %s not found", id)
	}

	return nil
}
