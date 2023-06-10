package customer

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/utils"
	"strings"
)

var CustomerArr []Customer

type Customer struct {
	Name  string
	Phone string
}

func NewCustomer(name string, phone string) Customer {
	return Customer{
		Name:  name,
		Phone: phone,
	}
}

/*
 * Lista todos los clientes sin mostrar los nombres repetidos
 */
func listCustomers() {
	utils.ClearConsole()
	if len(CustomerArr) != 0 {
		w := utils.CreateTabs()

		fmt.Println("Clientes en nuestros registros: ")
		fmt.Fprintln(w, "\nNombre\tTelefono")

		// Mapa para realizar un seguimiento de los nombres ya mostrados
		seenNames := make(map[string]bool)

		for _, cus := range CustomerArr {
			lowerName := strings.ToLower(cus.Name)

			if !seenNames[lowerName] {
				fmt.Fprintf(w, "%s\t%s\n", cus.Name, cus.Phone)
				seenNames[lowerName] = true
			}
		}
		w.Flush()
	} else {
		fmt.Println("No posee clientes :(")
	}
}

/*
 * Crea opciones de clientes
 */
func CustomerOpt() {
	utils.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Listar clientes")
		fmt.Println("2. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			listCustomers()
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			utils.ClearConsole()
			return
		default:
			fmt.Println("Opción inválida")
			utils.PausedConsole()
			utils.ClearConsole()
		}
	}

}
