package booking

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/DacioSB/booking-service/internal/domain"
	"github.com/DacioSB/booking-service/internal/outbox"
	"github.com/DacioSB/booking-service/internal/saga"
)

type sqlTxAdapter struct {
	*sql.Tx
}

type Service struct {
	db         *sql.DB
	bookinRepo Repository
	outboxRepo outbox.Repository
}

func NewService(db *sql.DB, bRepo Repository, outRepo outbox.Repository) *Service {
	return &Service{
		db:         db,
		bookinRepo: bRepo,
		outboxRepo: outRepo,
	}
}

// NewBooking saves a reserved booking and its outbox event atomically.
func (s *Service) NewBooking(ctx context.Context, bookingRequest BookingRequest) (*domain.Booking, error) {
	bookingID := fmt.Sprintf("book-%s-%s", bookingRequest.RoomID, bookingRequest.UserID)
	newBooking := &domain.Booking{
		ID:     bookingID,
		RoomID: bookingRequest.RoomID,
		UserID: bookingRequest.UserID,
		Status: "RESERVED",
	}

	newEvent := saga.BookingCreatedEvent{
		BookingID: bookingID,
		RoomID:    bookingRequest.RoomID,
		UserID:    bookingRequest.UserID,
	}

	eventBytes, err := json.Marshal(newEvent)
	if err != nil {
		return nil, fmt.Errorf("marshal booking created event: %w", err)
	}

	outboxEvent := &outbox.Event{
		ID:      fmt.Sprintf("evt-%s", bookingID),
		Payload: eventBytes,
		Status:  "PENDING",
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin booking transaction: %w", err)
	}

	// Rollback handles failure paths and is harmless after a successful commit.
	defer func() {
		_ = tx.Rollback()
	}()

	sqlTx := &sqlTxAdapter{Tx: tx}

	if err := s.bookinRepo.Save(ctx, sqlTx, newBooking); err != nil {
		return nil, fmt.Errorf("save booking: %w", err)
	}

	if err := s.outboxRepo.Save(ctx, sqlTx, outboxEvent); err != nil {
		return nil, fmt.Errorf("save outbox event: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit booking transaction: %w", err)
	}

	return newBooking, nil
}
