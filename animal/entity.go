package animal

import "time"

type Animals struct {
	ID        int
	Name      string
	Class     string
	Legs      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
