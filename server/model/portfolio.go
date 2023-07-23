package model

type Portfolio struct {
	ID              int32             `json:"id"`
	UserID          int32             `json:"user_id"`
	Name            string            `json:"name"`
	Value           float64           `json:"value"`
	OpenPositions   []*Position       `json:"open-positions"`
	ClosedPositions []*ClosedPosition `json:"closed-positions"`
}
