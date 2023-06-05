package shop

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/utils"
	"strconv"
)

var ShopsArr []*shop

type shop struct {
	Id       int
	Name     string
	Location string
}

func newShop(id int, name string, location string) *shop {
	return &shop{
		Id:       id,
		Name:     name,
		Location: location,
	}
}

func MaintenanceShop() shop {
	var idShop string
	var shopSelect *shop

	listShop()

	fmt.Print("\nIntroduzca el ID de la tienda a seleccionar: ")
	fmt.Scanln(&idShop)

	id, _ := strconv.Atoi(idShop)

	for i := 0; i < len(ShopsArr); i++ {
		if ShopsArr[i].Id == id {
			shopSelect = ShopsArr[i]
			break
		}
	}

	return shop {
		Id: shopSelect.Id,
		Name: shopSelect.Name,
		Location: shopSelect.Location,
	}
}

func ShopConstuct() []*shop {

	shopUno := newShop(1, "Centro de Servicio Automotriz 'Autotec'", "Santiago, Chile")
	shopDos := newShop(2, "Taller Mecánico 'Mecánica Rápida'", "Valparaíso, Chile")
	shopTres := newShop(3, "Garaje 'Mantenimiento Total'", "Concepción, Chile")
	shopCuatro := newShop(4, "Centro de Mantenimiento y Reparación 'AutoCare'", "Temuco, Chile")
	shopCinco := newShop(5, "Tienda de Servicio Rápido 'Soluciones Automotrices'", "Antofagasta, Chile")
	shopSeis := newShop(6, "Mecánica 'El Motorista'", "Viña del Mar, Chile")
	shopSiete := newShop(7, "Taller de Servicio Técnico 'La Excelencia'", "La Serena, Chile")

	ShopsArr = append(ShopsArr, shopUno)
	ShopsArr = append(ShopsArr, shopDos)
	ShopsArr = append(ShopsArr, shopTres)
	ShopsArr = append(ShopsArr, shopCuatro)
	ShopsArr = append(ShopsArr, shopCinco)
	ShopsArr = append(ShopsArr, shopSeis)
	ShopsArr = append(ShopsArr, shopSiete)

	return ShopsArr
}

func listShop() {
	utils.ClearConsole()

	// Crea tabs
	w := utils.CreateTabs()

	// Mostrar servicios disponibles
	fmt.Println("Tiendas operativas: ")
	fmt.Fprintln(w, "\nID\tNombre\tDirección")
	for _, svc := range ShopsArr {
		fmt.Fprintf(w, "%d\t%s\t%s\n", svc.Id, svc.Name, svc.Location)
	}

	w.Flush()
}

func createShop() {
	utils.ClearConsole()

	var name string
	var location string

	fmt.Print("Introduzca el nombre de la tienda: ")
	fmt.Scanln(&name)

	fmt.Print("Introduzca la ubicación de la tienda: ")
	fmt.Scanln(&location)

	newShopInput := newShop(ShopsArr[len(ShopsArr)-1].Id+1, name, location)
	ShopsArr = append(ShopsArr, newShopInput)

}

func deleteShop() {
	var idInput string
	flag := false

	listShop()

	fmt.Print("\nIntroduzca el ID de la tienda a eliminar: ")
	fmt.Scanln(&idInput)

	id, _ := strconv.Atoi(idInput)

	// Buscar el objeto por su ID y eliminarlo
	for i := 0; i < len(ShopsArr); i++ {
		if ShopsArr[i].Id == id {
			flag = true
			ShopsArr = append(ShopsArr[:i], ShopsArr[i+1:]...)
			break
		}
	}

	if flag {
		utils.ClearConsole()
		fmt.Printf("\nTienda con ID %d eliminada!\n", id)
	} else {
		utils.ClearConsole()
		fmt.Printf("\nTienda con ID %d no encontrada\n", id)
	}

	utils.PausedConsole()
	utils.ClearConsole()
}

func Shops() {
	utils.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

	// Crea menú para opciones de la aplicación
	for {
		fmt.Println("Seleccione una opción: ")
		fmt.Println("1. Listar tiendas")
		fmt.Println("2. Agregar tienda")
		fmt.Println("3. Eliminar tienda")
		fmt.Println("4. Salir")
		fmt.Print("Ingrese su opción: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			listShop()
			utils.PausedConsole()
			utils.ClearConsole()
		case "2":
			createShop()
			listShop()
			utils.PausedConsole()
			utils.ClearConsole()
		case "3":
			deleteShop()
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
