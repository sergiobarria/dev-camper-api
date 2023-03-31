package repositories

import (
	"fmt"

	"github.com/sergiobarria/dev-camper-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BootcampRepo interface {
	InsertOne(bootcamp *models.Bootcamp) error
	FindAll() (*[]models.Bootcamp, error)
	FindByID(id string) (*models.Bootcamp, error)
	UpdateOne(id string, bootcamp *models.Bootcamp) error
	DeleteOne(id string) error
}

type BootcampImpl struct {
	db *mongo.Client
}

func NewBootcampRepo(db *mongo.Client) *BootcampImpl {
	return &BootcampImpl{db: db}
}

func (r *BootcampImpl) InsertOne(bootcamp *models.Bootcamp) error {
	fmt.Println("Inserting bootcamp")
	return nil
}

func (r *BootcampImpl) FindAll() (*[]models.Bootcamp, error) {
	fmt.Println("Finding all bootcamps")
	return nil, nil
}

func (r *BootcampImpl) FindByID(id string) (*models.Bootcamp, error) {
	fmt.Println("Finding bootcamp by id", id)
	return nil, nil
}

func (r *BootcampImpl) UpdateOne(id string, bootcamp *models.Bootcamp) error {
	fmt.Println("Updating bootcamp by id", id)
	return nil
}

func (r *BootcampImpl) DeleteOne(id string) error {
	fmt.Println("Deleting bootcamp by id", id)
	return nil
}
