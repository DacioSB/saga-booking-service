package outbox

import (
	"context"

	"github.com/DacioSB/booking-service/internal/transaction"
)

type Event struct {
	ID      string
	Status  string
	Payload []byte
}

type Repository interface {
	Save(ctx context.Context, tx transaction.Tx, event *Event) error
	FetchPending(ctx context.Context) ([]Event, error)
	MarkProcessed(ctx context.Context, eventID string) error
}
