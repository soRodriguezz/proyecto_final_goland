package utils

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
)

type Utils interface {
	PausedConsole()
	ClearConsole()
	CreateTabs() *tabwriter.Writer
}

type UtilsImpl struct{}

func NewUtils() Utils {
	return &UtilsImpl{}
}

/*
 * Pausa la consta hasta dar ENTER
 */
func (c UtilsImpl) PausedConsole() {
	fmt.Println("\nPresione Enter para continuar...")
	fmt.Scanln()
}

/*
 * Limpia la consola
 */
func (c UtilsImpl) ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
 * Crea tabla con tabs
 */
func (c UtilsImpl) CreateTabs() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	return w
}
