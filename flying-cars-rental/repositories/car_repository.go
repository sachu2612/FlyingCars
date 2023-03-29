package repositories

import (
"context"
"database/sql"
"flying-cars-rental/models"
)
type CarRepository struct {
db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
return &CarRepository{db: db}
}

func (c *CarRepository) GetAllCars(ctx context.Context) ([]*models.Car, error) {
sqlStatement := `SELECT * FROM cars`
rows, err := c.db.QueryContext(ctx, sqlStatement)
if err != nil {
return nil, err
}
defer rows.Close()
var cars []*models.Car
for rows.Next() {
	car := &models.Car{}
	err = rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price)
	if err != nil {
		return nil, err
	}
	cars = append(cars, car)
}

if err = rows.Err(); err != nil {
	return nil, err
}
return cars, nil
}

func (c *CarRepository) GetCarById(ctx context.Context, id int64) (*models.Car, error) {
sqlStatement := `SELECT * FROM cars WHERE id=$1`
row := c.db.QueryRowContext(ctx, sqlStatement, id)
car := &models.Car{}
err := row.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price)
if err != nil {
return nil, err
}
return car, nil
}

func (c *CarRepository) CreateCar(ctx context.Context, car *models.Car) (*models.Car, error) {
sqlStatement := `INSERT INTO cars (brand, model, year, price) VALUES ($1, $2, $3, $4) RETURNING id`
err := c.db.QueryRowContext(ctx, sqlStatement, car.Brand, car.Model, car.Year, car.Price).Scan(&car.ID)
if err != nil {
return nil, err
}
return car, nil
}

func (c *CarRepository) UpdateCar(ctx context.Context, car *models.Car) error {
sqlStatement := `UPDATE cars SET brand=$1, model=$2, year=$3, price=$4 WHERE id=$5`
_, err := c.db.ExecContext(ctx, sqlStatement, car.Brand, car.Model, car.Year, car.Price, car.ID)
if err != nil {
return err
}
return nil
}

func (c *CarRepository) DeleteCar(ctx context.Context, id int64) error {
sqlStatement := `DELETE FROM cars WHERE id=$1`
_, err := c.db.ExecContext(ctx, sqlStatement, id)
if err != nil {
return err
}
return nil
}