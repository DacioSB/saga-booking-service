package saga

type BookingCreatedEvent struct {
	BookingID string `json:"booking_id"`
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
}

type PaymentFailedEvent struct {
	BookingID string `json:"booking_id"`
}