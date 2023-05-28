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

	scanner := bufio.NewScanner(os.Stdin)

	// Crea menú para opciones de la aplicación
	for {
		fmt.Println("=== Sistema de gestión de mantenimiento de flota de vehículos ===")
		fmt.Println("1. Programar mantenimiento")
		fmt.Println("2. Servicios")
		fmt.Println("3. Ver tiendas operativas")
		fmt.Println("4. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			maintenance.CreateMaintenance()
		case "2":
			service.Services()
		case "3":
			shop.CreateShop()
		case "4":
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
