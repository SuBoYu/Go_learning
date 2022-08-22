package main

import (
	"fmt"
)

// type SyncedBuffer struct {
// 	lock   sync.Mutex
// 	buffer bytes.Buffer
// 	foo    int
// 	bar    string
// }

// new(T) 宣告會直接拿到儲存位置，並且配置 Zero Value (初始化)，也就是數字型態為 0，字串型態就是 ""

// func main() {
// 	p := new(SyncedBuffer)
// 	fmt.Println("foo:", p.foo)
// 	fmt.Println("bar:", p.bar)
// 	fmt.Printf("%#v\n", p)
// }

// 通常在宣告 slice, map 及 channel 則會使用 Go 提供的另一個宣告方式 make

func main() {
	var p *map[string]string
	// new
	p = new(map[string]string)
	*p = map[string]string{
		"bar": "foo",
	}
	people := *p
	people["foo"] = "bar"

	fmt.Println(people)
	fmt.Println(p)

	//make
	foobar := make(map[string]string)
	foobar["foo"] = "bar"
	foobar["bar"] = "foo"
	fmt.Println(foobar)

}

// 總結底下 make 跟 new 的區別

// make 能夠分配並且初始化所需要的記憶體空間跟結構，而 new 只能回傳指標位置
// make 只能用在三種類型 slice, map 及 channel
// make 可以初始化上述三種格式的長度跟容量以便提供效率跟減少開銷