package repository

import (
	"Go-UserManagement/model/domain"
	"context"
	"database/sql"
	"errors"
)

// Create a struct implementation that follows user_repository interface
type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (User *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, U domain.User) (domain.User, error) {
	// Define the SQL command
	SQL := `INSERT INTO users(email, password, name) values($1, $2, $3) RETURNING id, created_at`
	// Execute the SQL command using DB.QueryRowContext (PostgreSQL)
	err := tx.QueryRowContext(ctx, SQL, U.Email, U.Password, U.Name).Scan(&U.ID, &U.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return U, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.User, error) {
	script := "SELECT id, name, email, password, created_at FROM users WHERE email = ?"
	
	row := db.QueryRowContext(ctx, script, email)
	
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")  // ✅ Better error message
		}
		return user, err
	}
	
	return user, nil
}