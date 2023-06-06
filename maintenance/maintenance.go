package maintenance

import (
	"fmt"
	"proyecto_final_goland/utils"
	"proyecto_final_goland/vehicle"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/customer"
	"time"
	"bufio"
	"os"
	"strconv"
)

var MaintenanceArr []Maintenance

type Maintenance struct {
	Id              int
	Patente					string
	IntervalHours   float64
	ServiceTime     time.Time
	Interval        time.Duration
	NameCustomer    string
	PhoneCustomer   string
	Service					string
	Shop						string
	ServicePrice		int
	Finished        bool
}

func CreateMaintenance() {
	var patente string
	var intervalHoursResult float64
	var nameCustomerResult string
	var phoneCustomerResult string
	var id int = 1

	if len(MaintenanceArr) > 0  {
		id = MaintenanceArr[len(MaintenanceArr)-1].Id + 1
	}

	fmt.Print("Introduzca patente del vehículo: ")
	fmt.Scanln(&patente)

	serviceTimeResult := time.Now()

	fmt.Print("Introduzca el intervalo de servicio en horas: ")
	fmt.Scanln(&intervalHoursResult)

	fmt.Print("Introduzca su nombre para asociar el servicio: ")
	fmt.Scanln(&nameCustomerResult)

	fmt.Print("Introduzca su telefono para asociar a su servicio (9 caracteres): ")
	fmt.Scanln(&phoneCustomerResult)

	intervalResult := time.Duration(intervalHoursResult) * time.Hour

	// Obtener servicio seleccionado
	serviceSelect := service.MaintenanceService()

	// Obtener tienda seleccionada
	shopSelect := shop.MaintenanceShop()
	utils.ClearConsole()

	maintenance := Maintenance{
		Id:							 id,
		Patente:         patente,
		ServiceTime: 		 serviceTimeResult,
		Interval:        intervalResult,
		NameCustomer:    nameCustomerResult,
		PhoneCustomer:   phoneCustomerResult,
		Service:				 serviceSelect.Name,
		Shop:            shopSelect.Name,
		ServicePrice:    serviceSelect.Price,
		Finished:				 false,
	}

	MaintenanceArr = append(MaintenanceArr, maintenance)

	newVehicle := vehicle.NewVehicle(maintenance.Patente, maintenance.ServiceTime, maintenance.Interval)
	vehicle.VehicleArr = append(vehicle.VehicleArr, newVehicle)

	newCustomer := customer.NewCustomer(maintenance.NameCustomer, maintenance.PhoneCustomer)
	customer.CustomerArr = append(customer.CustomerArr, newCustomer)
}

func FinishMaintenance() {
	ListMaintenance()

	var id string

	fmt.Print("Introduzca el ID de mantención a finalizar: ")
	fmt.Scanln(&id)

	idInput, _ := strconv.Atoi(id)

	for i := 0; i < len(MaintenanceArr); i++ {
		if MaintenanceArr[i].Id == idInput {
			MaintenanceArr[i].Finished = true
			break
		}
	}

}

func ListMaintenance() {
	// Crea tabs
	w := utils.CreateTabs()

	// Mostrar mantenimiento programado después de agregar nuevo vehículo y finalizar un mantenimiento
	fmt.Fprintln(w, "\nID Vehiculo\tCosto\tHora Inicio\tDuración\tNombre\tTelefono\tServicio\tTienda")
	for _, m := range MaintenanceArr {
		if !m.Finished {
			fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%.0f minutos\t%s\t%s\t%s\t%s\n", m.Id, m.Patente, m.ServicePrice, m.ServiceTime.Format("02-01-2006 15:04:05"), m.Interval.Minutes(), m.NameCustomer, m.PhoneCustomer, m.Service, m.Shop)
		}
	}

	w.Flush()
}

func AllMaintenance() {
		// Crea tabs
		w := utils.CreateTabs()

		// Mostrar mantenimiento programado después de agregar nuevo vehículo y finalizar un mantenimiento
		fmt.Fprintln(w, "\nID Vehiculo\tFinalizado\tCosto\tHora Inicio\tDuración\tNombre\tTelefono\tServicio\tTienda")
		for _, m := range MaintenanceArr {
			var finished string = "NO"
			if m.Finished {
				finished = "SI"
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%.0f minutos\t%s\t%s\t%s\t%s\n", m.Id, m.Patente, finished, m.ServicePrice, m.ServiceTime.Format("02-01-2006 15:04:05"), m.Interval.Minutes(), m.NameCustomer, m.PhoneCustomer, m.Service, m.Shop)
		}
		w.Flush()
}

func Maintenances() {
	utils.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Ingresar manteminimiento")
		fmt.Println("2. Ver mantenimientos pendientes")
		fmt.Println("3. Finalizar mantenimiento")
		fmt.Println("4. Mostrar todos los mantenimientos")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			utils.ClearConsole()
			CreateMaintenance()
			fmt.Println("Mantenimiento agendado!")
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			utils.ClearConsole()
			ListMaintenance()
			utils.PausedConsole()
			utils.ClearConsole()
		case "3":
			utils.ClearConsole()
			FinishMaintenance()
			utils.PausedConsole()
			utils.ClearConsole()
		case "4":
			utils.ClearConsole()
			AllMaintenance()
			utils.PausedConsole()
			utils.ClearConsole()
		case "5":
			utils.ClearConsole()
			return
		default:
			fmt.Println("Opción inválida")
			utils.PausedConsole()
			utils.ClearConsole()
		}

		fmt.Println()
	}

}
