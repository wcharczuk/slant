package main

import (
	"fmt"
	"os"
	"strings"

	"go.charczuk.com/sdk/slant"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
	}
	slant.Print(os.Stdout, strings.Join(os.Args[1:], " "))
}
