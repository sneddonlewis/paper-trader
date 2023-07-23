package model

type Portfolio struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
}
