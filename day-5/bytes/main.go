package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("bytes/cute-cat.webp")
	cat := bufio.NewReader(f)

	r, n, _ := cat.ReadRune()

	fmt.Printf("rune:%v\nbytes:%d", r, n)

}
