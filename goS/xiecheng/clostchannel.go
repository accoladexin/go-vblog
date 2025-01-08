package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5) // 缓冲，未关闭通道
	fmt.Printf("c1: %d, %d, %v\n", len(c1), cap(c1), c1)
	c1 <- 111
	c1 <- 222
	c1 <- 333
	fmt.Println(<-c1, "###") // 故意读走一个
	close(c1)                // 关闭通道，不许再进数据
	for v := range c1 {
		fmt.Println(v, "~~~") // 打印元素
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~") // 打印了这一句
}
