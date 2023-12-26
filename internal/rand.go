package internal

import (
	"fmt"
	"math/rand"
)

func Rand10Times() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}
