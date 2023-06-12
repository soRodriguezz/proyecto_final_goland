package main

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/customer"
	"proyecto_final_goland/maintenance"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/utils"
)

func main() {
	utils := utils.UtilsImpl{}

	utils.ClearConsole()

	service.ServicesInit()
	shop.ShopInit()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("=== Sistema de agenda de mantención de vehículos ===")
		fmt.Println("1. Mantenimientos")
		fmt.Println("2. Servicios")
		fmt.Println("3. Tiendas")
		fmt.Println("4. Clientes")
		fmt.Println("5. Vehiculos")
		fmt.Println("6. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			maintenance.MaintenancesOpt()
		case "2":
			service.ServicesOpt()
		case "3":
			shop.ShopsOpt()
		case "4":
			customer.CustomerOpt()
		case "5":
		case "6":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
			utils.PausedConsole()
			utils.ClearConsole()
		}
	}

}
