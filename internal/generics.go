package internal

import (
	"fmt"
	"strconv"
)

/*
課題1: 汎用的なスタックの実装
スタックは後入れ先出し（LIFO）のデータ構造で、push、pop、peek などの基本的な操作を実行します。ジェネリクスを使って、任意の型に対応するスタックを実装してください。

要件:

Push メソッドで要素を追加。
Pop メソッドで最後の要素を取り出し、その要素を返す。
Peek メソッドで最後の要素を見るだけ。
空の状態で Pop や Peek を呼び出すときは、適切に処理する。
URL:https://chatgpt.com/c/567c414b-2a56-45fd-817e-96320c7d4920
*/

type Stringer interface {
	String() string
}

// 構造体に[T any]とつける行為がジェネリクスを使った構造体の定義をしている
type Stack[T any] struct {
	items []T
}

// ...Tは可変長引数の意味。stack.Push(MyInt(1), MyInt(2), MyInt(3), MyString("hoge"))みたいなことが可能になり、引数なしもいける
func (s *Stack[T]) Push(items ...T) {
	s.items = append(s.items, items...)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // T型のゼロ値を返す
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

type MyInt int

// .String()はGoのfmtパッケージがString()メソッドを持つ値を自動で検出して文字列を取得するため自動実行される
func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

// 一見意味のないように見えるstring->stringだけど、独自メソッドや統一性から作った方がいいらしい。
type MyString string

func (s MyString) String() string {
	return string(s)
}

func GenericsTest() {
	fmt.Println("Generics Test")
	stack := &Stack[Stringer]{}
	stack.Push(MyInt(1), MyInt(2), MyInt(3), MyString("hoge"))

	item, ok := stack.Pop()
	if ok {
		fmt.Println(item.String())
	}

	item, ok = stack.Peek()
	if ok {
		fmt.Println(item.String())
	}
	stack.Push()
	item, ok = stack.Peek()
	if ok {
		fmt.Println(item.String())
	}
}

// func GenericsTest() {
// 	fmt.Println("generics test")
// 	fmt.Println(lifoPush([]MyInt{1, 2, 3}))
// 	fmt.Println(lifoPush([]MyString{"hoge"}))
// 	fmt.Println(lifoPop())
// 	fmt.Println(lifoPeek())
// }

// var result []string

// func lifoPush[T Stringer](xs []T) []string {
// 	for _, x := range xs {
// 		result = append(result, x.String())
// 	}
// 	return result
// }

// func lifoPop() string {
// 	pop := result[len(result)-1]
// 	result = result[:len(result)-1]
// 	return pop
// }

// func lifoPeek() string {
// 	peek := result[len(result)-1]
// 	return peek
// }

// type Stringer interface {
// 	String() string
// }

// type MyInt int

// func (i MyInt) String() string {
// 	return strconv.Itoa(int(i))
// }

// type MyString string

// func (s MyString) String() string {
// 	return string(s)
// }
