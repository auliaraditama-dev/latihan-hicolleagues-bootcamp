package main

import "fmt"

const Pi = 3.14

func main() {
	//? Latihan 1 sesi 2
	nama := "Budi"
	umur := 20
	kota := "Jakarta"

	fmt.Printf("Nama : %s\n", nama)
	fmt.Printf("Umur : %d\n", umur)
	fmt.Printf("Kota : %s\n", kota)

	//? Latihan 1 sesi 2
	var (
		C = 10.0
		F = C * 9/5 + 32.0
	)

	fmt.Print("\nHasil Fahrenheit dari Cecius adalah ")
	fmt.Printf("Fahrenheit: %.1f \n", F)

	//? Latihan 3 sesi 2
	const Pi = 3.14

	var (
		r = 10.0
		luas = Pi * r * r
		keliling = 2 * Pi * r
	)

	fmt.Printf("\nLuas Lingkaran: %.2f\n", luas)
	fmt.Printf("Kelilind Lingkaran: %.2f\n", keliling)
	fmt.Println("Bilangan Genap:",42%2==0)
}
