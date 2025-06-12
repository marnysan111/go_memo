package main

import (
	"fmt"
	"github.com/marnysan111/go_memo/internal"
)

func main() {
	pointerStruct := internal.NewPointerStruct(1, "mani", "Tokyo")
	fmt.Println(pointerStruct)
	pointerStruct = internal.NewPointerStruct(2, "marny", "")
	fmt.Println(pointerStruct)

	ps := internal.PointerStruct{
		ID:   3,
		Name: "marny",
	}
	pointerStruct.Output(&ps)

}
