package impl

import (
	"context"
	"errors"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (s *userServiceImpl) Register(ctx context.Context, req model.UserRegisterModel) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.UserRepository.Register(ctx, req.Nama, req.Email, req.NoTelp, string(hashedPassword))
}

func (s *userServiceImpl) Login(ctx context.Context, req model.UserLoginModel) (int, string, error) {
	hashedPassword, customerID, role, err := s.UserRepository.Login(ctx, req.Email)
	if err != nil {
		return 0, "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return 0, "", errors.New("password incorrect")
	}

	return customerID, role, nil
}
