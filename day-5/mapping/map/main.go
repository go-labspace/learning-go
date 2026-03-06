package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int

	invalid_utf := 0
	// mapping()
	wordFreq()

	fmt.Printf("rune\tcount\n")

	for r, c := range counts {
		fmt.Printf("%v\t%d\n", r, c)

	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid_utf > 0 {
		fmt.Printf("%d\t invalid unicode characters\n", invalid_utf)
	}

}
func mapping() {
	var letterCount, digitCount, punctuationCount int
	var utf_count [utf8.UTFMax + 1]int
	invalid_xter := 0

	i := bufio.NewReader(os.Stdin)

	for {
		r, s, err := i.ReadRune()

		if err == io.EOF {
			break
		}

		if r == unicode.ReplacementChar && s == 1 {
			invalid_xter++
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if unicode.IsLetter(r) {
			letterCount++
		}
		if unicode.IsDigit(r) {
			digitCount++

		}
		if !(unicode.IsDigit(r) || unicode.IsLetter(r)) {
			punctuationCount++

		}
		utf_count[s]++
	}

	fmt.Println("Letters\tDigits\tOther")
	fmt.Printf("%d\t%d\t%d\n", letterCount, digitCount, punctuationCount)

}

func wordFreq() {
	i := bufio.NewReader(os.Stdin)

	wordCount := make(map[string]int)
	for {

		input, err := i.ReadString('\n')
		if err == io.EOF {
			break
		}
		input = strings.TrimSpace(input)
		freq := strings.Split(input, " ")

		for _, word := range freq {
			wordCount[word]++

		}

	}

	fmt.Println("word\tcount")
	for w, c := range wordCount {
		fmt.Printf("%v\t%d\n", w, c)
	}

}
