package repository

import "context"

type ReservationRepository interface {
	CreateReservation(ctx context.Context, customerID, slotID, tableID, guestCount int) error
}
