package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuasSegitiga(t *testing.T) {
	alas := 2
	tinggi := 5
	expected := 10
	result := LuasSegitiga(alas, tinggi)

	assert.Equal(t, expected, result, "Luas salah")
}
