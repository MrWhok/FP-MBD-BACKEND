package model

import "time"

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ReservationRequest struct {
	SlotID     int `json:"slot_id" validate:"required"`
	TableID    int `json:"table_id" validate:"required"`
	GuestCount int `json:"guest_count" validate:"required"`
}

type RescheduleRequest struct {
	ReservationID int `json:"reservation_id" validate:"required"`
	NewSlotID     int `json:"new_slot_id" validate:"required"`
	NewGuestCount int `json:"new_guest_count" validate:"required"`
}

type Reservation struct {
	ReservationID int       `json:"reservation_id"`
	CustomerID    int       `json:"customer_id"`
	SlotID        int       `json:"slot_id"`
	TableID       int       `json:"table_id"`
	GuestCount    int       `json:"guest_count"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
