package internal

import (
	"fmt"
	"slices"
)

func Slices() {
	strSlice := []string{"a", "b", "c"}
	intSlice := []int{1, 2, 3}
	if slices.Contains(strSlice, "a") {
		fmt.Println("Found a in strSlice")
	}
	if slices.Contains(intSlice, 1) {
		fmt.Println("Found 1 in intSlice")
	}
}
