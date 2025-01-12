package common_test

import (
	"fmt"
	"testing"
)

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func describeAnimal(animal Animal) {
	fmt.Print("Animal says: ")
	animal.Speak()

	// 断言获取实际类型和值
	// 1. Comma-ok 断言
	if dog, ok := animal.(Dog); ok {
		fmt.Println("It's a dog, name:", dog.Name)
	} else if cat, ok := animal.(Cat); ok {
		fmt.Println("It's a cat, name:", cat.Name)
	} else {
		fmt.Println("Unknown animal type")
	}

	// 2. Panic 断言 (不推荐直接使用，可能导致panic)
	if cat, ok := animal.(Cat); ok {
		fmt.Println("It's a cat using panic assertion, name:", cat.Name)
	}
}

func Test_tmain(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	describeAnimal(dog)
	describeAnimal(cat)

	var animal Animal = dog
	dog.Name = "newDog"
	fmt.Println(animal)
	fmt.Println(dog)

	// 1. 使用 comma-ok 断言
	d, ok := animal.(Dog)
	if ok {
		fmt.Println("断言成功: ", d.Name)
	} else {
		fmt.Println("断言失败")
	}

	// 2. 使用 panic 断言，注意要使用 `_` 忽略未使用的值
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic: ", r)
		}
	}()

	c := animal.(Cat) // 触发panic
	fmt.Println(c)    // 该行代码不会被执行，因为上面的代码触发panic

}
