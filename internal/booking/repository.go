package booking

import (
	"context"

	"github.com/DacioSB/booking-service/internal/domain"
	"github.com/DacioSB/booking-service/internal/transaction"
)

// Tx is an alias retained for compatibility with existing implementations.
type Tx = transaction.Tx

type Repository interface {
	Save(ctx context.Context, tx Tx, booking *domain.Booking) error
}
