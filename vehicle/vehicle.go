package vehicle

import "time"

var VehicleArr []*Vehicle

type Vehicle struct {
	ID              string
	LastService     time.Time
	ServiceInterval time.Duration
}

func NewVehicle(id string, timeService time.Time, interval time.Duration) *Vehicle {
	return &Vehicle{
		ID:              id,
		LastService:     timeService,
		ServiceInterval: interval,
	}
}
