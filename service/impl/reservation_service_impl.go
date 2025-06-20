package impl

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
	"github.com/MrWhok/FP-MBD-BACKEND/repository"
)

type reservationServiceImpl struct {
	repo repository.ReservationRepository
}

func NewReservationServiceImpl(repo repository.ReservationRepository) *reservationServiceImpl {
	return &reservationServiceImpl{repo: repo}
}

func (s *reservationServiceImpl) Reserve(ctx context.Context, customerID int, req model.ReservationRequest) error {
	return s.repo.CreateReservation(ctx, customerID, req.SlotID, req.TableID, req.GuestCount)
}
