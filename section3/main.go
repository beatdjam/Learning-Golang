package main

import "fmt"

func outer() {
	s4 := "outer"
	fmt.Println(s4)
}

func main() {
	// 型を指定した宣言
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	// 同じ型は並べて宣言出来る
	var t, f bool = true, false
	fmt.Println(t, f)

	// 異なる型は括弧でくくって宣言出来る(タプル？)
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 宣言した時点で初期値が入る
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 値はミュータブルなので後から変更できる
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 型推論
	// 一度定義した変数に := で再代入することはできない
	i4 := 400
	fmt.Println(i4)

	// 同様に変更できる
	i4 = 450
	fmt.Println(i4)

	// 一度定義した変数は異なる型で再代入することはできない
	// 型推論は関数外で定義することはできない

	outer()

	// 未使用変数が定義された場合エラーになる
}
