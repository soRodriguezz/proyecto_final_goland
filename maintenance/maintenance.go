package maintenance

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/customer"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/utils"
	"proyecto_final_goland/vehicle"
	"strconv"
	"time"
)

var MaintenanceArr []*Maintenance
var utilsImpl utils.Utils

type Maintenance struct {
	Id            int
	Patente       string
	IntervalHours float64
	ServiceTime   time.Time
	Interval      time.Duration
	NameCustomer  string
	PhoneCustomer string
	Service       string
	Shop          string
	ServicePrice  int
	Finished      bool
}

func init() {
	utilsImpl = utils.NewUtils()
}

/*
 * Crear nuevo mantenimiento
 */
func createMaintenance() {
	var patente string
	var intervalHoursResult float64
	var nameCustomerResult string
	var phoneCustomerResult string
	var id int = 1

	if len(MaintenanceArr) > 0 {
		id = MaintenanceArr[len(MaintenanceArr)-1].Id + 1
	}

	fmt.Print("Introduzca patente del vehículo: ")
	fmt.Scanln(&patente)

	serviceTimeResult := time.Now()

	fmt.Print("Introduzca el intervalo de servicio en horas: ")
	fmt.Scanln(&intervalHoursResult)

	fmt.Print("Introduzca su nombre para asociar el servicio: ")
	fmt.Scanln(&nameCustomerResult)

	fmt.Print("Introduzca su telefono para asociar a su servicio: ")
	fmt.Scanln(&phoneCustomerResult)

	intervalResult := time.Duration(intervalHoursResult) * time.Hour

	serviceSelect := service.MaintenanceService()

	shopSelect := *shop.SelectShop()
	utilsImpl.ClearConsole()

	maintenance := &Maintenance{
		Id:            id,
		Patente:       patente,
		ServiceTime:   serviceTimeResult,
		Interval:      intervalResult,
		NameCustomer:  nameCustomerResult,
		PhoneCustomer: phoneCustomerResult,
		Service:       serviceSelect.Name,
		Shop:          shopSelect.Name,
		ServicePrice:  serviceSelect.Price,
		Finished:      false,
	}

	MaintenanceArr = append(MaintenanceArr, maintenance)

	newVehicle := vehicle.NewVehicle(maintenance.Patente, maintenance.ServiceTime, maintenance.Interval)
	vehicle.VehicleArr = append(vehicle.VehicleArr, newVehicle)

	newCustomer := customer.NewCustomer(maintenance.NameCustomer, maintenance.PhoneCustomer)
	customer.CustomerArr = append(customer.CustomerArr, newCustomer)
}

/*
 * Finzalizar un mantenimiento por id
 */
func finishMaintenance() {
	pendingMaintenance()

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

/*
 * Listar mantenciones pendientes
 */
func pendingMaintenance() {
	utils := utils.UtilsImpl{}
	w := utils.CreateTabs()

	fmt.Fprintln(w, "\nID Vehiculo\tCosto\tHora Inicio\tDuración\tNombre\tTelefono\tServicio\tTienda")
	for _, m := range MaintenanceArr {
		if !m.Finished {
			fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%.0f minutos\t%s\t%s\t%s\t%s\n", m.Id, m.Patente, m.ServicePrice, m.ServiceTime.Format("02-01-2006 15:04:05"), m.Interval.Minutes(), m.NameCustomer, m.PhoneCustomer, m.Service, m.Shop)
		}
	}

	w.Flush()
}

/*
 * Listar todas las mantenciones
 */
func allMaintenance() {
	utils := utils.UtilsImpl{}
	w := utils.CreateTabs()

	fmt.Fprintln(w, "Patente\tFinalizado\tCosto\tHora Inicio\tDuración\tNombre\tTelefono\tServicio\tTienda")
	for _, m := range MaintenanceArr {
		var finished string = "NO"
		if m.Finished {
			finished = "SI"
		}
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%.0f minutos\t%s\t%s\t%s\t%s\n", m.Patente, finished, m.ServicePrice, m.ServiceTime.Format("02-01-2006 15:04:05"), m.Interval.Minutes(), m.NameCustomer, m.PhoneCustomer, m.Service, m.Shop)
	}
	w.Flush()
}

/*
 * Opciones de mantenciones
 */
func MaintenancesOpt() {
	utils := utils.UtilsImpl{}
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
			createMaintenance()
			fmt.Println("Mantenimiento agendado!")
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			utils.ClearConsole()
			pendingMaintenance()
			utils.PausedConsole()
			utils.ClearConsole()
		case "3":
			utils.ClearConsole()
			finishMaintenance()
			utils.PausedConsole()
			utils.ClearConsole()
		case "4":
			utils.ClearConsole()
			allMaintenance()
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
	}

}
