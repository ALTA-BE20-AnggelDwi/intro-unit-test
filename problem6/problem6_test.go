package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHargaAkhir(t *testing.T) {
	var hargaAwal float64 = 370000
	var diskon float64 = 15
	var expected float64 = 314500
	result := HargaAkhir(hargaAwal, diskon)

	assert.Equal(t, expected, result, "Harga salah")
}

//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out -o coverage.html
//start coverage.html
