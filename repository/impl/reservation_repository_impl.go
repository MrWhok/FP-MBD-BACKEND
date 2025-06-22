package impl

import (
	"context"
	"errors"
	"fmt"

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
		// Optionally parse PG error here
		return errors.New("failed to create reservation: " + err.Error())
	}
	return nil
}

// GetReservationByID fetches a reservation by its ID using GORM's Raw method.
func (r *reservationRepositoryImpl) GetReservationByID(ctx context.Context, reservationID int) (*model.Reservation, error) {
	var res model.Reservation
	query := `SELECT reservation_id, customer_id, slot_id, table_id, guest_count, status, created_at
			  FROM reservation WHERE reservation_id = ?` // GORM uses '?' placeholders
	err := r.db.WithContext(ctx).Raw(query, reservationID).Scan(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Check for GORM's specific "no rows" error
			return nil, fmt.Errorf("reservation with ID %d not found", reservationID)
		}
		return nil, fmt.Errorf("failed to get reservation: %w", err)
	}
	return &res, nil
}

// RescheduleReservation calls the PostgreSQL stored procedure 'reschedule_reservation' using GORM's Exec.
func (r *reservationRepositoryImpl) RescheduleReservation(ctx context.Context, reservationID, newSlotID, newGuestCount int, notificationChannel string) error {
	// Using CALL for stored procedure in PostgreSQL 11+
	query := "CALL reschedule_reservation(?, ?, ?, ?)" // GORM uses '?' placeholders
	err := r.db.WithContext(ctx).Exec(query, reservationID, newSlotID, newGuestCount, notificationChannel).Error
	if err != nil {
		// Log error more detailed for debugging
		fmt.Printf("Error calling reschedule_reservation procedure for reservation ID %d: %v\n", reservationID, err)
		return fmt.Errorf("failed to reschedule reservation: %w", err)
	}
	return nil
}

// FindAvailableTableForSlot calls the PostgreSQL function 'find_available_table_for_slot' using GORM's Raw.
func (r *reservationRepositoryImpl) FindAvailableTableForSlot(ctx context.Context, slotID int, guestCount int) (int, error) {
	var tableID int
	query := "SELECT find_available_table_for_slot(?, ?)" // GORM uses '?' placeholders
	err := r.db.WithContext(ctx).Raw(query, slotID, guestCount).Scan(&tableID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || tableID == 0 { // Check if no row or if function returns 0/NULL
			return 0, nil // Return 0 if no table is found (function returns NULL/0)
		}
		return 0, fmt.Errorf("failed to find available table: %w", err)
	}
	return tableID, nil
}
