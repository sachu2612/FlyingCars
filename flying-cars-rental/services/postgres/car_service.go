package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/example/car-api/models"
	"github.com/example/car-api/repositories"
	"github.com/example/car-api/services"
)

type carService struct {
	carRepo repositories.CarRepository
}

// NewCarService creates a new instance of the car service
func NewCarService(carRepo repositories.CarRepository) services.CarService {
	return &carService{
		carRepo: carRepo,
	}
}

func (cs *carService) CreateCar(ctx context.Context, car *models.Car) (*models.Car, error) {
	return cs.carRepo.CreateCar(ctx, car)
}

func (cs *carService) GetCarByID(ctx context.Context, id int64) (*models.Car, error) {
	return cs.carRepo.GetCarByID(ctx, id)
}

func (cs *carService) GetCars(ctx context.Context) ([]*models.Car, error) {
	return cs.carRepo.GetCars(ctx)
}

func (cs *carService) UpdateCar(ctx context.Context, id int64, car *models.Car) (*models.Car, error) {
	// Ensure car exists before updating
	_, err := cs.carRepo.GetCarByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCarNotFound
		}
		return nil, err
	}

	return cs.carRepo.UpdateCar(ctx, id, car)
}

func (cs *carService) DeleteCar(ctx context.Context, id int64) error {
	// Ensure car exists before deleting
	_, err := cs.carRepo.GetCarByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrCarNotFound
		}
		return err
	}

	return cs.carRepo.DeleteCar(ctx, id)
}
