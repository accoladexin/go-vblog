package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	age  int
}

type Person1 struct {
	Name string `json:"name"`
	age  int
}

func main() {
	var v1 chan int
	v2 := make(chan int, 10)
	v3 := make(chan int, 10)
	fmt.Printf("v1==> %T %d %d %v\n", v1, len(v1), cap(v1), v1)
	fmt.Printf("v2==> %T %d %d %v\n", v2, len(v2), cap(v2), v2)
	fmt.Printf("v3==> %T %d %d %v\n", v3, len(v3), cap(v3), v3)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		v3 <- i
	}
	for i := 0; i < 10; i++ {
		temp := <-v3
		fmt.Println(temp)
	}
}
