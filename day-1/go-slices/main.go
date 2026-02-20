package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, substr string
	s = "abababcd"
	substr = "ab"
	arr := []string{"1", "2", "3", "4", "5"}
	for idx, val := range os.Args[0:] {

		fmt.Println(idx, val)
	}

	fmt.Println(strings.Join(arr, " "))
	fmt.Println("this is the OS: ", os.Args[0])

	fmt.Println(strings.Count(s, substr))

}
