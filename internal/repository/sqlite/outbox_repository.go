package sqlite

import (
	"context"
	"database/sql"

	"github.com/DacioSB/booking-service/internal/outbox"
)


type OutboxRepository struct {
	db *sql.DB
}

func NewOutboxRepository(con *sql.DB) *OutboxRepository {
	return &OutboxRepository{db: con}
}

func (r *OutboxRepository) Save(ctx context.Context, tx *sql.Tx, event *outbox.Event) error {
	_, err := tx.ExecContext(ctx, 
	"INSERT INTO outbox (id, status, payload) VALUES (?, ?, ?)",
	event.ID, event.Status, event.Payload)
	return err
}

func (r *OutboxRepository) FetchPending(ctx context.Context) ([]outbox.Event, error)  {
	rows, err := r.db.QueryContext(ctx, "SELECT id, payload FROM outbox WHERE status == 'PENDING'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []outbox.Event
	for rows.Next() {
		var event outbox.Event
		if err := rows.Scan(&event.ID, &event.Payload); err != nil {
			return nil, err
		}
		event.Status = "PENDING"
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (r *OutboxRepository) MarkProcessed(ctx context.Context, eventId string) error   {
	_, err := r.db.ExecContext(ctx, "UPDATE outbox SET STATUS = 'PROCESSED' WHERE id = ?", eventId)
	return err
}