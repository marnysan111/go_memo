package internal

import (
	"fmt"
)

func createGenerator(start, step int) func() int {
	// ヒント：
	// この関数の外で、現在の数値を覚えておく変数が必要になります。
	// そして、その変数を操作する「無名関数」を返します。
	current := start - step // ちょっとしたコツ：こうすると最初の呼び出しでちょうど`start`になる

	return func() int {
		// ここに、呼び出されるたびに行う処理を書く
		// 1. current の値を step だけ増やす
		// 2. 増やした後の current の値を返す
		current += step // ここで current を更新
		return current  // ← これを書き換える
	}
}

func FuncTest() {
	// 1. 10から始まり、1ずつ増えるジェネレータを作る
	countUp := createGenerator(10, 1)
	fmt.Println("1ずつアップ:")
	fmt.Println(countUp())
	fmt.Println(countUp())
	fmt.Println(countUp())

	fmt.Println("--------------------")

	// 2. 0から始まり、5ずつ増えるジェネレータを作る
	countBy5 := createGenerator(0, 5)
	fmt.Println("5ずつアップ:")
	fmt.Println(countBy5())
	fmt.Println(countBy5())
	fmt.Println(countBy5())

	// countUpをもう一回呼んでも、countBy5には影響されないことを確認
	fmt.Println("--------------------")
	fmt.Println("もう一度1ずつアップ:")
	fmt.Println(countUp())
}
