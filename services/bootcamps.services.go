package services

import (
	"context"
	"errors"
	"time"

	"github.com/sergiobarria/dev-camper-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BootcampsService interface {
	CreateBootcamp(*models.Bootcamp) (*models.Bootcamp, error)
	FindBootcamps(page, limit int, sort, selectFields string) ([]*models.Bootcamp, error)
	FindBootcampByID(string) (*models.Bootcamp, error)
	UpdateBootcampByID(string, *models.Bootcamp) (*models.Bootcamp, error)
	DeleteBootcampByID(string) error
}

// ===============================
// === Services Implementation ===
// ===============================
type BootcampsServiceImpl struct {
	bootcampsCollection *mongo.Collection
	ctx                 context.Context
}

func NewBootcampsService(bootcampsColl *mongo.Collection, ctx context.Context) BootcampsService {
	return &BootcampsServiceImpl{
		bootcampsCollection: bootcampsColl,
		ctx:                 ctx,
	}
}

func (s *BootcampsServiceImpl) CreateBootcamp(bootcamp *models.Bootcamp) (*models.Bootcamp, error) {
	bootcamp.CreatedAt = time.Now()

	res, err := s.bootcampsCollection.InsertOne(s.ctx, bootcamp)
	if err != nil {
		if e, ok := err.(mongo.WriteException); ok && e.WriteErrors[0].Code == 11000 {
			return nil, errors.New("duplicate key error")
		}
	}

	opt := options.Index().SetUnique(true) // set unique index to name

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt} // create index model

	if _, err := s.bootcampsCollection.Indexes().CreateOne(s.ctx, index); err != nil {
		return nil, errors.New("error creating index")
	}

	var newBootcamp *models.Bootcamp
	query := bson.M{"_id": res.InsertedID}
	if err := s.bootcampsCollection.FindOne(s.ctx, query).Decode(&newBootcamp); err != nil {
		return nil, errors.New("error decoding bootcamp")
	}

	return newBootcamp, nil
}

func (s *BootcampsServiceImpl) FindBootcamps(page int, limit int, sort, selectFields string) ([]*models.Bootcamp, error) {
	if page < 1 || page == 0 {
		page = 1
	}

	if limit < 1 || limit == 0 {
		limit = 1
	}

	if sort == "" {
		sort = "-createdAt"
	}

	// TODO: improve filters options

	var bootcamps []*models.Bootcamp
	skip := (page - 1) * limit

	opts := options.FindOptions{}
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))
	// opts.SetSort(bson.M{sort: 1})
	// opts.SetProjection(bson.M{selectFields: 1})

	query := bson.M{}

	cursor, err := s.bootcampsCollection.Find(s.ctx, query, &opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var bootcamp *models.Bootcamp
		if err := cursor.Decode(&bootcamp); err != nil {
			return nil, err
		}

		bootcamps = append(bootcamps, bootcamp)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(bootcamps) == 0 {
		return []*models.Bootcamp{}, nil
	}

	return bootcamps, nil
}

func (s *BootcampsServiceImpl) FindBootcampByID(id string) (*models.Bootcamp, error) {
	var bootcamp *models.Bootcamp
	docID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": docID}

	if err := s.bootcampsCollection.FindOne(s.ctx, query).Decode(&bootcamp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return bootcamp, nil
}

func (s *BootcampsServiceImpl) UpdateBootcampByID(id string, bootcamp *models.Bootcamp) (*models.Bootcamp, error) {
	docID, _ := primitive.ObjectIDFromHex(id)
	bootcamp.UpdatedAt = time.Now()

	query := bson.M{"_id": docID}
	data, err := bson.Marshal(bootcamp)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(data, &bootcamp)
	if err != nil {
		return nil, err
	}

	update := bson.M{"$set": bootcamp}

	var updatedBootcamp *models.Bootcamp

	res := s.bootcampsCollection.FindOneAndUpdate(s.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	if err := res.Decode(&updatedBootcamp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("bootcamp not found")
		}
		return nil, err
	}

	return updatedBootcamp, nil
}

func (s *BootcampsServiceImpl) DeleteBootcampByID(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": docID}

	res, err := s.bootcampsCollection.DeleteOne(s.ctx, query)
	if err != nil {
		return models.ErrNotFound
	}

	if res.DeletedCount == 0 {
		return models.ErrNotFound
	}

	return nil
}
