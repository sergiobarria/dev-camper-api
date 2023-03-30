package repositories

import (
	"fmt"

	"github.com/sergiobarria/dev-camper-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BootcampRepository interface {
	InsertOne(bootcamp *models.Bootcamp) (*models.Bootcamp, error)
	FindAll() (*[]models.Bootcamp, error)
	FindByID(id string) (*models.Bootcamp, error)
	UpdateOne(id string, bootcamp models.Bootcamp) (*models.Bootcamp, error)
	DeleteOne(id string) (*models.Bootcamp, error)
}

type BootcampImpl struct {
	db *mongo.Client
}

func NewBootcampRepo(db *mongo.Client) *BootcampImpl {
	return &BootcampImpl{db}
}

func (r *BootcampImpl) InsertOne(bootcamp *models.Bootcamp) (*models.Bootcamp, error) {
	fmt.Println("InsertOne() called")
	return &models.Bootcamp{}, nil
}

func (r *BootcampImpl) FindAll() (*[]models.Bootcamp, error) {
	fmt.Println("FindAll() called")
	return &[]models.Bootcamp{}, nil
}

func (r *BootcampImpl) FindByID(id string) (*models.Bootcamp, error) {
	fmt.Println("Get() called")
	return &models.Bootcamp{}, nil
}

func (r *BootcampImpl) UpdateOne(id string, bootcamp models.Bootcamp) (*models.Bootcamp, error) {
	fmt.Println("UpdateOne() called")
	return &models.Bootcamp{}, nil
}

func (r *BootcampImpl) DeleteOne(id string) (*models.Bootcamp, error) {
	fmt.Println("DeleteOne() called")
	return &models.Bootcamp{}, nil
}
