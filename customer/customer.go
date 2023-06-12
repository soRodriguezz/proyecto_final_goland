package customer

import (
	"fmt"
	"proyecto_final_goland/utils"
	"strings"
)

type CustomerUtils interface {
	ListCustomers()
	CustomerOptions()
}

var (
	CustomerArr []Customer
	utilsImpl   = utils.UtilsImpl{}
)

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

// ListCustomers lists all customers without displaying repeated names.
func ListCustomers() {
	utilsImpl.ClearConsole()
	if len(CustomerArr) != 0 {
		w := utilsImpl.CreateTabs()

		fmt.Println("Clientes en nuestros registros: ")
		fmt.Fprintln(w, "\nNombre\tTelefono")

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

// CustomerOptions creates customer options.
func CustomerOptions() {
	utilsImpl.ClearConsole()

	var option string

	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Listar clientes")
		fmt.Println("2. Salir")
		fmt.Print("Ingrese su opción: ")

		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Ocurrió un error al leer la entrada. Por favor intente de nuevo.")
			continue
		}

		switch option {
		case "1":
			ListCustomers()
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		case "2":
			utilsImpl.ClearConsole()
			return
		default:
			fmt.Println("Opción inválida")
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		}
	}
}
