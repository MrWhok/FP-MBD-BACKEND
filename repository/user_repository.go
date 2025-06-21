package repository

import (
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, nama, email, noTelp, password string) error
	Login(ctx context.Context, email string) (string, int, string, error)
	FindRolesByCustomerID(ctx context.Context, customerID int) ([]map[string]interface{}, error)
}
