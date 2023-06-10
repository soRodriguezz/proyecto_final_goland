package utils

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
)

/*
 * Util que pausa consola hasta realizar un enter
 */
func PausedConsole() {
	fmt.Println("\nPresione Enter para continuar...")
	fmt.Scanln()
}

/*
 * Util para borrar consola
 */
func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
 * Util que configura tabla con tabulaciones
 */
func CreateTabs() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	return w
}
