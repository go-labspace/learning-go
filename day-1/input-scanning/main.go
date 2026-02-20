package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	i := 5
	for i > 0 {
		input.Scan()
		counts[input.Text()]++
		i--

	}
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("key value pair: %d and %v\n", n, line)
		}
	}
}
