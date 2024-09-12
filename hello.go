package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtener la fecha y hora actual
	currentTime := time.Now()

	// Mostrar la fecha y hora en formato estándar
	fmt.Println("Fecha y hora actual:", currentTime.Format("2006-01-02 15:04:05"))
}
