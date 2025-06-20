package entity

import "time"

type Notification struct {
	ID            int       `gorm:"primaryKey;column:notification_id"`
	CustomerID    int       `gorm:"column:customer_id"`
	ReservationID int       `gorm:"column:reservation_id"`
	Message       string    `gorm:"column:message"`
	SentAt        time.Time `gorm:"column:sent_at"`
	Channel       string    `gorm:"column:channel"`
	IsSent        bool      `gorm:"column:is_sent"`
	Email         string    `gorm:"column:email"`
}
