package repositories

import (
	"context"
	"database/sql"

	"flying-cars-rental/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	sqlStatement := `SELECT * FROM users WHERE id=$1`
	row := u.db.QueryRowContext(ctx, sqlStatement, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	row := u.db.QueryRowContext(ctx, sqlStatement, email)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	sqlStatement := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id`
	err := u.db.QueryRowContext(ctx, sqlStatement, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
