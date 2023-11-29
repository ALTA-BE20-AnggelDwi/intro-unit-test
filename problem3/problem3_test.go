package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSapaNama(t *testing.T) {
	name := "kobar!"
	expected := "Hello kobar! Saya Golang bahasa yang sangat menyenangkan."
	result := SapaNama(name)

	assert.Equal(t, expected, result, "Greeting message mismatch")
}
