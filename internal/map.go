package internal

import (
	"fmt"
)

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

type Users struct {
	Name  string
	Count int
}

// 投票システムの作成
// 名前をキーとし、得票数を値とするmapを用いて、簡単な投票システムを作成します。ユーザーから名前を入力させ、
// その名前の得票数を1増やします。特定のキー（例："exit"）が入力されたら、全員の名前と得票数を表示して終了します。
func MapTest2() {
	var str string
	users := make(map[string]*Users)
	for i := 0; i < 5; i++ {
		fmt.Scan(&str)
		if _, ok := users[str]; !ok {
			users[str] = &Users{Name: str, Count: 0}
		}
		users[str].Count++
	}
	fmt.Println(users)
	for _, user := range users {
		fmt.Println(user.Name, user.Count)
	}
}

func MapTestToSlice() {
	var str string
	var users []Users
	for i := 0; i < 5; i++ {
		fmt.Scan(&str)
		found := false
		for j := 0; j < len(users); j++ {
			if users[j].Name == str {
				users[j].Count++
				found = true
				break
			}
		}
		if !found {
			users = append(users, Users{Name: str, Count: 1})
		}
	}
	fmt.Println(users)
}
