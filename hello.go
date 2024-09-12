package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtain actual date and time, and declare my CWID
	currentTime := time.Now()
	CWID := "99715"
	// Print Hello World!
	fmt.Println("Hello world!")
	// Print date and time, as well as the last 4 digits of my CWID
	fmt.Printf("It's: %s now and my CWID ending in %s.\n", currentTime.Format("2006-01-02 15:04:05"), CWID)
}
