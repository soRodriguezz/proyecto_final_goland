package main

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/customer"
	"proyecto_final_goland/maintenance"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/usecases"
	"proyecto_final_goland/utils"
	"sync"
)

var utilsImpl utils.Utils

func useCaseCreateMaintenance() {
	usecases.CreateMaintenanceUseCase()
}

// El proyecto se trata de una aplicación para manejar la programación de mantenimientos de vehículos.
func main() {
	utilsImpl = utils.NewUtils()

	utilsImpl.ClearConsole()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		service.InitServices()
		wg.Done()
	}()

	go func() {
		shop.InitShops()
		wg.Done()
	}()

	wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== Sistema de agenda de mantención de vehículos ===")
		fmt.Println("1. Mantenimientos")
		fmt.Println("2. Servicios")
		fmt.Println("3. Tiendas")
		fmt.Println("4. Clientes")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			maintenance.MaintenancesOpt()
		case "2":
			service.ServicesOptions()
		case "3":
			shop.ShopsOptions()
		case "4":
			customer.CustomerOptions()
		case "5":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		}
	}
}
