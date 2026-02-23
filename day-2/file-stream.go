package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)

	} else {
		for _, args := range files {
			file, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup error: %v", err)
				continue
			}

			countLines(file, counts)
			file.Close()
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Println("This line occured in: ")
			for _, file := range files {
				check_file, _ := os.Open(file)
				word := bufio.NewScanner(check_file)
				for word.Scan() {

					if strings.Contains(word.Text(), line) {
						fmt.Print(file)
					}
				}
				fmt.Print("\n")

			}

		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	file := bufio.NewScanner(f)

	for file.Scan() {
		counts[file.Text()]++
	}
}
