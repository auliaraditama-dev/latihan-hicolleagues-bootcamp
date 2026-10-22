package main

import (
	"errors"
	"fmt"
)

func validasiInput(a, b float64) (bool, error) {
	if a < 0 || b < 0 {
		return false, errors.New("input tidak boleh bernilai negatif")
	}
	return true, nil
}

func tambah(a, b float64) float64 {
	return a + b
}

func kurang(a, b float64) float64 {
	return a - b
}

func kali(a, b float64) float64 {
	return a * b
}

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak dapat membagi dengan nol")
	}
	return a / b, nil
}

func main() {
	var num1, num2 float64 = 15.5, 5.0

	fmt.Printf("Mencoba operasi kalkulator untuk angka %.2f dan %.2f\n", num1, num2)
	fmt.Println("-----------------------------------------------------")

	valid, errValidasi := validasiInput(num1, num2)
	if errValidasi != nil {
		fmt.Println("Error Validasi:", errValidasi)
		return 
	}

	if valid {
		fmt.Println("Status Validasi: Input Valid")

		hasilTambah := tambah(num1, num2)
		fmt.Printf("Hasil Tambah : %.2f\n", hasilTambah)

		hasilKurang := kurang(num1, num2)
		fmt.Printf("Hasil Kurang : %.2f\n", hasilKurang)

		hasilKali := kali(num1, num2)
		fmt.Printf("Hasil Kali   : %.2f\n", hasilKali)

		hasilBagi, errBagi := bagi(num1, num2)
		if errBagi != nil {
			fmt.Println("Error Bagi   :", errBagi)
		} else {
			fmt.Printf("Hasil Bagi   : %.2f\n", hasilBagi)
		}
	}

	fmt.Println("\n--- Uji coba error handling ---")
	hasilBagiNol, errBagiNol := bagi(10, 0)
	if errBagiNol != nil {
		fmt.Println("Uji Bagi Nol - Error:", errBagiNol)
	} else {
		fmt.Printf("Uji Bagi Nol - Hasil: %.2f\n", hasilBagiNol)
	}

	_, errValidasiGagal := validasiInput(-5, 10)
	if errValidasiGagal != nil {
		fmt.Println("Uji Validasi - Error:", errValidasiGagal)
	}
}