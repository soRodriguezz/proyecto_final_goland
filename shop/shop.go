package shop

import (
	"bufio"
	"fmt"
	"os"
	"proyecto_final_goland/utils"
	"strconv"
)

var (
	ShopsArr  []*Shop
	utilsImpl = utils.UtilsImpl{}
)

type Shop struct {
	Id       int
	Name     string
	Location string
}

func NewShop(id int, name, location string) *Shop {
	return &Shop{
		Id:       id,
		Name:     name,
		Location: location,
	}
}

// SelectShop gets the object of a shop by ID
func SelectShop() *Shop {
	var idShop string
	var shopSelect *Shop

	ListShops()

	fmt.Print("\nIntroduzca el ID de la tienda a seleccionar: ")
	fmt.Scanln(&idShop)

	id, err := strconv.Atoi(idShop)
	if err != nil {
		fmt.Println("Error: input must be a number")
		return nil
	}

	for _, shop := range ShopsArr {
		if shop.Id == id {
			shopSelect = shop
			break
		}
	}

	if shopSelect == nil {
		fmt.Println("Error: shop not found")
		return nil
	}

	return NewShop(shopSelect.Id, shopSelect.Name, shopSelect.Location)
}

// InitShops creates start object to fill shops
func InitShops() []*Shop {
	ShopsArr = append(ShopsArr,
		NewShop(1, "Centro de Servicio Automotriz 'Autotec'", "Santiago, Chile"),
		NewShop(2, "Taller Mecánico 'Mecánica Rápida'", "Valparaíso, Chile"),
		NewShop(3, "Garaje 'Mantenimiento Total'", "Concepción, Chile"),
		NewShop(4, "Centro de Mantenimiento y Reparación 'AutoCare'", "Temuco, Chile"),
		NewShop(5, "Tienda de Servicio Rápido 'Soluciones Automotrices'", "Antofagasta, Chile"),
		NewShop(6, "Mecánica 'El Motorista'", "Viña del Mar, Chile"),
		NewShop(7, "Taller de Servicio Técnico 'La Excelencia'", "La Serena, Chile"),
	)
	return ShopsArr
}

// ListShops shows available shops of the company
func ListShops() {
	utilsImpl.ClearConsole()

	w := utilsImpl.CreateTabs()

	fmt.Println("Tiendas operativas: ")
	fmt.Fprintln(w, "\nID\tNombre\tDirección")
	for _, shop := range ShopsArr {
		fmt.Fprintf(w, "%d\t%s\t%s\n", shop.Id, shop.Name, shop.Location)
	}

	w.Flush()
}

// CreateShop creates a new shop for the company
func CreateShop() {
	utilsImpl.ClearConsole()

	var name, location string

	fmt.Print("Introduzca el nombre de la tienda: ")
	fmt.Scanln(&name)

	fmt.Print("Introduzca la ubicación de la tienda: ")
	fmt.Scanln(&location)

	ShopsArr = append(ShopsArr, NewShop(len(ShopsArr)+1, name, location))
}

// DeleteShop performs a hard delete for a shop of the company
func DeleteShop() {
	var idInput string

	ListShops()

	fmt.Print("\nIntroduzca el ID de la tienda a eliminar: ")
	fmt.Scanln(&idInput)

	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Error: input must be a number")
		return
	}

	for i, shop := range ShopsArr {
		if shop.Id == id {
			ShopsArr = append(ShopsArr[:i], ShopsArr[i+1:]...)
			fmt.Printf("\nTienda con ID %d eliminada!\n", id)
			return
		}
	}

	fmt.Printf("\nTienda con ID %d no encontrada\n", id)
	utilsImpl.PausedConsole()
	utilsImpl.ClearConsole()
}

// ShopsMenu creates menu of shop options
func ShopsOptions() {
	utilsImpl.ClearConsole()

	scanner := bufio.NewScanner(os.Stdin)

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
			ListShops()
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		case "2":
			CreateShop()
			ListShops()
			fmt.Println("\nTienda agregada!")
			utilsImpl.PausedConsole()
			utilsImpl.ClearConsole()
		case "3":
			DeleteShop()
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
