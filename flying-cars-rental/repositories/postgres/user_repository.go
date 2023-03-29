package postgres

import (
	"database/sql"
	"errors"

	"flying-cars-rental/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	err := r.db.QueryRow("INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}
