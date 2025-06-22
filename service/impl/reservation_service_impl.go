package impl

import (
	"context"
	"fmt"
	"time"

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

func (s *reservationServiceImpl) Reschedule(ctx context.Context, customerID int, req model.RescheduleRequest) error {
	existingReservation, err := s.reservationRepo.GetReservationByID(ctx, req.ReservationID)
	if err != nil {
		return fmt.Errorf("failed to get existing reservation: %w", err)
	}

	if existingReservation.CustomerID != customerID {
		return fmt.Errorf("you are not authorized to reschedule this reservation")
	}
	if existingReservation.Status == "Confirmed" {
		oldSlotDate, err := s.getSlotDate(ctx, existingReservation.SlotID) // Anda perlu membuat func ini
		if err != nil {
			return fmt.Errorf("failed to get old slot date: %w", err)
		}
		if time.Now().Add(24 * time.Hour).After(oldSlotDate) { // Today + 1 day should be before old slot date
			return fmt.Errorf("reschedule is only allowed up to H-1 before the original reservation date")
		}
	}

	notificationChannel := "email"
	err = s.reservationRepo.RescheduleReservation(ctx, req.ReservationID, req.NewSlotID, req.NewGuestCount, notificationChannel)
	if err != nil {
		return fmt.Errorf("reschedule failed: %w", err)
	}

	return nil
}
