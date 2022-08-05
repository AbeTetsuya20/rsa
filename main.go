package main

import (
	"fmt"
	"rsa/application"
	"time"
)

func main() {

	// 素数を生成する時間を計測する
	start := time.Now()
	prime, err := application.MakePrime(2048)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
	}
	end := time.Now()
	fmt.Println("素数を生成しました:", prime)
	fmt.Println("生成にかかった時間:", end.Sub(start))

}

// 100:534.959µs
// 500:6.955625ms
// 1024:57.996625ms
// 1500:121.347625ms
// 2048:642.911959ms
