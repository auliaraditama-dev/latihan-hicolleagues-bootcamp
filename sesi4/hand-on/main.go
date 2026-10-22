package main

import (
	"errors"
	"fmt"
	"math"
)

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak dapat dibagi dengan nol")
	}
	return a / b, nil
}

func main() {
	min, max := minMax(25, 10)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	hasilBagi, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil bagi: %.2f\n", hasilBagi)
	}

	hasilBagiNol, errNol := bagi(5, 0)
	if errNol != nil {
		fmt.Println("Error:", errNol)
	} else {
		fmt.Printf("Hasil bagi: %.2f\n", hasilBagiNol)
	}

	pangkat := math.Pow(3, 2)
	akar := math.Sqrt(16)
	fmt.Printf("3 pangkat 2: %.0f\n", pangkat)
	fmt.Printf("Akar 16: %.0f\n", akar)
}