package main

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/maintenance"
	"proyecto_final_goland/service"
	"proyecto_final_goland/shop"
	"proyecto_final_goland/utils"
)

func main() {
	utils.ClearConsole()

	service.ServicesConstuct()
	shop.ShopConstuct()

	scanner := bufio.NewScanner(os.Stdin)

	// Crea menú para opciones de la aplicación
	for {
		fmt.Println("=== Sistema de gestión de mantenimiento de flota de vehículos ===")
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
			maintenance.Maintenances()
		case "2":
			service.Services()
		case "3":
			shop.Shops()
		case "4":
			return
		case "5":
			return
		case "6":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
			utils.PausedConsole()
			utils.ClearConsole()
		}

		fmt.Println()
	}

}
