package main

import "fmt"

func main1() {
	var a, b, c, d, e int
	var hasil int
	fmt.Print("Masukkan 5 angka yang ingin dijumlahkan:")
	fmt.Scanln(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.Println("Hasil penjumlahan", a, b, c, d, e, "adalah", hasil)
}
