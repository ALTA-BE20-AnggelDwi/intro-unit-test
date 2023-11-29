package main

import "fmt"

func SapaNama(nama string) string {
	return "Hello " + nama + " Saya Golang bahasa yang sangat menyenangkan."
}

func main() {
	fmt.Println(SapaNama("kobar!"))
}
