
package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	prima := angka >= 2

	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan bilangan prima")
	}

	nilai := [5]int{80, 90, 75, 85, 70}

	total := 0

	for _, angka := range nilai {
		total += angka
	}

	rataRata := float64(total) / float64(len(nilai))

	fmt.Println("\nArray nilai:", nilai)
	fmt.Println("Total nilai:", total)
	fmt.Println("Rata-rata:", rataRata)

	angkaArray := [5]int{12, 45, 7, 89, 34}

	terbesar := angkaArray[0]

	for _, angka := range angkaArray {
		if angka > terbesar {
			terbesar = angka
		}
	}

	fmt.Println("\nArray angka:", angkaArray)
	fmt.Println("Nilai terbesar:", terbesar)
}
