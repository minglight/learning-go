<!-- MARKDOWN_CELL -->
# 測試範例

這是一個簡單的測試，驗證轉換器是否正常運作。

**測試功能：**
- Markdown 解析
- Code cell 轉換
- 多個 cells 處理

<!-- CODE_CELL -->
```go
/* 簡單範例 - 變數宣告與輸出 */
package main

import "fmt"

func main() {
    message := "轉換成功!"
    fmt.Println(message)
}
// 輸出: 轉換成功!
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 第二個 Section

讓我們測試更多 Go 程式碼...

<!-- CODE_CELL -->
```go
/* 測試函式定義 */
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Printf("3 + 5 = %d\n", result)
}
// 輸出: 3 + 5 = 8
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 總結

如果你能看到這個 notebook 正常顯示，那麼轉換器就成功了！✨
