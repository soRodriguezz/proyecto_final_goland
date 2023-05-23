package vehicle

import "time"

type Maintenance struct {
	ID        string
	Cost      float64
	StartTime time.Time
	Duration  time.Duration
}
