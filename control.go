package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Binary: %b\\%b", 4, 5)
	}
}
