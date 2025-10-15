package repository

import (
	"Go-UserManagement/model/domain"
	"context"
	"database/sql"
)

// In go always use context for best practices
// Transactional
// Data
// Return data
type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, User domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.User, error)
}