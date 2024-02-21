package internal

import "fmt"

func MapTest() {
	sampleStr := "banana"
	countStr := make(map[rune]int)
	for _, s := range sampleStr {
		countStr[s]++
	}
	for char, count := range countStr {
		fmt.Printf("'%c': %d\n", char, count)
	}

	for char, count := range countStr {
		fmt.Println(char, ":", count)
	}
}
