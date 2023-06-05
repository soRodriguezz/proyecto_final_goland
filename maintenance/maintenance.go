package maintenance

import (
	"fmt"
	"proyecto_final_goland/customer"
	"proyecto_final_goland/scheduler"
	"proyecto_final_goland/utils"
	"proyecto_final_goland/vehicle"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"time"
)

type Maintenance struct {
	Id              string
	LastService     string
	IntervalHours   float64
	LastServiceTime time.Time
	Interval        time.Duration
	NameCustomer    string
	PhoneCustomer   string
}

func NewMaintenance() Maintenance {
	var idResult string
	var lastServiceResult string
	var intervalHoursResult float64
	var nameCustomerResult string
	var phoneCustomerResult string

	fmt.Print("Introduzca el ID del vehículo: ")
	fmt.Scanln(&idResult)

	fmt.Print("Introduzca la fecha del último servicio (YYYY-MM-DD): ")
	fmt.Scanln(&lastServiceResult)

	lastServiceTimeResult, err := time.Parse("2006-01-02", lastServiceResult)
	if err != nil {
		fmt.Println("Fecha inválida, intente nuevamente.")
		return NewMaintenance()
	}

	fmt.Print("Introduzca el intervalo de servicio en horas: ")
	fmt.Scanln(&intervalHoursResult)

	fmt.Print("Introduzca su nombre para asociar el servicio: ")
	fmt.Scanln(&nameCustomerResult)

	fmt.Print("Introduzca su telefono para asociar a su servicio (9 caracteres): ")
	fmt.Scanln(&phoneCustomerResult)

	if len(phoneCustomerResult) != 9 {
		fmt.Println("El campo no tiene 9 caracteres.")
		return NewMaintenance()
	}

	intervalResult := time.Duration(intervalHoursResult) * time.Hour

	return Maintenance{
		Id:              idResult,
		LastServiceTime: lastServiceTimeResult,
		Interval:        intervalResult,
		NameCustomer:    nameCustomerResult,
		PhoneCustomer:   phoneCustomerResult,
	}
}

func CreateMaintenance() {
	utils.ClearConsole()
	// Crea instancias de Vehicle
	vehicles := []*vehicle.Vehicle{}

	// Crea instancia de Scheduler
	s := scheduler.NewScheduler(vehicles)

	// Validación de entrada para agregar nuevo vehículo
	input := NewMaintenance()

	// Obtener servicio seleccionado
	serviceSelect := service.MaintenanceService()

	// Obtener tienda seleccionada
	shopSelect := shop.MaintenanceShop()
	utils.ClearConsole()

	cust := customer.NewCustomer(input.NameCustomer, input.PhoneCustomer)
	newVehicle := vehicle.NewVehicle(input.Id, input.LastServiceTime, input.Interval)
	s.Vehicles = append(s.Vehicles, newVehicle)

	// Programa mantenimiento para el nuevo vehículo
	s.ScheduleMaintenanceAt(newVehicle, time.Now())

	// Crea tabs
	w := utils.CreateTabs()

	// Mostrar mantenimiento programado después de agregar nuevo vehículo y finalizar un mantenimiento
	fmt.Fprintln(w, "\nID Vehiculo\tCosto\tHora Inicio\tDuración\tNombre\tTelefono\tServicio\tTienda")
	for _, m := range s.Maintenance {
		fmt.Fprintf(w, "%s\t%d\t%s\t%.0f minutes\t%s\t%s\t%s\t%s\n", m.ID, serviceSelect.Price, m.StartTime.Format("02-01-2006 15:04:05"), m.Duration.Minutes(), cust.Name, cust.Phone, serviceSelect.Name, shopSelect.Name)
	}

	w.Flush()

	utils.PausedConsole()
	utils.ClearConsole()
}
