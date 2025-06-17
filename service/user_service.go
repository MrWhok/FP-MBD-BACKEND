package service

import (
	"context"
	"github.com/MrWhok/FP-MBD-BACKEND/entity"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
}
