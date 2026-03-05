package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int

	invalid_utf := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, " charcounts:\n%v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid_utf++

		}

		counts[r]++
		utflen[n]++

		fmt.Printf("rune\tcount\n")

		for c, r := range counts {
			fmt.Printf("%v\t%d\n", c, r)

		}

		// fmt.Print("\nlen\tcount\n")
		// for i, n := range utflen {
		// 	if i > 0 {
		// 		fmt.Printf("%d\t%d\n", i, n)
		// 	}
		// }
		// if invalid_utf > 0 {
		// 	fmt.Printf("%d\t invalid unicode characters\n", invalid_utf)
		// }

	}

}
