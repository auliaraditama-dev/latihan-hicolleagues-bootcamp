package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"latihan-hicolleagues-bootcamp/sesi6/hand-on/belanja" 
)

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	s, _ := in.ReadString('\n')
	return strings.TrimSpace(s)
}

func cetakTabel(items []belanja.Item) {
	if len(items) == 0 {
		fmt.Println("Data kosong.")
		return
	}
	fmt.Printf("\n%-5s | %-20s | %-12s | %-5s\n", "ID", "Nama", "Harga", "Qty")
	fmt.Println(strings.Repeat("-", 53))
	for _, it := range items {
		fmt.Printf("%-5d | %-20s | %-12.2f | %-5d\n", it.ID, it.Nama, it.Harga, it.Qty)
	}
	fmt.Println()
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMenu: 1 Tambah | 2 Lihat | 3 Ubah Qty | 4 Hapus | 5 Total | 6 Cari | 0 Keluar")
		pilih := baca(in, "Pilih: ")

		switch pilih {
		case "1":
			nama := baca(in, "Nama: ")
			harga, err := strconv.ParseFloat(baca(in, "Harga: "), 64)
			if err != nil {
				fmt.Println("Error: Harga harus angka")
				continue
			}
			qty, err := strconv.Atoi(baca(in, "Qty: "))
			if err != nil {
				fmt.Println("Error: Qty harus angka")
				continue
			}
			_, err = belanja.Tambah(nama, harga, qty)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item berhasil ditambahkan.")
			}

		case "2":
			cetakTabel(belanja.Semua())

		case "3":
			id, err := strconv.Atoi(baca(in, "ID Item: "))
			if err != nil {
				fmt.Println("Error: ID harus angka")
				continue
			}
			qty, err := strconv.Atoi(baca(in, "Qty Baru: "))
			if err != nil {
				fmt.Println("Error: Qty harus angka")
				continue
			}
			if err := belanja.Ubah(id, qty); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Qty berhasil diubah.")
			}

		case "4":
			id, err := strconv.Atoi(baca(in, "ID Item yang dihapus: "))
			if err != nil {
				fmt.Println("Error: ID harus angka")
				continue
			}
			if err := belanja.Hapus(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item berhasil dihapus.")
			}

		case "5":
			fmt.Printf("Total Harga: %.2f\n", belanja.Total())

		case "6":
			nama := baca(in, "Masukkan nama yang dicari: ")
			hasil := belanja.Cari(nama)
			cetakTabel(hasil)

		case "0":
			fmt.Println("Keluar dari aplikasi.")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}