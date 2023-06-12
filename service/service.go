package service

import (
	"fmt"
	"proyecto_final_goland/utils"
	"strconv"
)

var (
	ServicesArr []*Service
	utilsImpl   = utils.UtilsImpl{}
)

type Service struct {
	Id    int
	Name  string
	Price int
}

func NewService(id int, name string, price int) *Service {
	return &Service{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

/*
 * Selecciona un servicio por ID
 */
func MaintenanceService() *Service {
	var idService string
	var serviceSelect Service

	ListServices()

	fmt.Print("\nIntroduzca el ID del servicio a seleccionar: ")
	fmt.Scanln(&idService)

	id, err := strconv.Atoi(idService)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return nil
	}

	for i := 0; i < len(ServicesArr); i++ {
		if ServicesArr[i].Id == id {
			serviceSelect = *ServicesArr[i]
			break
		}
	}

	return &Service{
		Id:    serviceSelect.Id,
		Name:  serviceSelect.Name,
		Price: serviceSelect.Price,
	}
}

/*
 * Inserta registros de servicios al iniciar la app
 */
func ServicesInit() []*Service {

	servicioUno := NewService(1, "Cambio de aceite", 50000)
	servicioDos := NewService(2, "Rotación de neumáticos", 30000)
	servicioTres := NewService(3, "Alineación de ruedas", 14000)
	servicioCuatro := NewService(4, "Cambio de bujías", 11000)
	servicioCinco := NewService(5, "Cambio de batería", 5000)
	servicioSeis := NewService(6, "Revisión de frenos", 34000)

	ServicesArr = append(ServicesArr, servicioUno, servicioDos, servicioTres, servicioCuatro, servicioCinco, servicioSeis)

	return ServicesArr
}

/*
 * Listar todos los servicios
 */
func ListServices() {
	utilsImpl.ClearConsole()

	w := utilsImpl.CreateTabs()

	fmt.Println("Servicios de la empresa: ")
	fmt.Fprintln(w, "\nID\tNombre\tPrecio")
	for _, svc := range ServicesArr {
		fmt.Fprintf(w, "%d\t%s\t%d\n", svc.Id, svc.Name, svc.Price)
	}

	w.Flush()
}

/*
 * Crea un nuevo servicio
 */
func CreateServices() {
	utilsImpl.ClearConsole()

	var name string
	var price int

	fmt.Print("Introduzca el nombre del servicio: ")
	fmt.Scanln(&name)

	fmt.Print("Introduzca el precio del servicio: ")
	fmt.Scanln(&price)

	newServiceInput := NewService(ServicesArr[len(ServicesArr)-1].Id+1, name, price)
	ServicesArr = append(ServicesArr, newServiceInput)

}

/*
 * Elimina un servicio por ID
 */
func DeleteServices() {
	var idInput string
	flag := false

	ListServices()

	fmt.Print("\nIntroduzca el ID del servicio a eliminar: ")
	fmt.Scanln(&idInput)

	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	// Search for the object by its ID and delete it
	for i := 0; i < len(ServicesArr); i++ {
		if ServicesArr[i].Id == id {
			flag = true
			ServicesArr = append(ServicesArr[:i], ServicesArr[i+1:]...)
			break
		}
	}

	if flag {
		utilsImpl.ClearConsole()
		fmt.Printf("\nServicio con ID %d eliminado!\n", id)
	} else {
		utilsImpl.ClearConsole()
		fmt.Printf("\nServicio con ID %d no encontrado\n", id)
	}

	utilsImpl.PausedConsole()
	utilsImpl.ClearConsole()
}

/*
 * Crea un menu de opciones para los servicios
 */
func ServicesOptions() {
	utilsImpl.ClearConsole()

	var option string

	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Listar servicios")
		fmt.Println("2. Agregar servicio")
		fmt.Println("3. Eliminar Servicio")
		fmt.Println("4. Salir")
		fmt.Print("Ingrese su opción: ")

		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Ocurrió un error al leer la entrada. Por favor intente de nuevo.")
			continue
		}

		switch option {
		case "1":
			ListServices()
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		case "2":
			CreateServices()
			ListServices()
			fmt.Println("\nServicio agregado!")
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		case "3":
			DeleteServices()
		case "4":
			utilsImpl.ClearConsole()
			return
		default:
			fmt.Println("Opción inválida")
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		}
	}

}
