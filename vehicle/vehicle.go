package vehicle

import "time"

var VehicleArr []Vehicle

type Vehicle struct {
	ID              string
	LastService     time.Time
	ServiceInterval time.Duration
}

func NewVehicle(id string, timeService time.Time, interval time.Duration) Vehicle {
	return Vehicle{
		ID:              id,
		LastService:     timeService,
		ServiceInterval: interval,
	}
}

func (v *Vehicle) NeedsService(currentTime time.Time) bool {
	return currentTime.Sub(v.LastService) > v.ServiceInterval
}


// func NewMaintenance(id string, startTime time.Time, duration time.Duration) *Maintenance {
// 	return &Maintenance{
// 		ID:        id,
// 		StartTime: startTime,
// 		Duration:  duration,
// 		Finished:  false,
// 	}
// }

// func (m *Maintenance) FinishMaintenance() {
// 	m.Finished = true
// }
