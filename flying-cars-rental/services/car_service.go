package services

import (
	"errors"
	"fmt"
	"github.com/example/car-api/models"
	"github.com/example/car-api/repositories"
)

type CarService struct {
	repo repositories.CarRepository
}

func NewCarService(repo repositories.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (cs *CarService) CreateCar(car *models.Car) (*models.Car, error) {
	if car == nil {
		return nil, errors.New("car cannot be nil")
	}

	return cs.repo.CreateCar(car)
}

func (cs *CarService) GetCarById(id int64) (*models.Car, error) {
	if id <= 0 {
		return nil, errors.New("invalid car id")
	}

	return cs.repo.GetCarById(id)
}

func (cs *CarService) UpdateCar(id int64, car *models.Car) (*models.Car, error) {
	if car == nil {
		return nil, errors.New("car cannot be nil")
	}

	if id <= 0 {
		return nil, errors.New("invalid car id")
	}

	existingCar, err := cs.GetCarById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing car: %v", err)
	}

	existingCar.Make = car.Make
	existingCar.Model = car.Model
	existingCar.Year = car.Year

	updatedCar, err := cs.repo.UpdateCar(existingCar)
	if err != nil {
		return nil, fmt.Errorf("failed to update car: %v", err)
	}

	return updatedCar, nil
}

func (cs *CarService) DeleteCar(id int64) error {
	if id <= 0 {
		return errors.New("invalid car id")
	}

	return cs.repo.DeleteCar(id)
}
