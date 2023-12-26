package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand10Times() {
	for i := 0; i < 10; i++ {
		//rand.Seed(time.Now().UnixNano())
		randTime := rand.Intn(10)
		fmt.Println(randTime)
		time.Sleep(time.Duration(randTime) * time.Second)
	}
}
