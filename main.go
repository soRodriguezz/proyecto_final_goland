package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"proyecto_final_goland/scheduler"
	"proyecto_final_goland/utils"
	"proyecto_final_goland/vehicle"
)

func main() {
	vehicles := []*vehicle.Vehicle{
		vehicle.NewVehicle("V1", time.Now().Add(-time.Hour*25), time.Hour*24),
		vehicle.NewVehicle("V2", time.Now().Add(-time.Hour*25), time.Hour*24),
	}

	s := scheduler.NewScheduler(vehicles)
	s.ScheduleMaintenance()

	w := new(tabwriter.Writer)

	// Inicializar tabwriter
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	// Escribir los encabezados
	fmt.Fprintln(w, "ID de Vehículo\tCosto\tHora de inicio\tDuración")

	// Mostrar todos los mantenimientos programados
	for _, m := range s.Maintenance {
		fmt.Fprintf(w, "%s\t%.2f\t%s\t%.0f minutos\n", m.ID, m.Cost, m.StartTime.Format(time.RFC3339), m.Duration.Minutes())
	}

	fmt.Fprintln(w, "") // Imprimir una línea en blanco
	w.Flush()           // Asegurar que todos los datos sean escritos

	// Permitir al usuario programar mantenimientos manualmente
	id, lastServiceTime, interval := utils.InputValidation()
	newVehicle := vehicle.NewVehicle(id, lastServiceTime, interval)
	s.Vehicles = append(s.Vehicles, newVehicle)

	// Programar mantenimiento para el nuevo vehículo
	s.ScheduleMaintenance()

	// Mostrar todos los mantenimientos programados después de añadir el nuevo vehículo
	fmt.Fprintln(w, "ID de Vehículo\tCosto\tHora de inicio\tDuración")
	for _, m := range s.Maintenance {
		fmt.Fprintf(w, "%s\t%.2f\t%s\t%.0f minutos\n", m.ID, m.Cost, m.StartTime.Format(time.RFC3339), m.Duration.Minutes())
	}

	fmt.Fprintln(w, "") // Imprimir una línea en blanco
	w.Flush()           // Asegurar que todos los datos sean escritos
}
