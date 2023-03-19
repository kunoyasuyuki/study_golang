package main

import (
	"fmt"
	"math"
)

func main() {

	// エクスポートされていない名前(小文字ではじまる名前)は、外部のパッケージからアクセスすることはできません。
	//fmt.Println(math.pi)

	// 下記が正解
	fmt.Println(math.Pi)
}
