package main

import "fmt"

func main() {

	//? Latihan 1 sesi 1
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	//? Latihan 2 sesi 1
	var bilangan1, bilangan2 float64
	var operator string
	var hasil float64
	// var hasilPasti = int(hasil) //* Konversi tipe data float64 ke int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&bilangan1)
	fmt.Print("Masukkan Operator: ")
	fmt.Scan(&operator)
	fmt.Print("Masukkan Bilangan kedua: ")
	fmt.Scan(&bilangan2)

	
	switch operator {
	case "+":
		hasil = bilangan1 + bilangan2
	case "-":
		hasil = bilangan1 - bilangan2
	case "*":
		hasil = bilangan1 * bilangan2
	case "/":
		if bilangan2 == 0 {
			fmt.Print("Tidak Bisa Dibagi Dengan Angka Nol!!!")
			return
		} else {
			hasil = bilangan1 / bilangan2
		}
	default:
		fmt.Print("Operator tidak di dukung")
		return
	}
	fmt.Println("Hasilnya adalah: ", hasil)
	// fmt.Printf("Hasilnya adalah: %d", hasilPasti) //* Digunakan saat ada var data yang di konversi

	//? Latihan 3 sesi 1
	var n int
	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++{
		println("hasil perulangan: ",i)
	}

	total := 0

	for i := 1; i <= n; i++{
		total += i
	}
	fmt.Println("hasil penjumlahan:",total)



}
