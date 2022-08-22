package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// one thread
// func main() {
//     say("world")
//     say("hello")
// }

// multi thread waitgroup
// 多執行緒下，經常需要處理的是執行緒之間的狀態管理，其中一個經常發生的事情是等待，
// 例如A執行緒需要等B執行緒計算並取得資料後才能繼續往下執行，在這情況下等待就變得十分重要。
func main() {
	//核心數
	fmt.Println(runtime.NumCPU())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go say("world", wg)
	go say("hello", wg)

	wg.Wait()
}

// multi thread channel
// 最後介紹的是使用 Channel 等待, 原為 Goroutine 溝通時使用的，
// 但因其阻塞的特性，使其可以當作等待 Goroutine 的方法
