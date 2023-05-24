package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"proyecto_final_goland/customer"
	"proyecto_final_goland/scheduler"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/utils"
	"proyecto_final_goland/vehicle"
)

func main() {
	// Crear instancia de Customer
	cust := customer.NewCustomer("John Doe", "+1234567890")

	// Crear instancias de Vehicle
	vehicles := []*vehicle.Vehicle{
		vehicle.NewVehicle("V1", time.Now().Add(-time.Hour*25), time.Hour*24),
		vehicle.NewVehicle("V2", time.Now().Add(-time.Hour*25), time.Hour*24),
	}

	// Crear instancias de Service
	services := []*service.Service{
		service.NewService("Oil Change", 50.0),
		service.NewService("Tire Rotation", 30.0),
	}

	// Crear instancias de Shop
	shops := []*shop.Shop{
		shop.NewShop("AutoCare Center", "123 Main St"),
		shop.NewShop("Quick Fix Garage", "456 Elm St"),
	}

	// Crear instancia de Scheduler
	s := scheduler.NewScheduler(vehicles)

	// Programar mantenimiento
	s.ScheduleMaintenance()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Vehicle ID\tCost\tStart Time\tDuration\tFinished")

	// Mostrar mantenimiento programado
	for _, m := range s.Maintenance {
		fmt.Fprintf(w, "%s\t%.2f\t%s\t%.0f minutes\t%t\n", m.ID, m.Cost, m.StartTime.Format(time.RFC3339), m.Duration.Minutes(), m.Finished)
	}

	fmt.Fprintln(w, "")
	w.Flush()

	// Validación de entrada para agregar nuevo vehículo
	id, lastServiceTime, interval := utils.InputValidation()
	newVehicle := vehicle.NewVehicle(id, lastServiceTime, interval)
	s.Vehicles = append(s.Vehicles, newVehicle)

	// Programar mantenimiento para el nuevo vehículo
	s.ScheduleMaintenanceAt(newVehicle, time.Now().Add(time.Hour*2))

	// Finalizar mantenimiento
	s.FinishMaintenance(s.Maintenance[0])

	// Mostrar mantenimiento programado después de agregar nuevo vehículo y finalizar un mantenimiento
	fmt.Fprintln(w, "Vehicle ID\tCost\tStart Time\tDuration\tFinished")
	for _, m := range s.Maintenance {
		fmt.Fprintf(w, "%s\t%.2f\t%s\t%.0f minutes\t%t\n", m.ID, m.Cost, m.StartTime.Format(time.RFC3339), m.Duration.Minutes(), m.Finished)
	}

	fmt.Fprintln(w, "")
	w.Flush()

	// Calcular el costo total de mantenimiento
	fmt.Printf("Total maintenance cost: %.2f\n", s.TotalMaintenanceCost())

	// Mostrar información del cliente
	fmt.Printf("Customer: %s\n", cust.Name)
	fmt.Printf("Phone: %s\n", cust.Phone)

	// Mostrar servicios disponibles
	fmt.Println("Services:")
	for _, svc := range services {
		fmt.Printf("- %s ($%.2f)\n", svc.Name, svc.Price)
	}

	// Mostrar tiendas disponibles
	fmt.Println("Shops:")
	for _, sh := range shops {
		fmt.Printf("- %s (%s)\n", sh.Name, sh.Location)
	}
}
