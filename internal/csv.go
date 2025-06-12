package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSV(filePath string) {
	fmt.Println("Reading CSV file:", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		log.Fatal(err)
	}

	// [][]stringなのでループする
	for _, v := range rows {
		fmt.Println(v)
	}

}
