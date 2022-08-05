package main

import (
	"fmt"
	"rsa/application"
	"time"
)

func main() {

	// 素数生成する時間を3回計測して、その平均を出力する
	var total time.Duration
	for i := 0; i < 3; i++ {
		start := time.Now()
		_, err := application.MakePrime(3000)
		if err != nil {
			panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
		}
		end := time.Now()
		total += end.Sub(start)
	}
	fmt.Println("生成にかかった時間:", total/3)

}

// 100:534.959µs
// 500:6.955625ms
// 1024:57.996625ms
// 1500:121.347625ms
// 2048:642.911959ms
