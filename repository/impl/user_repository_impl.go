package impl

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (r *userRepositoryImpl) Register(ctx context.Context, nama, email, noTelp, password string) error {
	return r.DB.WithContext(ctx).Exec(`CALL register_customer(?, ?, ?, ?)`, nama, email, noTelp, password).Error
}

func (r *userRepositoryImpl) Login(ctx context.Context, email string) (string, int, error) {
	var customerID int
	var hashedPassword string

	row := r.DB.WithContext(ctx).
		Raw(`SELECT * FROM login_customer(?)`, email).Row()

	err := row.Scan(&customerID, &hashedPassword)
	if err != nil {
		return "", 0, err
	}

	return hashedPassword, customerID, nil
}
