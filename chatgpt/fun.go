package main

import (
	"errors"
	"fmt"
)

// 1. 寫一個函數可同時回傳商、餘數與 error（除數為 0 時）
func divide(dividend int, divisor int) (quotient int, remainder int, err error) {
	if divisor == 0 {
		return 0, 0, errors.New("除數不能為 0")
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return quotient, remainder, nil
}

func main() {
	fmt.Println("=== 除法函數測試 ===")

	// 測試正常情況
	q, r, err := divide(10, 3)
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 3 = %d 餘 %d\n", q, r)
	}

	// 測試除數為 0 的情況
	q2, r2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("錯誤: %v\n", err2)
	} else {
		fmt.Printf("10 ÷ 0 = %d 餘 %d\n", q2, r2)
	}

	// 更多測試案例
	testCases := []struct {
		dividend, divisor int
	}{
		{15, 4},
		{20, 5},
		{7, 2},
	}

	for _, tc := range testCases {
		q, r, err := divide(tc.dividend, tc.divisor)
		if err != nil {
			fmt.Printf("錯誤: %v\n", err)
		} else {
			fmt.Printf("%d ÷ %d = %d 餘 %d\n", tc.dividend, tc.divisor, q, r)
		}
	}
}
