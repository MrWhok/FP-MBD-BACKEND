package impl

import (
	"context"
	"errors"

	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

type reservationRepositoryImpl struct {
	*gorm.DB
}

func NewReservationRepositoryImpl(db *gorm.DB) repository.ReservationRepository {
	return &reservationRepositoryImpl{DB: db}
}

func (r *reservationRepositoryImpl) CreateReservation(ctx context.Context, customerID, slotID, tableID, guestCount int) error {
	err := r.WithContext(ctx).Exec(
		`CALL create_reservation(?, ?, ?, ?)`,
		customerID, slotID, tableID, guestCount,
	).Error

	if err != nil {
		// Optionally parse PG error here
		return errors.New("failed to create reservation: " + err.Error())
	}
	return nil
}
