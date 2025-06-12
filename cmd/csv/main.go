package main

import (
	"fmt"

	"github.com/marnysan111/go_memo/internal"
)

func main() {
	fmt.Println("Hello, World!")
	internal.ReadCSV("./data/sample.csv")
}
