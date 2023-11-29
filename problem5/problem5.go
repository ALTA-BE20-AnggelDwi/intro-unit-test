package main

import (
	"fmt"
)

func LuasPermukaanTabung(T, r float64) float64 {
	return 2 * 3.14 * r * (r + T)
}

func main() {

	fmt.Println(LuasPermukaanTabung(20, 4))
}
