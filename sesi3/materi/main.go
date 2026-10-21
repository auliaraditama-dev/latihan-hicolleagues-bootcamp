package main

import (
	"fmt"
)

func main() {
	nilai := 80

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	umur := 20
	punyaKTP := true

	if umur >= 17 && punyaKTP {
		fmt.Println("Boleh Daftar")
	}

	if n := len("Go"); n > 1 {
		fmt.Println("Panjang", n)
	}

	hari := 9

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	default:
		fmt.Println("Undefined")
	}

	switch hari {
	case 6, 7:
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari Kerja")
	}

	switch {
	case nilai >= 85:
		fmt.Println("A")
	case nilai >= 75:
		fmt.Println("B")
	default:
		fmt.Println("c")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	n := 3

	for n > 0 {
		fmt.Println(n)
		n--
	}

	for {
		fmt.Println("Hello World!")
		break
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	buah := []string{"apel", "jeruk"}
	buah = append(buah, "mangga", "pisang")

	fmt.Println(buah, len(buah))


}
