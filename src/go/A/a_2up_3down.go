//go:build run
// +build run

package main

import (
	"fmt"
)

/*
問題文
高橋君が100階建てのビルにいます。

高橋君は2階分までの上り、または、3階分までの下りであれば移動には階段を使い、そうでないときエレベーターを使います。

高橋君がX階からY階への移動に使うのは階段ですか？

入力
入力は以下の形式で標準入力から与えられる。
X Y

出力
移動に使うのが階段ならば Yes、エレベーターならば No を出力せよ。

入力例
1 4
*/

func main() {
	var X, Y int

	fmt.Scan(&X, &Y)

	stairRange := []int{-3, -2, -1, 0, 1, 2}

	movingRange := Y - X

	for _, v := range stairRange {
		if movingRange == v {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
