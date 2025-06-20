package service

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

type UserService interface {
	Register(ctx context.Context, request model.UserRegisterModel) error
	Login(ctx context.Context, req model.UserLoginModel) (int, []map[string]interface{}, error)
}
