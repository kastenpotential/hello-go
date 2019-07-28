package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go for github ci test v0.1.0")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}
}
