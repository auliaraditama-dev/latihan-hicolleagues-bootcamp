package main

import "fmt"

func main() {
	var nama string = "Budi"
	var umur int = 25
	var tinggi float64 = 170.5
	var aktif bool = true

	fmt.Printf("Nama: %s Umur: %d Tinggi: %.1f Aktif: %t\n", nama, umur, tinggi, aktif)

	fmt.Println("\nHasil Kelulusan Mahasiswa:")
	
	nilai := []int{80, 65, 90, 55, 75}

	for _, n := range nilai {
		if n >= 70 {
			fmt.Printf("Nilai %d -> Lulus\n", n)
		} else {
			fmt.Printf("Nilai %d -> Tidak Lulus\n", n)
		}
	}
}