package service

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

type ReservationService interface {
	Reserve(ctx context.Context, customerID int, req model.ReservationRequest) error
	Reschedule(ctx context.Context, customerID int, req model.RescheduleRequest) error
	CancelReservation(ctx context.Context, reservationID int, customerID int) error
	EditReservation(ctx context.Context, customerID int, reservationID int, newGuestCount int) error
}
