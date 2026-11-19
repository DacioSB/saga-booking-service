package outbox

type Publisher struct {
	repo Repository
	messaging MessagingPublisher
}

type MessagingPublisher interface {
	Publish(queue string, payload []byte) error
}