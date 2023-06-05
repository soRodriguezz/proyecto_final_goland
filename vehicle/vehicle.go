package vehicle

import "time"

type Vehicle struct {
	ID              string
	LastService     time.Time
	ServiceInterval time.Duration
}

func NewVehicle(id string, lastService time.Time, interval time.Duration) *Vehicle {
	return &Vehicle{
		ID:              id,
		LastService:     lastService,
		ServiceInterval: interval,
	}
}

func (v *Vehicle) NeedsService(currentTime time.Time) bool {
	return currentTime.Sub(v.LastService) > v.ServiceInterval
}

type Maintenance struct {
	ID        string
	StartTime time.Time
	Duration  time.Duration
	Finished  bool
}

func NewMaintenance(id string, startTime time.Time, duration time.Duration) *Maintenance {
	return &Maintenance{
		ID:        id,
		StartTime: startTime,
		Duration:  duration,
		Finished:  false,
	}
}

func (m *Maintenance) FinishMaintenance() {
	m.Finished = true
}
