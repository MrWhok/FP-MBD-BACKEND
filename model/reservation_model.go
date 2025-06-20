package model

type ReservationRequest struct {
	SlotID     int `json:"slot_id" validate:"required"`
	TableID    int `json:"table_id" validate:"required"`
	GuestCount int `json:"guest_count" validate:"required"`
}
