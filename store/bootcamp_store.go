package store

import (
	"github.com/sergiobarria/dev-camper-api/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type BootcampStore interface {
	InsertOne(b *types.Bootcamp) error
	FindAll() (*[]types.Bootcamp, error)
	FindByID(id string) (*types.Bootcamp, error)
	UpdateOne(id string, b *types.Bootcamp) error
	DeleteOne(id string) error
}

type BootcampStoreImpl struct {
	db *mongo.Database
}

var coll string = "bootcamps"

func (s *BootcampStoreImpl) InsertOne(b *types.Bootcamp) error {
	return nil
}

func (s *BootcampStoreImpl) FindAll() (*[]types.Bootcamp, error) {
	return nil, nil
}

func (s *BootcampStoreImpl) FindByID(id string) (*types.Bootcamp, error) {
	return nil, nil
}

func (s *BootcampStoreImpl) UpdateOne(id string, b *types.Bootcamp) error {
	return nil
}

func (s *BootcampStoreImpl) DeleteOne(id string) error {
	return nil
}
