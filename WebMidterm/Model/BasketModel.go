package Model

import "time"

type Basket struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string
	Status    string
}
