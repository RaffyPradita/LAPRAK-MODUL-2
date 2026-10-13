package main

import "fmt"

func main4() {
	var nama, nim, kelas string

	fmt.Print("Masukkan nama, nim, dan kelas: ")
	fmt.Scan(&nama, &nim, &kelas)

	fmt.Println("Perkenalkan saya adalah", nama+", salah satu mahasiswa Prodi SI-1F dari kelas", kelas, "dengan NIM", nim+".")
}
