package outbox

import (
	"context"
	"log"
	"time"
)

type Publisher struct {
	repo Repository
	messaging MessagingPublisher
	ticker *time.Ticker
}

type MessagingPublisher interface {
	Publish(queue string, payload []byte) error
}

func NewPublisher(repository Repository, messaging MessagingPublisher) *Publisher  {
	return &Publisher{
		repo: repository,
		messaging: messaging,
		ticker: time.NewTicker(1 * time.Second),
	}
}

func (p *Publisher) Start(ctx context.Context) {
	log.Println("Starting Outbox Publisher...")
}