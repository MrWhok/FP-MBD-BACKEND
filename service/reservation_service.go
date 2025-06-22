package service

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

type ReservationService interface {
	Reserve(ctx context.Context, customerID int, req model.ReservationRequest) error
	Reschedule(ctx context.Context, customerID int, req model.RescheduleRequest) error
}
