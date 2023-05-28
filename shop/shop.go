package shop

import (
	"fmt"
	"proyecto_final_goland/utils"
)

type shop struct {
	id       int
	name     string
	location string
}

func newShop(id int, name string, location string) *shop {
	return &shop{
		id:       id,
		name:     name,
		location: location,
	}
}

func CreateShop() {
	utils.ClearConsole()
	// Crea instancias de Service
	shops := []*shop{
		newShop(1, "Centro de Servicio Automotriz 'Autotec'", "Santiago, Chile"),
		newShop(2, "Taller Mecánico 'Mecánica Rápida'", "Valparaíso, Chile"),
		newShop(3, "Garaje 'Mantenimiento Total'", "Concepción, Chile"),
		newShop(4, "Centro de Mantenimiento y Reparación 'AutoCare'", "Temuco, Chile"),
		newShop(5, "Tienda de Servicio Rápido 'Soluciones Automotrices'", "Antofagasta, Chile"),
		newShop(6, "Mecánica 'El Motorista'", "Viña del Mar, Chile"),
		newShop(7, "Taller de Servicio Técnico 'La Excelencia'", "La Serena, Chile"),
	}

	// Crea tabs
	w := utils.CreateTabs()

	// Mostrar tiendas operativas
	fmt.Println("Tiendas operativas: ")
	fmt.Fprintln(w, "\nID\tNombre\tDirección")
	for _, shp := range shops {
		fmt.Fprintf(w, "%d\t%s\t%s\n", shp.id, shp.name, shp.location)
	}

	w.Flush()
	utils.PausedConsole()
	utils.ClearConsole()
}
