# learning-go

這是一個學習Go的Project, 主要有三個任務, 當我要求時你要根據我的要求做不同的事情
1. 產生jupyter筆記 : 
- 分章節產生jupyter學習筆記
- 要完全根據指定的那一章的目錄產生筆記. 
- 每一小節都至少要有從初階到進階的詳細說明 as markdown block, 然後接著code blocks做詳細的範例, 每一段展示code都要有comment說明
- 如果有特別注意的點, 或是容易出錯的點都要說明

參考資訊 
- 要觀看目錄 ( ./目錄.md)
- 要看筆記 example (/.ch2/ch2_basic_types_declarations.ipynb)
- 要看考題 example (./ch2/ch2_interview_questions.ipynb)
- 要看考題答案example (./ch2/ch2_interview_answers_review.ipynb)
ex of 說明 : 
## 1. 內建的型態

Go 提供了豐富的內建型態，包括：

### 基本型態
- **布林型態**: `bool`
- **字串型態**: `string`
- **數字型態**:
  - 整數: `int`, `int8`, `int16`, `int32`, `int64`
  - 無符號整數: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - 浮點數: `float32`, `float64`
  - 複數: `complex64`, `complex128`
- **位元組**: `byte` (uint8 的別名)
- **Unicode 字元**: `rune` (int32 的別名)

ex of code block 

```go
package main

import "fmt"

func main() {
    // 展示各種內建型態
    var b bool = true
    var s string = "Hello, Go!"
    var i int = 42
    var f float64 = 3.14
    var c complex128 = 1 + 2i
    
    fmt.Printf("bool: %v\n", b)
    fmt.Printf("string: %v\n", s)
    fmt.Printf("int: %v\n", i)
    fmt.Printf("float64: %v\n", f)
    fmt.Printf("complex128: %v\n", c)
}
```

2. 產生面試考題
在同一個folder下產生一個 .ipynb file,  並且要寫一些面試常出現的考題, 你使用markdown寫題目, 題目要有標題跟內容, 然後要留 code block讓我寫答案.

3. 改考卷
當我寫完之後我會請你改考卷, 你要另外開一個jupyter file,  一開始要先寫summary, 說明哪裡寫錯或寫得不好,最後再一題一題講解copy我寫不好或寫錯的那一題的題目跟我寫的答案, 然後新增你的說明跟正確解答

