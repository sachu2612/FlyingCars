package postgres

import (
	"database/sql"
	"errors"
	"flying-cars-rental/models"
)

// CarRepository represents the repository for the cars
type CarRepository struct {
	DB *sql.DB
}

// GetAll returns all the cars
func (repo *CarRepository) GetAll() ([]*models.Car, error) {
	rows, err := repo.DB.Query("SELECT id, brand, model, price, year FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*models.Car{}

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Year); err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

// GetByID returns a car by ID
func (repo *CarRepository) GetByID(ID int64) (*models.Car, error) {
	var car models.Car
	row := repo.DB.QueryRow("SELECT id, brand, model, price, year FROM cars WHERE id=$1", ID)
	if err := row.Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Year); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return &car, nil
}

// Create inserts a new car into the database
func (repo *CarRepository) Create(car *models.Car) error {
	sql := `INSERT INTO cars (brand, model, price, year) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.DB.QueryRow(sql, car.Brand, car.Model, car.Price, car.Year).Scan(&car.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing car
func (repo *CarRepository) Update(ID int64, car *models.Car) error {
	sql := `UPDATE cars SET brand=$1, model=$2, price=$3, year=$4 WHERE id=$5`
	result, err := repo.DB.Exec(sql, car.Brand, car.Model, car.Price, car.Year, ID)
	if err != nil {
		return err
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsUpdated == 0 {
		return models.ErrNoRecord
	}
	return nil
}

// Delete deletes a car by ID
func (repo *CarRepository) Delete(ID int64) error {
	sql := `DELETE FROM cars WHERE id=$1`
	result, err := repo.DB.Exec(sql, ID)
	if err != nil {
		return err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsDeleted == 0 {
		return models.ErrNoRecord
	}
	return nil
}
