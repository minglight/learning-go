package main

import "fmt"

// 宣告常數
const (
	PI       = 3.14159
	MAX_SIZE = 100
	APP_NAME = "Go練習程式"
)

func main2() {
	// 1. 宣告常數與變數，印出常數與變數的值
	fmt.Println("=== 常數與變數 ===")
	var age int = 25
	var name string = "小明"
	var height float64 = 175.5

	fmt.Printf("常數 PI: %.5f\n", PI)
	fmt.Printf("常數 MAX_SIZE: %d\n", MAX_SIZE)
	fmt.Printf("常數 APP_NAME: %s\n", APP_NAME)
	fmt.Printf("變數 age: %d\n", age)
	fmt.Printf("變數 name: %s\n", name)
	fmt.Printf("變數 height: %.1f\n", height)

	// 2. 使用 if...else 判斷變數是否大於某個數字
	fmt.Println("\n=== if...else 判斷 ===")
	threshold := 20
	if age > threshold {
		fmt.Printf("年齡 %d 大於 %d\n", age, threshold)
	} else {
		fmt.Printf("年齡 %d 不大於 %d\n", age, threshold)
	}

	// 3. 使用 for 迴圈印出 1-10 的數字
	fmt.Println("\n=== for 迴圈印出 1-10 ===")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 4. 建立切片（Slice），動態加入數字後印出來
	fmt.Println("\n=== 切片（Slice）操作 ===")
	var numbers []int
	fmt.Printf("初始切片: %v\n", numbers)

	// 動態加入數字
	numbers = append(numbers, 10, 20, 30)
	fmt.Printf("加入 10, 20, 30 後: %v\n", numbers)

	numbers = append(numbers, 40, 50)
	fmt.Printf("再加入 40, 50 後: %v\n", numbers)

	// 5. 建立映射（Map）並新增、修改、刪除元素，最後印出結果
	fmt.Println("\n=== 映射（Map）操作 ===")

	// 建立 map
	students := make(map[string]int)
	fmt.Printf("初始 map: %v\n", students)

	// 新增元素
	students["Alice"] = 85
	students["Bob"] = 92
	students["Charlie"] = 78
	fmt.Printf("新增元素後: %v\n", students)

	// 修改元素
	students["Alice"] = 90
	fmt.Printf("修改 Alice 分數後: %v\n", students)

	// 刪除元素
	delete(students, "Charlie")
	fmt.Printf("刪除 Charlie 後: %v\n", students)

	// 檢查元素是否存在
	if score, exists := students["Bob"]; exists {
		fmt.Printf("Bob 的分數: %d\n", score)
	} else {
		fmt.Println("找不到 Bob 的分數")
	}
}
