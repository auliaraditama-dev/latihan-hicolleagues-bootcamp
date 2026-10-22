
package main

import "fmt"

func main() {
	var berat, tinggi float64

	fmt.Print("\nMasukkan berat badan (kg): ")
	fmt.Scan(&berat)

	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scan(&tinggi)

	if berat > 0 && tinggi > 0 {
		bmi := berat / (tinggi * tinggi)
		fmt.Printf("BMI: %.2f\n", bmi)
	} else {
		fmt.Println("Berat dan tinggi harus lebih dari 0.")
	}

	var detik int

	fmt.Print("\nMasukkan jumlah detik: ")
	fmt.Scan(&detik)

	if detik >= 0 {
		jam := detik / 3600
		sisaDetik := detik % 3600
		menit := sisaDetik / 60
		detikAkhir := sisaDetik % 60

		fmt.Printf("%d jam, %d menit, %d detik\n",
			jam, menit, detikAkhir)
	} else {
		fmt.Println("Jumlah detik tidak boleh negatif.")
	}
}
