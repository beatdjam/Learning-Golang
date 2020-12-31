package main

import "fmt"

func main() {
	// intだけで指定した場合、intの範囲は実行環境依存となる
	var i int = 100

	fmt.Println(i + 50)
	// 明示的に指定することは出来る
	var i2 int64 = 200

	// int と int64は異なる型と認識されるので演算できない
	fmt.Printf("%T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i))

	// 型を揃えてやれば演算ができる
	fmt.Println(i + int(i))
}
