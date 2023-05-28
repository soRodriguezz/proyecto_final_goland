package utils

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"
)

func TimeToDuration(t time.Time) time.Duration {
	now := time.Now()
	return now.Sub(t)
}

func PausedConsole() {
	fmt.Println("\nPresione Enter para continuar...")
	fmt.Scanln()
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CreateTabs() *tabwriter.Writer {
	// Configura tabla con tabulaciones
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	return w
}
