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
	existingReservation, err := s.repo.GetReservationByID(ctx, req.ReservationID)
	if err != nil {
		return fmt.Errorf("failed to get existing reservation: %w", err)
	}

	if existingReservation.CustomerID != customerID {
		return fmt.Errorf("you are not authorized to reschedule this reservation")
	}

	if existingReservation.Status == "Confirmed" {
		oldSlotDate, err := s.getSlotDate(ctx, existingReservation.SlotID)
		if err != nil {
			return fmt.Errorf("failed to get old slot date: %w", err)
		}
		if time.Now().Add(24 * time.Hour).After(oldSlotDate) {
			return fmt.Errorf("reschedule is only allowed up to H-1 before the original reservation date")
		}
	}

	notificationChannel := "email"
	if err := s.repo.RescheduleReservation(
		ctx,
		req.ReservationID,
		req.NewSlotID,
		req.NewGuestCount,
		notificationChannel,
	); err != nil {
		return fmt.Errorf("reschedule failed: %w", err)
	}
	return nil
}

func (s *reservationServiceImpl) CancelReservation(ctx context.Context, reservationID int, customerID int) error {
	reservation, err := s.repo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return err
	}
	if reservation.CustomerID != customerID {
		return fmt.Errorf("unauthorized")
	}
	return s.repo.CancelReservation(ctx, reservationID)
}

func (s *reservationServiceImpl) EditReservation(
	ctx context.Context,
	customerID int,
	reservationID int,
	newGuestCount int,
) error {

	reservation, err := s.repo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return fmt.Errorf("failed to get reservation: %w", err)
	}
	if reservation.CustomerID != customerID {
		return fmt.Errorf("unauthorized")
	}
	if reservation.Status == "Cancelled" {
		return fmt.Errorf("cannot edit a cancelled reservation")
	}
	if err := s.repo.EditReservation(ctx, reservationID, newGuestCount); err != nil {
		return fmt.Errorf("failed to edit reservation: %w", err)
	}
	return nil
}

func (s *reservationServiceImpl) getSlotDate(ctx context.Context, slotID int) (time.Time, error) {
	return time.Time{}, fmt.Errorf("getSlotDate not implemented")
}
