package impl_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func worker(ctx context.Context, id int) {
	fmt.Printf("Worker %d started\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			// 由于会一直执行，所以 这里需要sleep
			time.Sleep(1 * time.Second)
		}
	}
}

func Test_main(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() //确保在函数退出时取消上下文
	go worker(ctx, 1)
	go worker(ctx, 2)

	<-ctx.Done() // 等待上下文被取消或超时

	fmt.Println("Main function exiting")
}
