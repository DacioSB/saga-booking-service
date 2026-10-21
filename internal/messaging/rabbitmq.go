package messaging

import "github.com/streadway/amqp"


type RabbitMQPublisher struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func NewRabbitMQPublisher(url string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	//ensure queue exists
	_, err = channel.QueueDeclare("booking_events", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &RabbitMQPublisher{conn: conn, ch: channel}, nil
}

func (p *RabbitMQPublisher) Publish(queue string, payload []byte) error {
	return p.ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: payload,
		},
	)
}

func (p *RabbitMQPublisher) Close()  {
	p.ch.Close()
	p.conn.Close()
}