//go:build run
// +build run

package main

import (
	"fmt"
)

/*
問題文
N問の問題が出題されるプログラミングコンテストがあります。
i=1,2,…,N について、i 問目の配点はSiです。

配点がX 以下である問題すべての配点の合計を出力してください。

制約
入力される値は全て整数
・4 ≤ N ≤ 8
・100 ≤ Si ≤ 675
・100 ≤ X ≤ 675

入力&
入力は以下の形式で標準入力から与えられる。
N X
S1 S2… SN


出力
答えを出力せよ。

入力例
6 200
100 675 201 200 199 328
*/

func main() {
	var N, X int

	fmt.Scan(&N, &X)

	total := 0

	for i := 0; i < N; i++ {
		{
			var S int
			fmt.Scan(&S)

			if X < S {
				continue
			}

			total += S
		}
	}

	fmt.Println(total)
}
