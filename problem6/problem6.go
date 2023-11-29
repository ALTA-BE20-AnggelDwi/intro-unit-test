package main

import (
	"fmt"
)

func HargaAkhir(hargaAwal, diskon float64) float64 {
	hargaDiskon := (diskon / 100) * hargaAwal
	hargaAkhir := hargaAwal - hargaDiskon
	return hargaAkhir
}

func main() {

	fmt.Println("Harga baju setelah diskon: Rp.", HargaAkhir(370000, 15))
}
