package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/wcharczuk/slant/slant"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
	}
	slant.Print(os.Stdout, strings.Join(os.Args[1:], " "))
}
