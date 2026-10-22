
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	angka := []int{1, 2, 3, 2, 4, 1, 3, 2, 5, 1, 2}

	frekuensi := make(map[int]int)

	for _, nilai := range angka {
		frekuensi[nilai]++
	}

	fmt.Println("Data angka:", angka)
	fmt.Println("Jumlah kemunculan setiap angka:")

	for nilai, jumlah := range frekuensi {
		fmt.Println(nilai, "muncul", jumlah, "kali")
	}

	nama := "Aulia"
	bahasa := "Go"
	durasi := 1

	fmt.Println("Nama saya", nama)
	fmt.Println("Saya sedang belajar bahasa pemrograman", bahasa)
	fmt.Println("Saya berlatih presentasi selama", durasi, "menit.")
	fmt.Println("Saat presentasi, saya akan menjaga kontak mata")
	fmt.Println("dan berbicara dengan tempo yang jelas.")
}
