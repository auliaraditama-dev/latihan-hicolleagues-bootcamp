package belanja

import (
	"errors"
	"fmt"
	"strings"
)

type Item struct {
	ID    int
	Nama  string
	Harga float64
	Qty   int
}

var items []Item
var nextID = 1

func Tambah(nama string, harga float64, qty int) (Item, error) {
	if strings.TrimSpace(nama) == "" {
		return Item{}, errors.New("nama tidak boleh kosong")
	}
	if harga < 0 || qty <= 0 {
		return Item{}, errors.New("harga atau qty tidak valid")
	}
	it := Item{ID: nextID, Nama: nama, Harga: harga, Qty: qty}
	nextID++
	items = append(items, it)
	return it, nil
}

func Semua() []Item {
	return items
}

func Hapus(id int) error {
	for i, it := range items {
		if it.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item %d tidak ditemukan", id)
}

func Ubah(id, qty int) error {
	if qty <= 0 {
		return errors.New("qty tidak valid")
	}
	for i := range items {
		if items[i].ID == id {
			items[i].Qty = qty
			return nil
		}
	}
	return fmt.Errorf("item %d tidak ditemukan", id)
}

func Total() float64 {
	total := 0.0
	for _, it := range items {
		total += it.Harga * float64(it.Qty)
	}
	return total
}

func Cari(nama string) []Item {
	var hasil []Item
	cariLower := strings.ToLower(strings.TrimSpace(nama))
	for _, it := range items {
		if strings.Contains(strings.ToLower(it.Nama), cariLower) {
			hasil = append(hasil, it)
		}
	}
	return hasil
}