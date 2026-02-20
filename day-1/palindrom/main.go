package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Accept a sring via input
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	r := len(input.Text()) - 1
	palindrome_check := 1
	for l := 0; l < r; l++ {
		if input.Text()[l] != input.Text()[r-l] {
			palindrome_check = 0
			r--
			break
		}

	}

	if palindrome_check == 1 {
		fmt.Printf("%v is palindrome", input.Text())
	}

	if palindrome_check == 0 {
		fmt.Printf("%v is not palindrome\n", input.Text())
	}

}
