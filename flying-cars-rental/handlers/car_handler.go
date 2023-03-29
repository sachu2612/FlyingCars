package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var carDB *sql.DB

type Car struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Year     int    `json:"year"`
	RentalID int    `json:"rentalid"`
}

// CreateCarInput defines the input for creating a new car
type CreateCarInput struct {
	Name     string `json:"name" binding:"required"`
	Brand    string `json:"brand" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	RentalID string `json:"rental_id" binding:"required"`
}

// UpdateCarInput defines the input for updating an existing car
type UpdateCarInput struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Year  int    `json:"year"`
}

// GetCarsResponse defines the response for retrieving a list of cars
type GetCarsResponse struct {
	Cars []Car `json:"cars"`
}

// GetCarResponse defines the response for retrieving a single car
type GetCarResponse struct {
	Car Car `json:"car"`
}

// CreateCar handles the creation of a new car
func CreateCar(c *gin.Context) {
	var input CreateCarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    rentalID, err := strconv.Atoi(input.RentalID)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
    return
}

	result, err := carDB.Exec("INSERT INTO cars (name, brand, year, rental_id) VALUES (?, ?, ?, ?)", input.Name, input.Brand, input.Year, rentalID )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetCars handles the retrieval of a list of cars
func GetCars(c *gin.Context) {
	rows, err := carDB.Query("SELECT * FROM cars")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
}

defer rows.Close()

var cars []Car

for rows.Next() {
var car Car

err := rows.Scan(&car.ID, &car.Brand, &car.Name, &car.Year, &car.Year, &car.RentalID)
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
cars = append(cars, car)
}

if err := rows.Err(); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

c.JSON(http.StatusOK, cars)
}

// UpdateCar handles the updating of an existing car
func UpdateCar(c *gin.Context) {
id := c.Param("id")
var car Car
if err := c.ShouldBindJSON(&car); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

stmt, err := carDB.Prepare("UPDATE cars SET brand=?, model=?, year=?, price=? WHERE id=?")
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
defer stmt.Close()
_, err = stmt.Exec(car.Brand, car.RentalID, car.Year, car.Name, car.ID, id)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
}

// DeleteCar handles the deletion of an existing car
func DeleteCar(c *gin.Context) {
id := c.Param("id")
stmt, err := carDB.Prepare("DELETE FROM cars WHERE id=?")
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
defer stmt.Close()

_, err = stmt.Exec(id)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}