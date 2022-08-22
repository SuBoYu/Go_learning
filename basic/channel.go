package main

import (
	"fmt"
	"time"
)

// multi thread channel
// 最後介紹的是使用 Channel 等待, 原為 Goroutine 溝通時使用的，
// 但因其阻塞的特性，使其可以當作等待 Goroutine 的方法

// func say(s string, c chan string) {

// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	c <- "FINISH"
// }

// func main() {
// 	ch := make(chan string)

// 	go say("world", ch)
// 	go say("hello", ch)

// 	<-ch
// 	<-ch

// }

// 多執行緒下的共享變數

// 1.sync.Mutex

// type SafeNumber struct {
// 	v   int
// 	mux sync.Mutex
// }

// func main() {
// 	total := SafeNumber{v: 0}
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			total.mux.Lock()
// 			total.v++
// 			total.mux.Unlock()
// 		}()
// 	}
// 	time.Sleep(time.Second)
// 	total.mux.Lock()
// 	fmt.Println(total.v)
// 	total.mux.Unlock()
// }

// func main() {
// 	total := 0
// 	ch := make(chan int, 1)
// 	ch <- total
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			ch <- <-ch + 1
// 		}()
// 	}
// 	time.Sleep(time.Second)
// 	fmt.Println(<-ch)
// }

// unbuffered channel v.s. buffered channel
// unbuffered channel
// 只有一條 Goroutine：main
// 推入 1 後因為還沒有其他 Goroutine 拉取 Channel 中的資料，所以進入阻塞狀態
// 因為 main 已經在推入資料時阻塞，所以拉取的程式永遠不會被執行，造成死結
// func main() {
// 	ch := make(chan int)
// 	ch <- 1
// 	fmt.Println(<-ch)
// }

// buffered channel
// 推入 1 後 Channel 內的資料數為1並沒有超過 Buffer 的長度1，所以不會被阻塞
// 因為沒有阻塞，所以下一行拉取的程式碼可以被執行，並且完成執行
// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 1
// 	fmt.Println(<-ch)
// }

// loop 中的channel
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		close(c) // close channel
// 	}()
// 	for {
// 		v, ok := <-c
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(v)
// 	}
// }

// 使用 select 避免等待
// 在 Channel 推入/拉取時，會有一段等待的時間而造成 Goroutine 無法回應，
// 如果此 Goroutine 是負責處理畫面的，使用者就會看到畫面 lag 的情況，
// 這是我們不想見的情況。

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("calculate goroutine starts calculating")
		time.Sleep(time.Second) // Heavy calculation
		fmt.Println("calculate goroutine ends calculating")

		ch <- "FINISH"
		time.Sleep(time.Second)
		fmt.Println("calculate goroutine finished")
	}()

	for {
		select {
		case <-ch: //Channel 中有資料執行此區域
			fmt.Println("main goroutine finished")
			return
		default:
			fmt.Println("WAITING...")
			time.Sleep(500 * time.Millisecond)

		}
	}
}
