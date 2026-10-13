package main

import "fmt"

func main6() {
	var fahrenheit float64
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.0f\n", celsius)
}
