package service

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/utils"
	"strconv"
)

var ServicesArr []Service

type Service struct {
	Id    int
	Name  string
	Price int
}

func NewService(id int, name string, price int) Service {
	return Service{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func MaintenanceService() Service {
	var idService string
	var serviceSelect Service

	ListServices()

	fmt.Print("\nIntroduzca el ID del servicio a seleccionar: ")
	fmt.Scanln(&idService)

	id, _ := strconv.Atoi(idService)

	for i := 0; i < len(ServicesArr); i++ {
		if ServicesArr[i].Id == id {
			serviceSelect = ServicesArr[i]
			break
		}
	}

	return Service {
		Id: serviceSelect.Id,
		Name: serviceSelect.Name,
		Price: serviceSelect.Price,
	}
}

func ServicesConstuct() []Service {

	servicioUno := NewService(1, "Cambio de aceite", 50000)
	servicioDos := NewService(2, "Rotación de neumáticos", 30000)
	servicioTres := NewService(3, "Alineación de ruedas", 14000)
	servicioCuatro := NewService(4, "Cambio de bujías", 11000)
	servicioCinco := NewService(5, "Cambio de batería", 5000)
	servicioSeis := NewService(6, "Revisión de frenos", 34000)

	ServicesArr = append(ServicesArr, servicioUno)
	ServicesArr = append(ServicesArr, servicioDos)
	ServicesArr = append(ServicesArr, servicioTres)
	ServicesArr = append(ServicesArr, servicioCuatro)
	ServicesArr = append(ServicesArr, servicioCinco)
	ServicesArr = append(ServicesArr, servicioSeis)

	return ServicesArr
}

func ListServices() {
	utils.ClearConsole()

	// Crea tabs
	w := utils.CreateTabs()

	// Mostrar servicios disponibles
	fmt.Println("Seleccione el servicio: ")
	fmt.Fprintln(w, "\nID\tNombre\tPrecio")
	for _, svc := range ServicesArr {
		fmt.Fprintf(w, "%d\t%s\t%d\n", svc.Id, svc.Name, svc.Price)
	}

	w.Flush()
}

func createServices() {
	utils.ClearConsole()

	var name string
	var price int

	fmt.Print("Introduzca el nombre del servicio: ")
	fmt.Scanln(&name)

	fmt.Print("Introduzca el precio del servicio: ")
	fmt.Scanln(&price)

	newServiceInput := NewService(ServicesArr[len(ServicesArr)-1].Id+1, name, price)
	ServicesArr = append(ServicesArr, newServiceInput)

}

func deleteServices() {
	var idInput string
	flag := false

	ListServices()

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

func Services() {
	utils.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

	// Crea menú para opciones de la aplicación
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
			ListServices()
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			createServices()
			ListServices()
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

		fmt.Println()
	}

}
