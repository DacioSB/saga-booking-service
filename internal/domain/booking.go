package domain

type Booking struct {
	ID string `json:"id"`
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
	Status string `json:"status"`
}