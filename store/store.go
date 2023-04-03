package store

import "go.mongodb.org/mongo-driver/mongo"

type StoreImpl struct {
	Bootcamp BootcampStore
}

func New(db *mongo.Database) *StoreImpl {
	return &StoreImpl{
		Bootcamp: &BootcampStoreImpl{db: db},
		// Add more stores here 👇🏼
	}
}
