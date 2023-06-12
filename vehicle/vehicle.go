package vehicle

import "time"

var VehicleArr []*Vehicle

type Vehicle struct {
	ID              string
	LastService     time.Time
	ServiceInterval time.Duration
}

// NewVehicle crea una nueva instancia de Vehicle y la devuelve.
// id es el identificador del vehículo.
// timeService es la fecha y hora del último servicio.
// interval es el intervalo de tiempo entre servicios.
func NewVehicle(id string, timeService time.Time, interval time.Duration) *Vehicle {
	return &Vehicle{
		ID:              id,
		LastService:     timeService,
		ServiceInterval: interval,
	}
}
