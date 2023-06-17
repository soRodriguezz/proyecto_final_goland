package models

import (
	"proyecto_final_goland/maintenance"
	"time"
)

func NewMaintenance(patente string, intervalo float64, nameCus string, phoneCus string) *maintenance.Maintenance {
	return &maintenance.Maintenance{
		Id:            9999,
		Patente:       patente,
		IntervalHours: intervalo,
		ServiceTime:   time.Now(),
		Interval:      time.Duration(intervalo) * time.Hour,
		NameCustomer:  nameCus,
		PhoneCustomer: phoneCus,
		Service:       "Cambio de aceite",
		Shop:          "Taller Mecánico 'Mecánica Rápida'",
		ServicePrice:  200,
		Finished:      false,
	}
}
