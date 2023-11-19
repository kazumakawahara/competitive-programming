//go:build run
// +build run

package main

import (
	"fmt"
	"strings"
)

/*
問題文
英小文字からなる長さNの文字列Sが与えられます。
Sの中でaとbが隣接する箇所があれば Yes を、なければ No を出力してください。(a と b の順序は問いません。)

制約
2≤N≤100
Sは英小文字からなる長さNの文字列

入力
入力は以下の形式で標準入力から与えられる。
N,S

出力
Sの中でa と bが隣接する箇所があれば Yes を、なければ No を出力せよ。

入力例
3
abc
*/

func main() {
	var N int
	var S string

	fmt.Scan(&N, &S)

	if strings.Contains(S, "ab") || strings.Contains(S, "ba") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
