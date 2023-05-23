package utils

import (
	"fmt"
	"time"
)

func TimeToDuration(t time.Time) time.Duration {
	now := time.Now()
	return now.Sub(t)
}

func InputValidation() (string, time.Time, time.Duration) {
	var id string
	var lastService string
	var intervalHours float64

	fmt.Println("Por favor, introduzca el ID del vehículo:")
	fmt.Scanln(&id)

	fmt.Println("Introduzca la fecha del último servicio (YYYY-MM-DD):")
	fmt.Scanln(&lastService)

	lastServiceTime, err := time.Parse("2006-01-02", lastService)
	if err != nil {
		fmt.Println("Fecha inválida, por favor intente de nuevo.")
		return InputValidation()
	}

	fmt.Println("Por favor, introduzca el intervalo de servicio en horas:")
	fmt.Scanln(&intervalHours)

	interval := time.Duration(intervalHours) * time.Hour

	return id, lastServiceTime, interval
}
