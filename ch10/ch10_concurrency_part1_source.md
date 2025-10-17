<!-- MARKDOWN_CELL -->
# 第十章 Go 的並行 (Part 1: 基礎概念)

Go 語言的並行(concurrency)是其最強大的特性之一。Go 提供了 goroutine 和 channel 這兩個核心概念來處理並行程式設計。並行不同於並列(parallelism):並行是關於同時處理多個任務的能力,而並列是關於同時執行多個任務。

Go 的並行模型基於 CSP(Communicating Sequential Processes)理論,強調透過通訊來共享記憶體,而不是透過共享記憶體來通訊。
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
## 何時該使用並行

並行程式設計並不總是必要的,需要根據實際情況來決定是否使用。以下情況適合使用並行:

1. **I/O 密集型操作**:網路請求、檔案讀寫、資料庫查詢等
2. **獨立的任務處理**:可以同時處理多個不相關的任務
3. **生產者-消費者模式**:需要處理資料流的情況
4. **提升使用者體驗**:避免阻塞使用者介面
5. **利用多核心處理器**:CPU 密集型任務的並列處理

需要注意的是,並行也會帶來複雜性,包括資料競爭、死鎖等問題。
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 不使用並行的順序執行範例 */
func sequentialTasks() {
    fmt.Println("開始順序執行任務")
    start := time.Now()

    // 模擬三個耗時的任務
    task1()
    task2()
    task3()

    fmt.Printf("順序執行總時間: %v\n", time.Since(start)) // 輸出: 約 3 秒
}

func task1() {
    time.Sleep(1 * time.Second)
    fmt.Println("任務 1 完成")
}

func task2() {
    time.Sleep(1 * time.Second)
    fmt.Println("任務 2 完成")
}

func task3() {
    time.Sleep(1 * time.Second)
    fmt.Println("任務 3 完成")
}

func main() {
    sequentialTasks()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 使用並行的同時執行範例 */
func concurrentTasks() {
    fmt.Println("開始並行執行任務")
    start := time.Now()

    // 使用 channel 來等待所有任務完成
    done := make(chan bool, 3)

    // 同時啟動三個 goroutine
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("並行任務 1 完成")
        done <- true
    }()

    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("並行任務 2 完成")
        done <- true
    }()

    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("並行任務 3 完成")
        done <- true
    }()

    // 等待所有任務完成
    for i := 0; i < 3; i++ {
        <-done
    }

    fmt.Printf("並行執行總時間: %v\n", time.Since(start)) // 輸出: 約 1 秒
}

func main() {
    concurrentTasks()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## goroutine

Goroutine 是 Go 中的輕量級執行緒,由 Go runtime 管理。與傳統作業系統執行緒相比,goroutine 的建立成本非常低,記憶體占用小(初始只有 2KB 的堆疊),可以輕鬆建立成千上萬個 goroutine。

Goroutine 的特點:
- 使用 `go` 關鍵字啟動
- 在背景執行,不會阻塞主程式
- 由 Go scheduler 管理,採用 M:N 模型
- 支援協作式調度
- 當主 goroutine 結束時,所有其他 goroutine 也會終止
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 基本 goroutine 使用 */
func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello %s! (%d)\n", name, i+1)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    fmt.Println("開始執行程式")

    // 啟動一個 goroutine
    go sayHello("Alice")

    // 啟動另一個 goroutine
    go sayHello("Bob")

    // 主 goroutine 繼續執行
    fmt.Println("主程式繼續執行")

    // 等待 goroutines 執行完成
    time.Sleep(1 * time.Second)

    fmt.Println("程式結束")
    // 輸出會交錯顯示,因為 goroutines 並行執行
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "runtime"
)

/* 使用匿名函式建立 goroutine */
func main() {
    fmt.Printf("初始 goroutine 數量: %d\n", runtime.NumGoroutine()) // 輸出: 初始 goroutine 數量: 1

    // 使用匿名函式建立 goroutine
    for i := 0; i < 5; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d 開始執行\n", id)
            // 模擬一些工作
            for j := 0; j < 1000000; j++ {
                // 簡單的計算
            }
            fmt.Printf("Goroutine %d 執行完成\n", id)
        }(i)
    }

    fmt.Printf("建立後 goroutine 數量: %d\n", runtime.NumGoroutine()) // 輸出: 建立後 goroutine 數量: 6

    // 讓 goroutines 有時間執行
    runtime.Gosched() // 讓出 CPU 時間給其他 goroutines

    fmt.Println("主 goroutine 結束")
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

/* 使用 WaitGroup 確保 goroutines 執行完成 */
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 函式結束時通知 WaitGroup

    fmt.Printf("Worker %d 開始工作\n", id)

    // 模擬工作時間
    time.Sleep(time.Duration(id*100) * time.Millisecond)

    fmt.Printf("Worker %d 完成工作\n", id)
}

func main() {
    var wg sync.WaitGroup

    numWorkers := 3
    wg.Add(numWorkers) // 設定要等待的 goroutine 數量

    for i := 1; i <= numWorkers; i++ {
        go worker(i, &wg)
    }

    fmt.Println("等待所有 workers 完成...")
    wg.Wait() // 等待所有 goroutines 執行完成

    fmt.Println("所有工作完成!")
    // 輸出:
    // 等待所有 workers 完成...
    // Worker 1 開始工作
    // Worker 2 開始工作
    // Worker 3 開始工作
    // Worker 1 完成工作
    // Worker 2 完成工作
    // Worker 3 完成工作
    // 所有工作完成!
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## channel

Channel 是 Go 中用於 goroutine 之間通訊的管道。它基於 CSP 模型,讓 goroutines 能夠安全地交換資料。Channel 是型態安全的,只能傳遞特定型態的值。

Channel 的基本概念:
- 使用 `make(chan Type)` 建立
- 使用 `<-` 運算子進行讀寫操作
- 預設是同步的(無緩衝)
- 可以設定緩衝區大小
- 可以關閉 channel
- 支援方向性(只讀或只寫)
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import "fmt"

/* Channel 基本使用 - 簡單的資料傳遞 */
func main() {
    // 建立一個傳遞 string 的 channel
    messages := make(chan string)

    // 在 goroutine 中發送資料
    go func() {
        messages <- "Hello from goroutine!" // 發送資料到 channel
    }()

    // 在主 goroutine 中接收資料
    msg := <-messages // 從 channel 接收資料
    fmt.Println("收到訊息:", msg) // 輸出: 收到訊息: Hello from goroutine!
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* Channel 作為函式參數傳遞 */
func producer(ch chan<- int) { // 只能寫入的 channel
    for i := 1; i <= 5; i++ {
        fmt.Printf("生產者發送: %d\n", i)
        ch <- i
        time.Sleep(100 * time.Millisecond)
    }
    close(ch) // 關閉 channel 表示沒有更多資料
}

func consumer(ch <-chan int) { // 只能讀取的 channel
    for value := range ch { // 使用 range 讀取直到 channel 關閉
        fmt.Printf("消費者接收: %d\n", value)
        time.Sleep(200 * time.Millisecond)
    }
    fmt.Println("消費者完成")
}

func main() {
    ch := make(chan int)

    go producer(ch)
    go consumer(ch)

    // 等待足夠時間讓 goroutines 執行完成
    time.Sleep(2 * time.Second)

    fmt.Println("程式結束")
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 讀取、寫入與緩衝

Channel 的讀寫操作使用 `<-` 運算子。無緩衝的 channel 是同步的,發送和接收操作會相互阻塞直到另一方準備好。有緩衝的 channel 允許異步操作,直到緩衝區滿或空。

操作特點:
- 發送:`ch <- value`
- 接收:`value := <-ch` 或 `<-ch`
- 無緩衝:同步操作,必須有接收者才能發送
- 有緩衝:異步操作,可以發送到緩衝區
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 無緩衝 channel 的同步行為 */
func unbufferedDemo() {
    ch := make(chan string) // 無緩衝 channel

    go func() {
        fmt.Println("Goroutine 準備發送資料...")
        ch <- "同步訊息" // 這行會阻塞直到有接收者
        fmt.Println("Goroutine 發送完成")
    }()

    time.Sleep(1 * time.Second) // 模擬主程式忙碌
    fmt.Println("主程式準備接收...")

    message := <-ch // 接收資料,此時 goroutine 才能繼續
    fmt.Println("接收到:", message)

    time.Sleep(100 * time.Millisecond) // 等待 goroutine 完成輸出
}

func main() {
    unbufferedDemo()
    // 輸出順序:
    // Goroutine 準備發送資料...
    // (1秒後)
    // 主程式準備接收...
    // 接收到: 同步訊息
    // Goroutine 發送完成
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 有緩衝 channel 的異步行為 */
func bufferedDemo() {
    ch := make(chan string, 2) // 緩衝大小為 2 的 channel

    go func() {
        fmt.Println("發送第一個訊息")
        ch <- "訊息 1" // 不會阻塞,放入緩衝區

        fmt.Println("發送第二個訊息")
        ch <- "訊息 2" // 不會阻塞,放入緩衝區

        fmt.Println("發送第三個訊息")
        ch <- "訊息 3" // 會阻塞,因為緩衝區已滿

        fmt.Println("所有訊息發送完成")
    }()

    time.Sleep(1 * time.Second) // 讓 goroutine 先執行

    fmt.Println("開始接收訊息")
    fmt.Println("接收:", <-ch) // 接收: 訊息 1
    fmt.Println("接收:", <-ch) // 接收: 訊息 2
    fmt.Println("接收:", <-ch) // 接收: 訊息 3,此時 goroutine 才能完成

    time.Sleep(100 * time.Millisecond)
}

func main() {
    bufferedDemo()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import "fmt"

/* Channel 的容量和長度 */
func channelProperties() {
    // 無緩衝 channel
    unbuffered := make(chan int)
    fmt.Printf("無緩衝 channel - 容量: %d, 長度: %d\n", cap(unbuffered), len(unbuffered))
    // 輸出: 無緩衝 channel - 容量: 0, 長度: 0

    // 有緩衝 channel
    buffered := make(chan int, 3)
    fmt.Printf("空的有緩衝 channel - 容量: %d, 長度: %d\n", cap(buffered), len(buffered))
    // 輸出: 空的有緩衝 channel - 容量: 3, 長度: 0

    // 添加一些資料
    buffered <- 1
    buffered <- 2
    fmt.Printf("部分填充的 channel - 容量: %d, 長度: %d\n", cap(buffered), len(buffered))
    // 輸出: 部分填充的 channel - 容量: 3, 長度: 2

    // 填滿 channel
    buffered <- 3
    fmt.Printf("已滿的 channel - 容量: %d, 長度: %d\n", cap(buffered), len(buffered))
    // 輸出: 已滿的 channel - 容量: 3, 長度: 3
}

func main() {
    channelProperties()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### for-range 與 channel

`for-range` 迴圈可以用來從 channel 中持續讀取資料,直到 channel 被關閉。這是處理 channel 資料流的慣用方式。當 channel 被關閉且沒有更多資料時,迴圈會自動結束。

使用方式:
- `for value := range ch`:接收值並檢查 channel 是否關閉
- 當 channel 關閉時,迴圈自動結束
- 適合處理不定數量的資料流
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 使用 for-range 處理 channel 資料 */
func numberGenerator(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("生成數字: %d\n", i)
        ch <- i
        time.Sleep(200 * time.Millisecond)
    }
    close(ch) // 關閉 channel,通知沒有更多資料
}

func main() {
    numbers := make(chan int)

    go numberGenerator(numbers)

    fmt.Println("開始接收數字:")

    // 使用 for-range 接收所有資料
    for num := range numbers {
        fmt.Printf("接收到數字: %d\n", num)
        time.Sleep(100 * time.Millisecond)
    }

    fmt.Println("所有數字接收完成")
    // 輸出:
    // 開始接收數字:
    // 生成數字: 1
    // 接收到數字: 1
    // 生成數字: 2
    // 接收到數字: 2
    // ...
    // 所有數字接收完成
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import "fmt"

/* 比較不同的 channel 讀取方式 */
func demonstrateChannelReading() {
    ch := make(chan string, 3)

    // 發送一些資料
    ch <- "第一個"
    ch <- "第二個"
    ch <- "第三個"
    close(ch)

    fmt.Println("方法 1: 使用 for-range")
    // 重新建立 channel 因為上面的已經被讀完了
    ch1 := make(chan string, 3)
    ch1 <- "A"
    ch1 <- "B"
    ch1 <- "C"
    close(ch1)

    for value := range ch1 {
        fmt.Println("range 讀取:", value)
    }

    fmt.Println("\n方法 2: 使用 for 迴圈配合 ok")
    ch2 := make(chan string, 3)
    ch2 <- "X"
    ch2 <- "Y"
    ch2 <- "Z"
    close(ch2)

    for {
        value, ok := <-ch2
        if !ok {
            fmt.Println("channel 已關閉")
            break
        }
        fmt.Println("手動讀取:", value)
    }
}

func main() {
    demonstrateChannelReading()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 關閉 channel

關閉 channel 使用 `close()` 函式,它會通知接收者沒有更多資料會被發送。關閉 channel 不會清除其中的資料,接收者仍然可以讀取剩餘的資料。

關閉 channel 的重要概念:
- 只有發送者應該關閉 channel
- 向已關閉的 channel 發送資料會引發 panic
- 從已關閉的 channel 讀取會立即返回零值和 false
- 可以使用 `value, ok := <-ch` 檢查 channel 是否關閉
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import "fmt"

/* Channel 關閉的基本操作 */
func closeDemo() {
    ch := make(chan int, 2)

    // 發送一些資料
    ch <- 1
    ch <- 2

    fmt.Println("關閉前 channel 長度:", len(ch)) // 輸出: 關閉前 channel 長度: 2

    // 關閉 channel
    close(ch)

    fmt.Println("關閉後 channel 長度:", len(ch)) // 輸出: 關閉後 channel 長度: 2

    // 仍然可以讀取剩餘資料
    value1, ok1 := <-ch
    fmt.Printf("讀取 1: 值=%d, ok=%t\n", value1, ok1) // 輸出: 讀取 1: 值=1, ok=true

    value2, ok2 := <-ch
    fmt.Printf("讀取 2: 值=%d, ok=%t\n", value2, ok2) // 輸出: 讀取 2: 值=2, ok=true

    // 所有資料讀完後,再讀取會得到零值
    value3, ok3 := <-ch
    fmt.Printf("讀取 3: 值=%d, ok=%t\n", value3, ok3) // 輸出: 讀取 3: 值=0, ok=false
}

func main() {
    closeDemo()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 優雅關閉 channel 的模式 */
func gracefulShutdown() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // Worker goroutine
    go func() {
        for {
            job, ok := <-jobs
            if !ok {
                fmt.Println("沒有更多工作,worker 退出")
                done <- true
                return
            }
            fmt.Printf("處理工作: %d\n", job)
            time.Sleep(100 * time.Millisecond)
        }
    }()

    // 發送工作
    for i := 1; i <= 3; i++ {
        jobs <- i
    }

    // 關閉 jobs channel,通知沒有更多工作
    close(jobs)
    fmt.Println("已關閉 jobs channel")

    // 等待 worker 完成
    <-done
    fmt.Println("優雅關閉完成")
}

func main() {
    gracefulShutdown()
    // 輸出:
    // 已關閉 jobs channel
    // 處理工作: 1
    // 處理工作: 2
    // 處理工作: 3
    // 沒有更多工作,worker 退出
    // 優雅關閉完成
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### channel 的行為

Channel 在不同狀態下有不同的行為模式。理解這些行為對於編寫正確的並行程式至關重要。

Channel 行為總結:
- **nil channel**:讀寫都會永遠阻塞
- **開啟的無緩衝 channel**:讀寫必須同時進行
- **開啟的有緩衝 channel**:有空間時寫入不阻塞,有資料時讀取不阻塞
- **已關閉的 channel**:可以讀取剩餘資料,寫入會 panic
- **已關閉且空的 channel**:讀取返回零值和 false
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* Channel 不同狀態的行為示範 */
func channelBehaviors() {
    fmt.Println("=== 無緩衝 channel 行為 ===")
    unbuffered := make(chan int)

    // 嘗試寫入無緩衝 channel(會在 goroutine 中執行避免阻塞)
    go func() {
        fmt.Println("準備寫入無緩衝 channel...")
        unbuffered <- 42
        fmt.Println("無緩衝 channel 寫入完成")
    }()

    time.Sleep(100 * time.Millisecond) // 讓寫入操作先執行
    fmt.Println("準備從無緩衝 channel 讀取...")
    value := <-unbuffered
    fmt.Printf("從無緩衝 channel 讀取: %d\n\n", value)

    fmt.Println("=== 有緩衝 channel 行為 ===")
    buffered := make(chan int, 2)

    // 寫入有緩衝 channel(不會阻塞)
    fmt.Println("寫入第一個值到有緩衝 channel")
    buffered <- 1
    fmt.Println("寫入第二個值到有緩衝 channel")
    buffered <- 2
    fmt.Printf("緩衝 channel 狀態 - 長度: %d, 容量: %d\n", len(buffered), cap(buffered))

    // 讀取有緩衝 channel
    val1 := <-buffered
    val2 := <-buffered
    fmt.Printf("從有緩衝 channel 讀取: %d, %d\n\n", val1, val2)

    fmt.Println("=== 已關閉 channel 行為 ===")
    closed := make(chan int, 1)
    closed <- 100
    close(closed)

    // 從已關閉的 channel 讀取剩餘資料
    value1, ok1 := <-closed
    fmt.Printf("從已關閉 channel 讀取: 值=%d, ok=%t\n", value1, ok1)

    // 再次讀取已關閉且空的 channel
    value2, ok2 := <-closed
    fmt.Printf("從已關閉且空的 channel 讀取: 值=%d, ok=%t\n", value2, ok2)
}

func main() {
    channelBehaviors()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import "fmt"

/* Channel 方向性示範 */
func channelDirections() {
    ch := make(chan string, 1)

    // 只寫 channel 函式
    sendOnly := func(ch chan<- string, msg string) {
        fmt.Println("發送訊息:", msg)
        ch <- msg
        // value := <-ch // 編譯錯誤:不能從只寫 channel 讀取
    }

    // 只讀 channel 函式
    receiveOnly := func(ch <-chan string) string {
        msg := <-ch
        fmt.Println("接收訊息:", msg)
        // ch <- "test" // 編譯錯誤:不能向只讀 channel 寫入
        return msg
    }

    // 使用不同方向的 channel
    sendOnly(ch, "Hello, World!")
    result := receiveOnly(ch)

    fmt.Println("處理結果:", result)
    // 輸出:
    // 發送訊息: Hello, World!
    // 接收訊息: Hello, World!
    // 處理結果: Hello, World!
}

func main() {
    channelDirections()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## select

`select` 語句是 Go 並行程式設計的強大工具,類似於 `switch` 語句,但專門用於 channel 操作。它可以同時等待多個 channel 操作,哪個操作先準備好就執行哪個。

Select 的特點:
- 可以同時監聽多個 channel
- 隨機選擇一個準備好的 case
- 支援 default case 用於非阻塞操作
- 可用於實現超時機制
- 每次執行只會選擇一個 case
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* Select 基本用法 */
func basicSelect() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // 第一個 goroutine,1秒後發送資料
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "來自 channel 1"
    }()

    // 第二個 goroutine,2秒後發送資料
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "來自 channel 2"
    }()

    // 使用 select 等待第一個準備好的 channel
    select {
    case msg1 := <-ch1:
        fmt.Println("接收到:", msg1) // 這個會先執行
    case msg2 := <-ch2:
        fmt.Println("接收到:", msg2)
    }

    fmt.Println("基本 select 完成")
    // 輸出:
    // (1秒後) 接收到: 來自 channel 1
    // 基本 select 完成
}

func main() {
    basicSelect()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 使用 Select 實現超時機制 */
func timeoutSelect() {
    ch := make(chan string)

    // 模擬一個可能很慢的操作
    go func() {
        time.Sleep(2 * time.Second) // 模擬2秒的工作
        ch <- "工作完成"
    }()

    fmt.Println("等待工作完成,最多等待 1 秒...")

    select {
    case result := <-ch:
        fmt.Println("收到結果:", result)
    case <-time.After(1 * time.Second): // 1秒後觸發
        fmt.Println("操作超時!") // 這個會執行,因為1秒 < 2秒
    }

    fmt.Println("超時 select 完成")
    // 輸出:
    // 等待工作完成,最多等待 1 秒...
    // (1秒後) 操作超時!
    // 超時 select 完成
}

func main() {
    timeoutSelect()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 使用 Default Case 的非阻塞 Select */
func nonBlockingSelect() {
    ch := make(chan string)

    // 嘗試非阻塞讀取
    select {
    case msg := <-ch:
        fmt.Println("接收到訊息:", msg)
    default:
        fmt.Println("沒有訊息可讀取,繼續其他工作") // 這個會執行
    }

    // 嘗試非阻塞寫入(對無緩衝 channel)
    select {
    case ch <- "測試訊息":
        fmt.Println("成功發送訊息")
    default:
        fmt.Println("無法發送訊息,沒有接收者") // 這個會執行
    }

    // 使用有緩衝 channel 進行非阻塞操作
    bufferedCh := make(chan int, 1)

    select {
    case bufferedCh <- 42:
        fmt.Println("成功發送到有緩衝 channel") // 這個會執行
    default:
        fmt.Println("緩衝 channel 已滿")
    }

    select {
    case value := <-bufferedCh:
        fmt.Println("從緩衝 channel 讀取:", value) // 這個會執行,輸出: 42
    default:
        fmt.Println("緩衝 channel 為空")
    }
}

func main() {
    nonBlockingSelect()
    // 輸出:
    // 沒有訊息可讀取,繼續其他工作
    // 無法發送訊息,沒有接收者
    // 成功發送到有緩衝 channel
    // 從緩衝 channel 讀取: 42
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 使用 Select 實現多工處理 */
func multiplexingSelect() {
    // 建立多個 channel
    orders := make(chan string)
    payments := make(chan string)
    notifications := make(chan string)
    quit := make(chan bool)

    // 模擬不同來源的資料
    go func() {
        time.Sleep(1 * time.Second)
        orders <- "新訂單 #1001"
    }()

    go func() {
        time.Sleep(1500 * time.Millisecond)
        payments <- "付款 $100"
    }()

    go func() {
        time.Sleep(500 * time.Millisecond)
        notifications <- "系統通知"
    }()

    go func() {
        time.Sleep(3 * time.Second)
        quit <- true
    }()

    fmt.Println("開始處理多種事件...")

    // 使用 select 處理多種事件
    for {
        select {
        case order := <-orders:
            fmt.Println("處理訂單:", order)
        case payment := <-payments:
            fmt.Println("處理付款:", payment)
        case notification := <-notifications:
            fmt.Println("處理通知:", notification)
        case <-quit:
            fmt.Println("收到退出信號,停止處理")
            return
        case <-time.After(2 * time.Second):
            fmt.Println("等待事件中...")
        }
    }
}

func main() {
    multiplexingSelect()
    // 輸出:
    // 開始處理多種事件...
    // (500ms後) 處理通知: 系統通知
    // (1s後) 處理訂單: 新訂單 #1001
    // (1.5s後) 處理付款: 付款 $100
    // (2s後) 等待事件中...
    // (3s後) 收到退出信號,停止處理
}
```
<!-- END_CODE_CELL -->
