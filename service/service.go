package service

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/utils"
	"strconv"
)

var ServicesArr []*service

type service struct {
	Id    int
	Name  string
	Price int
}

func newService(id int, name string, price int) *service {
	return &service{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

/*
 * Selecciona un objeto de servicio por ID
 */
func MaintenanceService() *service {
	var idService string
	var serviceSelect service

	listServices()

	fmt.Print("\nIntroduzca el ID del servicio a seleccionar: ")
	fmt.Scanln(&idService)

	id, _ := strconv.Atoi(idService)

	for i := 0; i < len(ServicesArr); i++ {
		if ServicesArr[i].Id == id {
			serviceSelect = *ServicesArr[i]
			break
		}
	}

	return &service{
		Id:    serviceSelect.Id,
		Name:  serviceSelect.Name,
		Price: serviceSelect.Price,
	}
}

/*
 * Llenar colección con servicios al iniciar app
 */
func ServicesInit() []*service {

	servicioUno := newService(1, "Cambio de aceite", 50000)
	servicioDos := newService(2, "Rotación de neumáticos", 30000)
	servicioTres := newService(3, "Alineación de ruedas", 14000)
	servicioCuatro := newService(4, "Cambio de bujías", 11000)
	servicioCinco := newService(5, "Cambio de batería", 5000)
	servicioSeis := newService(6, "Revisión de frenos", 34000)

	ServicesArr = append(ServicesArr, servicioUno)
	ServicesArr = append(ServicesArr, servicioDos)
	ServicesArr = append(ServicesArr, servicioTres)
	ServicesArr = append(ServicesArr, servicioCuatro)
	ServicesArr = append(ServicesArr, servicioCinco)
	ServicesArr = append(ServicesArr, servicioSeis)

	return ServicesArr
}

/*
 * Lista todos los servicios
 */
func listServices() {
	utils.ClearConsole()

	w := utils.CreateTabs()

	fmt.Println("Servicios de la empresa: ")
	fmt.Fprintln(w, "\nID\tNombre\tPrecio")
	for _, svc := range ServicesArr {
		fmt.Fprintf(w, "%d\t%s\t%d\n", svc.Id, svc.Name, svc.Price)
	}

	w.Flush()
}

/*
 * Agrega un nuevo servicio en la colección
 */
func createServices() {
	utils.ClearConsole()

	var name string
	var price int

	fmt.Print("Introduzca el nombre del servicio: ")
	fmt.Scanln(&name)

	fmt.Print("Introduzca el precio del servicio: ")
	fmt.Scanln(&price)

	newServiceInput := newService(ServicesArr[len(ServicesArr)-1].Id+1, name, price)
	ServicesArr = append(ServicesArr, newServiceInput)

}

/*
 * Hace un hard delete de un servicio por ID
 */
func deleteServices() {
	var idInput string
	flag := false

	listServices()

	fmt.Print("\nIntroduzca el ID del servicio a eliminar: ")
	fmt.Scanln(&idInput)

	id, _ := strconv.Atoi(idInput)

	// Buscar el objeto por su ID y eliminarlo
	for i := 0; i < len(ServicesArr); i++ {
		if ServicesArr[i].Id == id {
			flag = true
			ServicesArr = append(ServicesArr[:i], ServicesArr[i+1:]...)
			break
		}
	}

	if flag {
		utils.ClearConsole()
		fmt.Printf("\nServicio con ID %d eliminado!\n", id)
	} else {
		utils.ClearConsole()
		fmt.Printf("\nServicio con ID %d no encontrado\n", id)
	}

	utils.PausedConsole()
	utils.ClearConsole()
}

/*
 * Crea opciones de servicios
 */
func ServicesOpt() {
	utils.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Listar servicios")
		fmt.Println("2. Agregar servicio")
		fmt.Println("3. Eliminar Servicio")
		fmt.Println("4. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			listServices()
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			createServices()
			listServices()
			fmt.Println("\nServicio agregado!")
			utils.PausedConsole()
			utils.ClearConsole()
		case "3":
			deleteServices()
		case "4":
			utils.ClearConsole()
			return
		default:
			fmt.Println("Opción inválida")
			utils.PausedConsole()
			utils.ClearConsole()
		}
	}

}
