package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuasPermukaanTabung(t *testing.T) {
	var T float64 = 20
	var r float64 = 4
	expected := 602.88
	result := LuasPermukaanTabung(T, r)

	assert.Equal(t, expected, result, "Luas salah")
}
