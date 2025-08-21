package main

import "fmt"

func testArray() {
	ary1 := [3]int{1, 2, 3}
	aryCopy := ary1
	aryCopy[1] = 99
	fmt.Println("original array :", ary1)
	fmt.Println("assigned array", aryCopy)

}

func testStruct() {
	type Person struct {
		Name string
	}
	var p1 = Person{Name: "Hank"}
	p2 := p1
	p2.Name = "Ham"

	fmt.Println("original Person :", p1)
	fmt.Println("assigned Person", p2)

}

func main() {
	// Closure 陷阱示例
	var funcs []func() int

	// 錯誤的方式
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() int {
			return i // 注意：這裡會捕獲到最終的 i 值
		})
	}

	fmt.Println("錯誤的 closure:")
	for j, f := range funcs {
		fmt.Printf("funcs[%d]() = %d\n", j, f()) // 全部都是 3
	}

	testArray()
	testStruct()
}
