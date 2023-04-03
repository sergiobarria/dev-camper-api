package models

import "go.mongodb.org/mongo-driver/mongo"

type Models struct {
	Bootcamp BootcampModel
}

func NewModels(db *mongo.Database) Models {
	return Models{
		Bootcamp: BootcampModel{db: db},
	}
}
