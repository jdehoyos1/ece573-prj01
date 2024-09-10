package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now().UTC()
	fmt.Printf("Hello world!\n")
	fmt.Printf("It's ")
	fmt.Printf("%d-%d-%d %d:%d:%d", dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second())
	fmt.Printf(" now and my CWID ending in 9715\n")

}
