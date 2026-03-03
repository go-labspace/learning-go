package main

import (
	"fmt"
)

func main() {
	age := map[string]int{
		"alice": 20,
		"bake":  30,
	}

	value, ok := age["bob"]

	fmt.Printf("the key %v exists: %b", value, ok)
}
