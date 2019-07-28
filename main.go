package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go for github ci test 001")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}
}
