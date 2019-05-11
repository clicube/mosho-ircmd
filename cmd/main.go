package main

import (
	"fmt"
	"os"

	"mosho-cmdsub/internal"
)

func main() {
	err := internal.Sub()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
	os.Exit(0)
}
