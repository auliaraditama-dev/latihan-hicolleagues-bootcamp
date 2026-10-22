
package main

import "fmt"

func main() {

	var nilai int

	fmt.Print("Masukkan nilai (0-100): ")
	fmt.Scan(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai tidak valid")
	} else if nilai >= 90 {
		fmt.Println("Grade: A")
	} else if nilai >= 80 {
		fmt.Println("Grade: B")
	} else if nilai >= 70 {
		fmt.Println("Grade: C")
	} else if nilai >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}

	nilaiSiswa := []int{80, 90, 75, 85, 70}

	total := 0
	nilaiTertinggi := nilaiSiswa[0]
	nilaiTerendah := nilaiSiswa[0]

	for _, nilai := range nilaiSiswa {
		total += nilai

		if nilai > nilaiTertinggi {
			nilaiTertinggi = nilai
		}

		if nilai < nilaiTerendah {
			nilaiTerendah = nilai
		}
	}

	rataRata := float64(total) / float64(len(nilaiSiswa))

	fmt.Println("Daftar nilai:", nilaiSiswa)
	fmt.Println("Jumlah siswa:", len(nilaiSiswa))
	fmt.Println("Total nilai:", total)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
	fmt.Println("Nilai tertinggi:", nilaiTertinggi)
	fmt.Println("Nilai terendah:", nilaiTerendah)

	type Siswa struct {
		nama  string
		umur  int
		nilai int
	}

	daftarSiswa := []Siswa{
		{nama: "Aulia", umur: 17, nilai: 90},
		{nama: "Budi", umur: 16, nilai: 85},
		{nama: "Citra", umur: 17, nilai: 78},
		{nama: "Dina", umur: 16, nilai: 95},
	}

	for nomor, siswa := range daftarSiswa {
		fmt.Printf("\nSiswa ke-%d\n", nomor+1)
		fmt.Println("Nama:", siswa.nama)
		fmt.Println("Umur:", siswa.umur)
		fmt.Println("Nilai:", siswa.nilai)

		if siswa.nilai >= 75 {
			fmt.Println("Status: Lulus")
		} else {
			fmt.Println("Status: Tidak lulus")
		}
	}

	fmt.Println("\nJumlah seluruh siswa:", len(daftarSiswa))
}
