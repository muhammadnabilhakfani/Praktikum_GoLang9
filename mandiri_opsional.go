package main

import "fmt"

func main() {
	var aksi string
	fmt.Scan(&aksi)

	var konsekuensi1, konsekuensi2 string

	if aksi == "mobil" {
		konsekuensi1 = "telat masuk kuliah"
		konsekuensi2 = "mendapat pengurangan nilai"
	}

	if aksi == "motor" {
		konsekuensi1 = "bajunya basah"
		konsekuensi2 = "harus membawa baju ganti"
	}

	fmt.Println("Aksi:", aksi)
	fmt.Println("Konsekuensi 1:", konsekuensi1)
	fmt.Println("Konsekuensi 2:", konsekuensi2)
}
