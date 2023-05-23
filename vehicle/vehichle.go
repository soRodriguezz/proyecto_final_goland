package vehicle

import (
	"time"
)

type Vehicle struct {
	ID              string
	LastServiceDate time.Time
	ServiceInterval time.Duration
}

func NewVehicle(id string, lastServiceDate time.Time, serviceInterval time.Duration) *Vehicle {
	return &Vehicle{
		ID:              id,
		LastServiceDate: lastServiceDate,
		ServiceInterval: serviceInterval,
	}
}

func (v *Vehicle) NeedsService(currentTime time.Time) bool {
	return currentTime.Sub(v.LastServiceDate) > v.ServiceInterval
}
