package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	file, _ := os.Open("file1.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

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
