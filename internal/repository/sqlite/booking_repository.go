package sqlite

import (
	"context"
	"database/sql"

	"github.com/DacioSB/booking-service/internal/domain"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(con *sql.DB) *BookingRepository {
	return &BookingRepository{db: con}
}

func (r *BookingRepository) Save(ctx context.Context, tx *sql.Tx, booking *domain.Booking) error {
	_, err := tx.ExecContext(ctx, 
		"INSERT INTO bookings (id, room_id, user_id, status) VALUES (?, ?, ?, ?)",
		booking.ID, booking.RoomID, booking.UserID, booking.Status,
	)
	return err
}