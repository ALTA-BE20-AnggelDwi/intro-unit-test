package main

import "fmt"

func LuasSegitiga(alas, tinggi int) int {
	return alas * tinggi
}

func main() {

	fmt.Println(LuasSegitiga(2, 5))
}
