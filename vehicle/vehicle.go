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
	Cost      float64
	StartTime time.Time
	Duration  time.Duration
	Finished  bool
}

func NewMaintenance(id string, cost float64, startTime time.Time, duration time.Duration) *Maintenance {
	return &Maintenance{
		ID:        id,
		Cost:      cost,
		StartTime: startTime,
		Duration:  duration,
		Finished:  false,
	}
}

func (m *Maintenance) FinishMaintenance() {
	m.Finished = true
}
