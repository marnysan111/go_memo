package internal

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

//商品のリストを受け取り、各商品の価格に消費税を加算した新しい価格リストを作りたい。

func ApplyTax[T float64 | int](prices []T, taxRate float64) []T {
	result := make([]T, len(prices))
	for i, p := range prices {
		r := float64(p) * (1 + taxRate)
		result[i] = T(r)
	}
	return result
}

func ApplyTax2[T constraints.Integer | constraints.Float](prices []T, taxRate float64) []T {
	result := make([]T, len(prices))
	for i, p := range prices {
		r := float64(p) * (1 + taxRate)
		if _, ok := any(p).(int); ok {
			result[i] = T(math.Round(r)) // 整数の場合は結果を丸める
		} else {
			result[i] = T(r) // 浮動小数点数の場合はそのままキャスト
		}
	}
	return result
}

func GenericsTest2() {
	taxRate := 0.10
	fmt.Println(ApplyTax2([]int{100, 200, 300, 404}, taxRate))
	fmt.Println(ApplyTax2([]float64{500.50, 600.60}, taxRate))
}
