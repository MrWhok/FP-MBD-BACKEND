package repository

import "context"

type ReservationRepository interface {
	CreateReservation(ctx context.Context, customerID, slotID, tableID, guestCount int) error
	GetReservationByID(ctx context.Context, reservationID int) (*model.Reservation, error)
	RescheduleReservation(ctx context.Context, reservationID, newSlotID, newGuestCount int, notificationChannel string) error
	FindAvailableTableForSlot(ctx context.Context, slotID int, guestCount int) (int, error)
}
