package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/MrWhok/FP-MBD-BACKEND/model"
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
	sql := `CALL create_reservation(?, ?, ?, ?)`
	err := r.DB.WithContext(ctx).Exec(sql, customerID, slotID, tableID, guestCount).Error

	if err != nil {
		return errors.New("failed to create reservation: " + err.Error())
	}
	return nil
}

func (r *reservationRepositoryImpl) GetReservationByID(ctx context.Context, reservationID int) (*model.Reservation, error) {
	var res model.Reservation
	query := `SELECT reservation_id, customer_id, slot_id, table_id, guest_count, status, created_at
			  FROM reservation WHERE reservation_id = ?`
	err := r.db.WithContext(ctx).Raw(query, reservationID).Scan(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("reservation with ID %d not found", reservationID)
		}
		return nil, fmt.Errorf("failed to get reservation: %w", err)
	}
	return &res, nil
}

func (r *reservationRepositoryImpl) RescheduleReservation(ctx context.Context, reservationID, newSlotID, newGuestCount int, notificationChannel string) error {
	query := "CALL reschedule_reservation(?, ?, ?, ?)"
	err := r.db.WithContext(ctx).Exec(query, reservationID, newSlotID, newGuestCount, notificationChannel).Error
	if err != nil {
		fmt.Printf("Error calling reschedule_reservation procedure for reservation ID %d: %v\n", reservationID, err)
		return fmt.Errorf("failed to reschedule reservation: %w", err)
	}
	return nil
}

func (r *reservationRepositoryImpl) FindAvailableTableForSlot(ctx context.Context, slotID int, guestCount int) (int, error) {
	var tableID int
	query := "SELECT find_available_table_for_slot(?, ?)"
	err := r.db.WithContext(ctx).Raw(query, slotID, guestCount).Scan(&tableID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || tableID == 0 {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to find available table: %w", err)
	}
	return tableID, nil
}

func (r *reservationRepositoryImpl) CancelReservation(ctx context.Context, reservationID int) error {
	err := r.DB.WithContext(ctx).Exec("SELECT cancel_reservation(?)", reservationID).Error
	return err
}

func (r *reservationRepositoryImpl) EditReservation(ctx context.Context, reservationID int, newGuestCount int) error {
	query := "CALL edit_reservation(?, ?)"
	err := r.DB.WithContext(ctx).Exec(query, reservationID, newGuestCount).Error
	if err != nil {
		fmt.Printf("Error calling edit_reservation procedure for reservation ID %d: %v\n", reservationID, err)
		return fmt.Errorf("failed to edit reservation: %w", err)
	}
	return nil
}
