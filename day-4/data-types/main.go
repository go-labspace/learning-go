package main

import (
	"fmt"
	"strconv"
)

var x int = 123

func main() {
	fmt.Println(strconv.FormatInt(int64(x), 16)) // "1111011"
}
