package main

import (
	"fmt"
	"os"
)

func main() {

	cute_cat, err := os.ReadFile("cute-cat.webp")

	if err != nil {
		panic(err)
	}

	fmt.Printf("this cute do is just %d bytes\n", len(cute_cat))

}
