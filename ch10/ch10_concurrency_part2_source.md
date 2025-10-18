<!-- MARKDOWN_CELL -->
# 第十章 Go 的並行 (Part 2: 進階應用)

這一部分將深入探討 Go 並行程式設計的進階概念和實踐模式。我們將學習如何在實際應用中正確使用 goroutines 和 channels,避免常見的陷阱,並掌握各種並行模式的最佳實踐。
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
## 並行實踐模式

並行程式設計有許多常見的模式和最佳實踐。理解這些模式可以幫助你編寫更安全、更高效的並行程式,避免常見的錯誤如資料競爭、死鎖和 goroutine 洩漏等問題。

本節將涵蓋以下重要模式:
- API 設計中的並行考量
- Goroutine 生命週期管理
- 各種同步和協調機制
- 錯誤處理和資源清理
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
### 不要運行放入 API

這個原則是指在設計 API 時,不要讓 API 函式內部自動啟動 goroutines,除非這是 API 的核心功能。讓呼叫者決定是否需要並行執行,這樣可以:

- 給呼叫者更多控制權
- 避免不必要的 goroutine 建立
- 讓 API 更容易測試
- 減少資源洩漏的風險

如果 API 確實需要內部並行,應該提供明確的控制機制。
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 錯誤的 API 設計:內部自動啟動 goroutine */
type BadProcessor struct{}

func (p *BadProcessor) ProcessData(data string) {
    // 錯誤:自動啟動 goroutine,呼叫者無法控制
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("處理完成:", data)
    }()
    // 函式立即返回,呼叫者不知道處理何時完成
}

func badAPIExample() {
    processor := &BadProcessor{}
    processor.ProcessData("測試資料")

    // 問題:不知道處理何時完成,無法等待
    fmt.Println("函式已返回,但不知道處理是否完成")

    // 必須猜測等待時間
    time.Sleep(1500 * time.Millisecond)
}

func main() {
    badAPIExample()
    // 輸出:
    // 函式已返回,但不知道處理是否完成
    // (1秒後) 處理完成: 測試資料
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

/* 良好的 API 設計:讓呼叫者控制並行 */
type GoodProcessor struct{}

// 同步版本:阻塞直到完成
func (p *GoodProcessor) ProcessData(data string) {
    time.Sleep(1 * time.Second) // 模擬處理時間
    fmt.Println("同步處理完成:", data)
}

// 異步版本:返回 channel 讓呼叫者控制
func (p *GoodProcessor) ProcessDataAsync(data string) <-chan string {
    result := make(chan string, 1)

    go func() {
        defer close(result)
        time.Sleep(1 * time.Second)
        result <- fmt.Sprintf("異步處理完成: %s", data)
    }()

    return result
}

func goodAPIExample() {
    processor := &GoodProcessor{}

    // 呼叫者可以選擇同步執行
    fmt.Println("=== 同步執行 ===")
    processor.ProcessData("同步資料")

    // 或者選擇異步執行
    fmt.Println("\n=== 異步執行 ===")
    resultCh := processor.ProcessDataAsync("異步資料")

    // 呼叫者可以決定何時等待結果
    fmt.Println("可以在等待時做其他事情...")
    time.Sleep(500 * time.Millisecond)
    fmt.Println("現在等待異步結果")

    result := <-resultCh
    fmt.Println("收到結果:", result)
}

func main() {
    goodAPIExample()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### goroutine、for 迴圈,更會變的變數

這是 Go 並行程式設計中最常見的陷阱之一。在迴圈中啟動 goroutines 時,如果直接使用迴圈變數,所有 goroutines 可能會使用相同的值(通常是最後一次迭代的值)。

問題的原因:
- 迴圈變數在所有迭代中是同一個記憶體位置
- Goroutines 是異步執行的,可能在迴圈結束後才讀取變數
- 所有 goroutines 看到的是變數的最終值

解決方法:
- 將迴圈變數作為參數傳遞給 goroutine
- 在迴圈內部建立本地變數副本
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

/* 錯誤示範:直接使用迴圈變數 */
func wrongWayDemo() {
    fmt.Println("=== 錯誤方式:直接使用迴圈變數 ===")
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 錯誤:直接使用迴圈變數 i
            // 所有 goroutines 很可能看到的都是 i=5(迴圈結束後的值)
            fmt.Printf("Goroutine 看到的 i 值: %d\n", i)
        }()
    }

    wg.Wait()
    fmt.Println("錯誤方式完成\n")
}

func main() {
    wrongWayDemo()
    // 可能的輸出(通常所有 goroutines 都輸出 5):
    // Goroutine 看到的 i 值: 5
    // Goroutine 看到的 i 值: 5
    // Goroutine 看到的 i 值: 5
    // Goroutine 看到的 i 值: 5
    // Goroutine 看到的 i 值: 5
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
)

/* 正確方式 1:將變數作為參數傳遞 */
func correctWayWithParameter() {
    fmt.Println("=== 正確方式 1:參數傳遞 ===")
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) { // 將 i 作為參數傳遞
            defer wg.Done()
            fmt.Printf("Goroutine 收到參數: %d\n", id)
        }(i) // 傳遞當前的 i 值
    }

    wg.Wait()
    fmt.Println("參數傳遞方式完成\n")
}

func main() {
    correctWayWithParameter()
    // 輸出(順序可能不同):
    // Goroutine 收到參數: 0
    // Goroutine 收到參數: 1
    // Goroutine 收到參數: 2
    // Goroutine 收到參數: 3
    // Goroutine 收到參數: 4
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
)

/* 正確方式 2:建立本地變數副本 */
func correctWayWithLocalVariable() {
    fmt.Println("=== 正確方式 2:本地變數副本 ===")
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)

        // 建立本地變數副本
        localI := i

        go func() {
            defer wg.Done()
            fmt.Printf("Goroutine 使用本地變數: %d\n", localI)
        }()
    }

    wg.Wait()
    fmt.Println("本地變數方式完成\n")
}

func main() {
    correctWayWithLocalVariable()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
)

/* 更複雜的例子:處理 slice 資料 */
func complexExample() {
    fmt.Println("=== 複雜例子:處理 slice 資料 ===")

    data := []string{"apple", "banana", "cherry", "date", "elderberry"}
    var wg sync.WaitGroup

    // 錯誤方式:直接使用 item 和 i
    fmt.Println("錯誤方式:")
    for i, item := range data {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 錯誤:可能所有 goroutines 都看到最後的值
            fmt.Printf("錯誤 - 索引: %d, 項目: %s\n", i, item)
        }()
    }
    wg.Wait()

    fmt.Println("\n正確方式:")
    for i, item := range data {
        wg.Add(1)
        go func(index int, value string) {
            defer wg.Done()
            // 正確:使用參數傳遞的值
            fmt.Printf("正確 - 索引: %d, 項目: %s\n", index, value)
        }(i, item)
    }
    wg.Wait()
}

func main() {
    complexExample()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 你一定要清理 goroutine

Goroutine 洩漏是並行程式中的嚴重問題。如果 goroutines 無法正常結束,它們會一直消耗記憶體和其他資源。因此,確保所有啟動的 goroutines 都能正常結束是非常重要的。

常見的 goroutine 洩漏原因:
- 無限迴圈沒有退出條件
- 阻塞在 channel 操作上
- 等待永遠不會發生的事件
- 沒有處理取消信號

防止洩漏的策略:
- 使用 context 來控制 goroutine 生命週期
- 設置明確的退出條件
- 使用 done channel 模式
- 設置超時機制
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

/* 問題示範:會造成 goroutine 洩漏的程式 */
func leakyGoroutines() {
    fmt.Printf("開始時的 goroutine 數量: %d\n", runtime.NumGoroutine())

    ch := make(chan string)

    // 啟動一些會洩漏的 goroutines
    for i := 0; i < 3; i++ {
        go func(id int) {
            // 這些 goroutines 會永遠阻塞在這裡
            // 因為沒有人會向 ch 發送資料
            msg := <-ch
            fmt.Printf("Goroutine %d 收到: %s\n", id, msg)
        }(i)
    }

    time.Sleep(100 * time.Millisecond) // 讓 goroutines 啟動

    fmt.Printf("啟動後的 goroutine 數量: %d\n", runtime.NumGoroutine())

    // 函式結束,但 goroutines 仍然在運行(洩漏)
    fmt.Println("函式結束,但 goroutines 仍在等待(洩漏)")
}

func main() {
    leakyGoroutines()

    time.Sleep(100 * time.Millisecond)
    fmt.Printf("main 函式中的 goroutine 數量: %d\n", runtime.NumGoroutine())

    // 輸出:
    // 開始時的 goroutine 數量: 1
    // 啟動後的 goroutine 數量: 4
    // 函式結束,但 goroutines 仍在等待(洩漏)
    // main 函式中的 goroutine 數量: 4 (應該是 1)
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### done channel 模式

Done channel 模式是控制 goroutine 生命週期的經典方法。透過一個專門的 channel 來通知 goroutines 何時應該停止工作。這種模式簡單易懂,廣泛用於各種並行程式中。

Done channel 的特點:
- 通常使用 `chan struct{}` 或 `chan bool`
- 發送者關閉 channel 來通知所有接收者
- 接收者使用 select 來檢查退出信號
- 可以同時通知多個 goroutines
- 輕量級,不傳遞實際資料
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

/* Done Channel 模式解決 goroutine 洩漏 */
func doneChannelPattern() {
    fmt.Printf("開始時的 goroutine 數量: %d\n", runtime.NumGoroutine())

    done := make(chan struct{}) // 使用空 struct 節省記憶體

    // 啟動多個 worker goroutines
    for i := 0; i < 3; i++ {
        go func(id int) {
            defer fmt.Printf("Goroutine %d 退出\n", id)

            ticker := time.NewTicker(200 * time.Millisecond)
            defer ticker.Stop()

            for {
                select {
                case <-done:
                    fmt.Printf("Goroutine %d 收到退出信號\n", id)
                    return // 正常退出
                case <-ticker.C:
                    fmt.Printf("Goroutine %d 正在工作\n", id)
                }
            }
        }(i)
    }

    time.Sleep(100 * time.Millisecond)
    fmt.Printf("啟動後的 goroutine 數量: %d\n", runtime.NumGoroutine())

    // 讓 goroutines 工作一段時間
    time.Sleep(500 * time.Millisecond)

    fmt.Println("發送退出信號...")
    close(done) // 關閉 done channel,通知所有 goroutines 退出

    // 等待 goroutines 退出
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("清理後的 goroutine 數量: %d\n", runtime.NumGoroutine())
}

func main() {
    doneChannelPattern()
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

/* Done Channel 與 WaitGroup 結合使用 */
func doneChannelWithWaitGroup() {
    var wg sync.WaitGroup
    done := make(chan struct{})

    // 啟動多個處理器
    numWorkers := 3
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            defer fmt.Printf("Worker %d 完成清理\n", workerID)

            workCount := 0
            ticker := time.NewTicker(150 * time.Millisecond)
            defer ticker.Stop()

            for {
                select {
                case <-done:
                    fmt.Printf("Worker %d 收到停止信號,完成了 %d 項工作\n", workerID, workCount)
                    return
                case <-ticker.C:
                    workCount++
                    fmt.Printf("Worker %d 完成第 %d 項工作\n", workerID, workCount)
                }
            }
        }(i)
    }

    // 讓 workers 工作一段時間
    fmt.Println("讓 workers 工作 1 秒...")
    time.Sleep(1 * time.Second)

    // 發送停止信號
    fmt.Println("\n發送停止信號並等待所有 workers 完成...")
    close(done)

    // 等待所有 workers 完成清理
    wg.Wait()
    fmt.Println("所有 workers 已安全退出")
}

func main() {
    doneChannelWithWaitGroup()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 使用取消函式來終結 goroutine

Go 的 `context` 套件提供了更強大和標準化的取消機制。Context 不僅可以用於取消操作,還可以設置截止時間、超時,以及傳遞請求範圍的值。

Context 的優勢:
- 標準化的取消介面
- 支援超時和截止時間
- 可以形成取消鏈
- 與許多標準函式庫整合
- 線程安全
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

/* 使用 context 控制 goroutine */
func contextCancellation() {
    // 建立可取消的 context
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 確保 context 被取消

    var wg sync.WaitGroup

    // 啟動多個 workers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }

    // 讓 workers 執行一段時間
    fmt.Println("讓 workers 執行 800ms...")
    time.Sleep(800 * time.Millisecond)

    // 取消所有 workers
    fmt.Println("\n取消所有 workers...")
    cancel()

    // 等待所有 workers 完成
    wg.Wait()
    fmt.Println("所有 workers 已完成")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Printf("Worker %d 退出\n", id)

    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()

    taskCount := 0

    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d 收到取消信號,完成了 %d 個任務\n", id, taskCount)
            return
        case <-ticker.C:
            taskCount++
            fmt.Printf("Worker %d 完成任務 #%d\n", id, taskCount)
        }
    }
}

func main() {
    contextCancellation()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "context"
    "fmt"
    "time"
)

/* 使用 context 設置超時 */
func contextWithTimeout() {
    // 建立有超時的 context(3秒後自動取消)
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    go longRunningTask(ctx, "長時間任務")

    // 等待任務完成或超時
    <-ctx.Done()

    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("任務因超時而取消")
    } else {
        fmt.Println("任務正常完成")
    }
}

func longRunningTask(ctx context.Context, taskName string) {
    fmt.Printf("開始執行 %s\n", taskName)

    // 模擬長時間運行的任務
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()

    step := 0

    for {
        select {
        case <-ctx.Done():
            fmt.Printf("%s 在第 %d 步被取消: %v\n", taskName, step, ctx.Err())
            return
        case <-ticker.C:
            step++
            fmt.Printf("%s 執行第 %d 步\n", taskName, step)

            // 假設任務在第10步完成
            if step >= 10 {
                fmt.Printf("%s 完成!\n", taskName)
                return
            }
        }
    }
}

func main() {
    contextWithTimeout()
    // 輸出會在3秒後顯示超時訊息,因為任務需要5秒才能完成(10步 × 0.5秒)
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 何時該使用有緩衝與無緩衝的 channel

選擇有緩衝還是無緩衝的 channel 是並行程式設計中的重要決策。每種類型都有其適用的場景和特點。

**無緩衝 channel 適用於:**
- 需要同步的場景
- 確保資料被立即處理
- 實現 handshake 機制
- 簡單的信號傳遞

**有緩衝 channel 適用於:**
- 解耦生產者和消費者
- 處理突發流量
- 實現工作佇列
- 避免 goroutine 阻塞
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 無緩衝 channel:同步通訊 */
func unbufferedChannelDemo() {
    fmt.Println("=== 無緩衝 channel 示範 ===")

    // 用於同步的 channel
    sync := make(chan struct{})

    go func() {
        fmt.Println("Goroutine: 開始執行重要任務...")
        time.Sleep(1 * time.Second) // 模擬工作
        fmt.Println("Goroutine: 重要任務完成!")

        // 發送完成信號(會阻塞直到主程式接收)
        sync <- struct{}{}
        fmt.Println("Goroutine: 確認主程式已收到完成信號")
    }()

    fmt.Println("主程式: 等待重要任務完成...")

    // 等待任務完成(會阻塞直到 goroutine 發送信號)
    <-sync

    fmt.Println("主程式: 收到完成信號,可以安全繼續")
    fmt.Println("主程式: 執行依賴於重要任務的操作\n")
}

func main() {
    unbufferedChannelDemo()
    // 輸出顯示嚴格的同步行為
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

/* 有緩衝 channel:非同步通訊和工作佇列 */
func bufferedChannelDemo() {
    fmt.Println("=== 有緩衝 channel 示範 ===")

    // 建立工作佇列
    jobs := make(chan string, 5) // 可以存放5個工作
    results := make(chan string, 5)

    // 啟動 worker
    go func() {
        for job := range jobs {
            fmt.Printf("Worker: 處理工作 '%s'\n", job)
            time.Sleep(300 * time.Millisecond) // 模擬處理時間
            results <- fmt.Sprintf("完成 %s", job)
        }
        close(results)
    }()

    // 快速提交多個工作(不會阻塞,因為有緩衝)
    fmt.Println("主程式: 快速提交工作到佇列...")
    jobs <- "任務A"
    jobs <- "任務B"
    jobs <- "任務C"
    fmt.Println("主程式: 所有工作已提交,可以做其他事")

    // 關閉工作佇列
    close(jobs)

    // 等待並收集結果
    fmt.Println("主程式: 收集處理結果...")
    for result := range results {
        fmt.Printf("主程式: 收到結果 - %s\n", result)
    }

    fmt.Println("主程式: 所有工作處理完成\n")
}

func main() {
    bufferedChannelDemo()
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

/* 比較不同緩衝大小的影響 */
func bufferSizeComparison() {
    fmt.Println("=== 緩衝大小對性能的影響 ===")

    // 測試小緩衝
    testBufferSize("小緩衝 (1)", 1)

    // 測試大緩衝
    testBufferSize("大緩衝 (10)", 10)
}

func testBufferSize(name string, bufferSize int) {
    fmt.Printf("\n--- %s ---\n", name)

    start := time.Now()
    ch := make(chan int, bufferSize)

    // 生產者
    go func() {
        for i := 1; i <= 5; i++ {
            fmt.Printf("生產者: 準備發送 %d\n", i)
            ch <- i
            fmt.Printf("生產者: 已發送 %d\n", i)
        }
        close(ch)
    }()

    // 消費者(故意慢一點)
    time.Sleep(200 * time.Millisecond) // 讓生產者先跑

    for value := range ch {
        fmt.Printf("消費者: 收到 %d\n", value)
        time.Sleep(100 * time.Millisecond) // 模擬處理時間
    }

    elapsed := time.Since(start)
    fmt.Printf("%s 總時間: %v\n", name, elapsed)
}

func main() {
    bufferSizeComparison()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 方法、select 與頻的 case

當在 `select` 語句中有多個 case 同時準備好時,Go 會隨機選擇一個執行。這種隨機性有時候可能不是我們想要的行為,特別是當我們需要優先處理某些操作時。

處理優先級的策略:
- 使用巢狀的 select
- 實現優先級佇列
- 使用不同的 channel 緩衝大小
- 週期性檢查高優先級操作
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 示範 select 的隨機性 */
func selectRandomness() {
    fmt.Println("=== Select 隨機性示範 ===")

    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)

    // 同時向兩個 channel 發送資料
    ch1 <- "來自 channel 1"
    ch2 <- "來自 channel 2"

    // 執行多次 select,觀察隨機選擇
    for i := 0; i < 10; i++ {
        // 重新填充 channels
        select {
        case msg := <-ch1:
            fmt.Printf("第 %d 次: 選擇了 ch1 - %s\n", i+1, msg)
            ch1 <- "來自 channel 1" // 重新填充
        case msg := <-ch2:
            fmt.Printf("第 %d 次: 選擇了 ch2 - %s\n", i+1, msg)
            ch2 <- "來自 channel 2" // 重新填充
        }
    }
}

func main() {
    selectRandomness()
    // 輸出會顯示隨機選擇不同的 channels
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

/* 實現優先級處理 */
func priorityHandling() {
    fmt.Println("\n=== 優先級處理示範 ===")

    highPriority := make(chan string, 2)
    lowPriority := make(chan string, 2)
    done := make(chan struct{})

    // 發送不同優先級的訊息
    go func() {
        highPriority <- "緊急任務 1"
        lowPriority <- "普通任務 1"
        time.Sleep(100 * time.Millisecond)

        lowPriority <- "普通任務 2"
        highPriority <- "緊急任務 2"
        time.Sleep(100 * time.Millisecond)

        close(highPriority)
        close(lowPriority)
    }()

    // 處理訊息時優先處理高優先級
    go func() {
        defer close(done)

        for {
            select {
            case msg, ok := <-highPriority:
                if !ok {
                    highPriority = nil // 防止重複選擇已關閉的 channel
                    break
                }
                fmt.Printf("[高優先級] 處理: %s\n", msg)
            default:
                // 只有在沒有高優先級訊息時才處理低優先級
                select {
                case msg, ok := <-lowPriority:
                    if !ok {
                        lowPriority = nil
                        break
                    }
                    fmt.Printf("[低優先級] 處理: %s\n", msg)
                default:
                    // 兩個 channel 都沒有訊息
                    if highPriority == nil && lowPriority == nil {
                        return
                    }
                    time.Sleep(10 * time.Millisecond) // 短暫等待
                }
            }
        }
    }()

    <-done
    fmt.Println("優先級處理完成")
}

func main() {
    priorityHandling()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 如何處理字串式溝通

在並行程式設計中,有時需要進行複雜的訊息交換,而不僅僅是簡單的資料傳遞。這種情況下,我們可能需要實現請求-回應模式、多方通訊或者訊息路由等複雜模式。

常見的字串式溝通模式:
- 請求-回應模式
- 發布-訂閱模式
- 訊息路由
- 工作分發與結果收集
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "time"
)

/* 請求-回應模式 */
type Request struct {
    ID       int
    Data     string
    Response chan string // 每個請求帶有自己的回應 channel
}

func requestResponsePattern() {
    fmt.Println("=== 請求-回應模式 ===")

    requests := make(chan Request, 3)

    // 啟動服務器
    go server(requests)

    // 發送多個請求
    for i := 1; i <= 3; i++ {
        response := make(chan string, 1)

        req := Request{
            ID:       i,
            Data:     fmt.Sprintf("資料 #%d", i),
            Response: response,
        }

        fmt.Printf("客戶端: 發送請求 #%d\n", i)
        requests <- req

        // 等待回應(可以同時處理多個請求)
        go func(reqID int, respCh <-chan string) {
            resp := <-respCh
            fmt.Printf("客戶端: 收到請求 #%d 的回應: %s\n", reqID, resp)
        }(i, response)
    }

    // 等待處理完成
    time.Sleep(1 * time.Second)
    close(requests)

    fmt.Println("請求-回應模式完成\n")
}

func server(requests <-chan Request) {
    for req := range requests {
        fmt.Printf("服務器: 處理請求 #%d - %s\n", req.ID, req.Data)

        // 模擬處理時間
        time.Sleep(200 * time.Millisecond)

        // 發送回應
        response := fmt.Sprintf("已處理 %s", req.Data)
        req.Response <- response
        close(req.Response)
    }
}

func main() {
    requestResponsePattern()
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

/* 發布-訂閱模式 */
type PubSub struct {
    subscribers map[string][]chan string
    mutex       sync.RWMutex
}

func NewPubSub() *PubSub {
    return &PubSub{
        subscribers: make(map[string][]chan string),
    }
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
    ps.mutex.Lock()
    defer ps.mutex.Unlock()

    ch := make(chan string, 5)
    ps.subscribers[topic] = append(ps.subscribers[topic], ch)
    return ch
}

func (ps *PubSub) Publish(topic, message string) {
    ps.mutex.RLock()
    defer ps.mutex.RUnlock()

    if subscribers, ok := ps.subscribers[topic]; ok {
        for _, ch := range subscribers {
            select {
            case ch <- message:
            default:
                fmt.Printf("警告: 訂閱者佇列已滿,丟棄訊息\n")
            }
        }
    }
}

func pubSubPattern() {
    fmt.Println("=== 發布-訂閱模式 ===")

    pubsub := NewPubSub()

    // 建立訂閱者
    news := pubsub.Subscribe("news")
    sports := pubsub.Subscribe("sports")
    general := pubsub.Subscribe("general")

    var wg sync.WaitGroup

    // 啟動訂閱者
    wg.Add(3)
    go subscriber("新聞訂閱者", news, &wg)
    go subscriber("體育訂閱者", sports, &wg)
    go subscriber("一般訂閱者", general, &wg)

    // 發布訊息
    time.Sleep(100 * time.Millisecond)

    fmt.Println("發布者: 開始發布訊息...")
    pubsub.Publish("news", "重要新聞: Go 1.21 發布!")
    pubsub.Publish("sports", "體育新聞: 世界盃開始了!")
    pubsub.Publish("general", "一般訊息: 天氣不錯")
    pubsub.Publish("news", "科技新聞: AI 技術突破")

    time.Sleep(500 * time.Millisecond)

    // 關閉 channels 以停止訂閱者
    close(news)
    close(sports)
    close(general)

    wg.Wait()
    fmt.Println("發布-訂閱模式完成\n")
}

func subscriber(name string, ch <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    for message := range ch {
        fmt.Printf("%s: 收到訊息 - %s\n", name, message)
        time.Sleep(100 * time.Millisecond) // 模擬處理時間
    }

    fmt.Printf("%s: 停止訂閱\n", name)
}

func main() {
    pubSubPattern()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 使用 WaitGroup

`sync.WaitGroup` 是 Go 中用於等待多個 goroutines 完成的同步原語。它提供了一種簡單而有效的方法來協調多個併發操作的完成。

WaitGroup 的使用方法:
- `Add(n)`: 增加等待的 goroutine 計數
- `Done()`: 標記一個 goroutine 完成
- `Wait()`: 阻塞直到所有 goroutines 完成

常見用法:
- 等待所有工作者完成
- 並行處理資料集
- 同步多個初始化步驟
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

/* WaitGroup 基本用法 */
func basicWaitGroup() {
    fmt.Println("=== WaitGroup 基本用法 ===")

    var wg sync.WaitGroup

    tasks := []string{"任務A", "任務B", "任務C", "任務D"}

    // 設定要等待的 goroutine 數量
    wg.Add(len(tasks))

    for i, task := range tasks {
        go func(id int, taskName string) {
            defer wg.Done() // 確保在函式結束時呼叫 Done()

            fmt.Printf("開始執行 %s (ID: %d)\n", taskName, id)

            // 模擬不同的執行時間
            duration := time.Duration(200+id*100) * time.Millisecond
            time.Sleep(duration)

            fmt.Printf("完成 %s (ID: %d)\n", taskName, id)
        }(i, task)
    }

    fmt.Println("等待所有任務完成...")
    wg.Wait() // 等待所有 goroutines 完成

    fmt.Println("所有任務完成!\n")
}

func main() {
    basicWaitGroup()
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

/* 並行處理資料集合 */
func parallelDataProcessing() {
    fmt.Println("=== 並行處理資料集合 ===")

    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    results := make([]int, len(data))
    var wg sync.WaitGroup

    fmt.Printf("原始資料: %v\n", data)

    // 並行處理每個元素
    for i, value := range data {
        wg.Add(1)
        go func(index, val int) {
            defer wg.Done()

            // 模擬複雜的計算(計算平方)
            time.Sleep(time.Duration(val*50) * time.Millisecond)
            result := val * val

            results[index] = result
            fmt.Printf("處理完成: %d² = %d\n", val, result)
        }(i, value)
    }

    // 等待所有處理完成
    wg.Wait()

    fmt.Printf("處理結果: %v\n", results)
    fmt.Println("並行處理完成\n")
}

func main() {
    parallelDataProcessing()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 只執行程式一次

`sync.Once` 確保某個操作只執行一次,即使在多個 goroutines 中同時呼叫也是如此。這對於初始化操作、單例模式或昂貴的一次性設定特別有用。

Once 的特點:
- 執行緒安全
- 只執行第一次呼叫
- 後續呼叫會被忽略
- 常用於延遲初始化
- 適合單例模式實現
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

/* sync.Once 基本用法 */
var (
    once     sync.Once
    instance string
)

func expensiveOperation() {
    fmt.Println("執行昂貴的初始化操作...")
    time.Sleep(500 * time.Millisecond) // 模擬昂貴的操作
    instance = "初始化完成的實例"
    fmt.Println("初始化操作完成")
}

func getInstance() string {
    once.Do(expensiveOperation) // 只會執行一次
    return instance
}

func syncOnceDemo() {
    fmt.Println("=== sync.Once 示範 ===")

    var wg sync.WaitGroup

    // 啟動多個 goroutines 同時呼叫 getInstance
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            fmt.Printf("Goroutine %d: 嘗試獲取實例...\n", id)
            result := getInstance()
            fmt.Printf("Goroutine %d: 獲得實例: %s\n", id, result)
        }(i)
    }

    wg.Wait()
    fmt.Println("所有 goroutines 完成\n")
}

func main() {
    syncOnceDemo()
    // 輸出會顯示初始化只執行了一次
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

/* 實際應用:配置載入 */
type Config struct {
    DatabaseURL string
    APIKey      string
    MaxRetries  int
}

var (
    config     *Config
    configOnce sync.Once
)

func loadConfig() {
    fmt.Println("開始載入配置文件...")

    // 模擬從檔案或環境變數載入配置
    time.Sleep(300 * time.Millisecond)

    config = &Config{
        DatabaseURL: "postgresql://localhost:5432/mydb",
        APIKey:      "secret-api-key-12345",
        MaxRetries:  3,
    }

    fmt.Println("配置載入完成")
}

func GetConfig() *Config {
    configOnce.Do(loadConfig)
    return config
}

func configLoadingDemo() {
    fmt.Println("=== 配置載入示範 ===")

    var wg sync.WaitGroup

    // 模擬多個服務同時需要配置
    services := []string{"Web服務", "API服務", "資料庫服務", "快取服務"}

    for _, service := range services {
        wg.Add(1)
        go func(serviceName string) {
            defer wg.Done()

            fmt.Printf("%s: 請求配置\n", serviceName)
            cfg := GetConfig()

            fmt.Printf("%s: 獲得配置 - DB: %s\n", serviceName, cfg.DatabaseURL)
        }(service)
    }

    wg.Wait()
    fmt.Println("所有服務已獲得配置\n")
}

func main() {
    configLoadingDemo()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 將目前的工具數一台記起來

這個概念指的是記錄和管理當前運行的 goroutines 數量,以及監控系統資源的使用情況。這對於除錯、性能調優和資源管理非常重要。

監控要點:
- Goroutine 數量追蹤
- 記憶體使用情況
- Channel 狀態監控
- 性能指標收集
- 資源洩漏檢測
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

/* Goroutine 監控 */
func monitorGoroutines() {
    fmt.Println("=== Goroutine 監控 ===")

    // 啟動監控 goroutine
    done := make(chan struct{})
    go func() {
        ticker := time.NewTicker(200 * time.Millisecond)
        defer ticker.Stop()

        for {
            select {
            case <-done:
                return
            case <-ticker.C:
                count := runtime.NumGoroutine()
                fmt.Printf("[監控] 當前 goroutine 數量: %d\n", count)
            }
        }
    }()

    var wg sync.WaitGroup

    // 分批啟動 goroutines
    for batch := 1; batch <= 3; batch++ {
        fmt.Printf("\n--- 啟動第 %d 批 goroutines ---\n", batch)

        for i := 0; i < 3; i++ {
            wg.Add(1)
            go func(batchNum, workerNum int) {
                defer wg.Done()

                fmt.Printf("批次 %d 工作者 %d 開始\n", batchNum, workerNum)
                time.Sleep(time.Duration(500+workerNum*100) * time.Millisecond)
                fmt.Printf("批次 %d 工作者 %d 完成\n", batchNum, workerNum)
            }(batch, i)
        }

        time.Sleep(300 * time.Millisecond) // 讓監控顯示變化
    }

    wg.Wait()
    close(done) // 停止監控

    fmt.Printf("\n最終 goroutine 數量: %d\n", runtime.NumGoroutine())
    fmt.Println("監控完成\n")
}

func main() {
    monitorGoroutines()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

/* 記憶體使用監控 */
func memoryMonitoring() {
    fmt.Println("=== 記憶體使用監控 ===")

    // 記錄初始狀態
    printMemStats("初始狀態")

    // 建立一些資料結構
    var data [][]int

    for i := 0; i < 5; i++ {
        // 分配記憶體
        batch := make([]int, 100000) // 分配 100k 個整數
        for j := range batch {
            batch[j] = j
        }
        data = append(data, batch)

        printMemStats(fmt.Sprintf("分配第 %d 批資料後", i+1))
        time.Sleep(100 * time.Millisecond)
    }

    // 強制垃圾回收
    fmt.Println("\n執行垃圾回收...")
    runtime.GC()
    runtime.GC() // 執行兩次確保完全清理

    printMemStats("垃圾回收後")

    // 保持資料引用避免被回收
    fmt.Printf("資料長度: %d (保持引用)\n", len(data))

    fmt.Println("記憶體監控完成\n")
}

func printMemStats(label string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    fmt.Printf("[%s]\n", label)
    fmt.Printf("  分配的記憶體: %.2f MB\n", float64(m.Alloc)/1024/1024)
    fmt.Printf("  總分配記憶體: %.2f MB\n", float64(m.TotalAlloc)/1024/1024)
    fmt.Printf("  系統記憶體: %.2f MB\n", float64(m.Sys)/1024/1024)
    fmt.Printf("  GC 次數: %d\n\n", m.NumGC)
}

func main() {
    memoryMonitoring()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 何時該用 mutex 來取代 channel

雖然 Go 鼓勵「通過通訊來共享記憶體」,但有些情況下使用 `sync.Mutex` 可能更適合。選擇 mutex 而不是 channel 的情況包括:

**使用 Mutex 的場景:**
- 保護共享狀態的簡單讀寫操作
- 性能要求極高的場景
- 需要多個 goroutines 同時讀取的情況(RWMutex)
- 簡單的計數器或狀態標記

**使用 Channel 的場景:**
- 需要傳遞資料或信號
- 實現複雜的協調模式
- 生產者-消費者模式
- 需要選擇性等待多個操作
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

/* Mutex vs Channel 比較:簡單計數器 */

// 使用 Mutex 的計數器
type MutexCounter struct {
    mu    sync.Mutex
    count int
}

func (c *MutexCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *MutexCounter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// 使用 Channel 的計數器
type ChannelCounter struct {
    ch chan int
}

func NewChannelCounter() *ChannelCounter {
    cc := &ChannelCounter{
        ch: make(chan int),
    }

    go func() {
        count := 0
        for {
            select {
            case <-cc.ch: // 收到增加信號
                count++
            }
        }
    }()

    return cc
}

func (c *ChannelCounter) Increment() {
    c.ch <- 1
}

func mutexVsChannelDemo() {
    fmt.Println("=== Mutex vs Channel 性能比較 ===")

    const numOperations = 100000
    const numGoroutines = 10

    // 測試 Mutex 計數器
    fmt.Println("測試 Mutex 計數器...")
    mutexCounter := &MutexCounter{}

    start := time.Now()
    var wg sync.WaitGroup

    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < numOperations/numGoroutines; j++ {
                mutexCounter.Increment()
            }
        }()
    }

    wg.Wait()
    mutexDuration := time.Since(start)

    fmt.Printf("Mutex 結果: 計數 = %d, 時間 = %v\n", mutexCounter.Get(), mutexDuration)

    // 測試 Channel 計數器
    fmt.Println("\n測試 Channel 計數器...")
    channelCounter := NewChannelCounter()

    start = time.Now()

    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < numOperations/numGoroutines; j++ {
                channelCounter.Increment()
            }
        }()
    }

    wg.Wait()
    channelDuration := time.Since(start)

    fmt.Printf("Channel 時間 = %v\n", channelDuration)
    fmt.Printf("\n性能比較: Mutex 比 Channel 快 %.2f 倍\n\n",
        float64(channelDuration)/float64(mutexDuration))
}

func main() {
    mutexVsChannelDemo()
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

/* RWMutex 用於讀多寫少的場景 */
type Cache struct {
    rwmu sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) (string, bool) {
    c.rwmu.RLock() // 讀鎖
    defer c.rwmu.RUnlock()

    value, ok := c.data[key]
    return value, ok
}

func (c *Cache) Set(key, value string) {
    c.rwmu.Lock() // 寫鎖
    defer c.rwmu.Unlock()

    c.data[key] = value
}

func rwMutexDemo() {
    fmt.Println("=== RWMutex 示範 ===")

    cache := NewCache()
    var wg sync.WaitGroup

    // 初始化一些資料
    cache.Set("user:1", "Alice")
    cache.Set("user:2", "Bob")
    cache.Set("user:3", "Charlie")

    // 啟動多個讀取者
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(readerID int) {
            defer wg.Done()

            for j := 0; j < 10; j++ {
                key := fmt.Sprintf("user:%d", (j%3)+1)
                if value, ok := cache.Get(key); ok {
                    fmt.Printf("讀取者 %d: %s = %s\n", readerID, key, value)
                }
                time.Sleep(50 * time.Millisecond)
            }
        }(i)
    }

    // 啟動一個寫入者
    wg.Add(1)
    go func() {
        defer wg.Done()

        for i := 4; i <= 6; i++ {
            key := fmt.Sprintf("user:%d", i)
            value := fmt.Sprintf("User%d", i)

            cache.Set(key, value)
            fmt.Printf("寫入者: 設定 %s = %s\n", key, value)

            time.Sleep(200 * time.Millisecond)
        }
    }()

    wg.Wait()
    fmt.Println("RWMutex 示範完成\n")
}

func main() {
    rwMutexDemo()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 原子運算一你應該要它們

`sync/atomic` 套件提供了原子操作,這是最快的同步機制。原子操作在硬體層面保證操作的原子性,不會被其他 goroutines 中斷。

原子操作適用於:
- 簡單的數值操作(增減、交換)
- 高性能要求的場景
- 無鎖程式設計
- 狀態標記和計數器

常用的原子操作:
- `atomic.AddInt64()`: 原子加法
- `atomic.LoadInt64()`: 原子讀取
- `atomic.StoreInt64()`: 原子寫入
- `atomic.SwapInt64()`: 原子交換
- `atomic.CompareAndSwapInt64()`: 比較並交換
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

/* 原子操作基本用法 */
func atomicOperationsDemo() {
    fmt.Println("=== 原子操作示範 ===")

    var counter int64
    var wg sync.WaitGroup

    numGoroutines := 10
    incrementsPerGoroutine := 10000

    start := time.Now()

    // 啟動多個 goroutines 進行原子增加操作
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            for j := 0; j < incrementsPerGoroutine; j++ {
                // 原子增加操作
                atomic.AddInt64(&counter, 1)
            }

            // 原子讀取操作
            current := atomic.LoadInt64(&counter)
            fmt.Printf("Goroutine %d 完成,當前計數: %d\n", id, current)
        }(i)
    }

    wg.Wait()
    duration := time.Since(start)

    finalCount := atomic.LoadInt64(&counter)
    expected := int64(numGoroutines * incrementsPerGoroutine)

    fmt.Printf("\n最終計數: %d\n", finalCount)
    fmt.Printf("預期計數: %d\n", expected)
    fmt.Printf("結果正確: %t\n", finalCount == expected)
    fmt.Printf("執行時間: %v\n\n", duration)
}

func main() {
    atomicOperationsDemo()
}
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

/* 比較不同同步機制的性能 */
func performanceComparison() {
    fmt.Println("=== 同步機制性能比較 ===")

    const numOperations = 1000000
    const numGoroutines = 10

    // 測試原子操作
    fmt.Println("測試原子操作...")
    var atomicCounter int64
    start := time.Now()

    var wg sync.WaitGroup
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < numOperations/numGoroutines; j++ {
                atomic.AddInt64(&atomicCounter, 1)
            }
        }()
    }
    wg.Wait()
    atomicDuration := time.Since(start)

    fmt.Printf("原子操作結果: 計數 = %d, 時間 = %v\n",
        atomic.LoadInt64(&atomicCounter), atomicDuration)

    // 測試 Mutex
    fmt.Println("\n測試 Mutex...")
    var mutexCounter int64
    var mu sync.Mutex
    start = time.Now()

    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < numOperations/numGoroutines; j++ {
                mu.Lock()
                mutexCounter++
                mu.Unlock()
            }
        }()
    }
    wg.Wait()
    mutexDuration := time.Since(start)

    fmt.Printf("Mutex 結果: 計數 = %d, 時間 = %v\n", mutexCounter, mutexDuration)

    // 性能比較
    fmt.Printf("\n性能比較:原子操作比 Mutex 快 %.2f 倍\n",
        float64(mutexDuration)/float64(atomicDuration))
}

func main() {
    performanceComparison()
}
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 進一步學習並行的地方

Go 的並行程式設計是一個深度的主題,本章介紹了基礎概念和常用模式。要進一步提升並行程式設計技能,可以關注以下方向:

**進階主題:**
- Context 套件的深入應用
- 並行模式設計(Pipeline、Fan-in/Fan-out、Worker Pool)
- 錯誤處理和恢復策略
- 性能調優和除錯技巧

**學習資源:**
- Go Blog 的並行相關文章
- "Go Concurrency Patterns" 系列講座
- 開源專案中的並行程式碼實例
- 效能分析工具(pprof)的使用

**實踐建議:**
- 從簡單的並行程式開始
- 多寫測試來驗證並行程式的正確性
- 學習使用 race detector
- 關注記憶體管理和 GC 影響
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
## 總結

Go 的並行程式設計提供了強大而優雅的工具來處理複雜的並行需求。通過本章的學習,我們掌握了:

### 核心概念:
1. **Goroutines**:輕量級的執行緒,Go 並行的基石
2. **Channels**:安全的通訊機制,實現「通過通訊來共享記憶體」
3. **Select**:多路選擇,實現複雜的協調邏輯

### 實踐模式:
4. **生命週期管理**:使用 done channels、context 來控制 goroutines
5. **同步機制**:WaitGroup、Once、Mutex 等工具的合理使用
6. **性能考量**:原子操作、合適的同步機制選擇

### 最佳實踐:
- 避免 goroutine 洩漏,總是確保 goroutines 能夠退出
- 正確處理迴圈變數與 goroutines 的關係
- 在適當的場景選擇適當的同步機制
- 重視程式的可測試性和可維護性

### 設計哲學:
Go 的並行模型鼓勵編寫簡潔、可理解的並行程式。透過 goroutines 和 channels 的組合,我們可以構建出既高效又易於理解的並行系統。記住 Go 的格言:「不要通過共享記憶體來通訊,而要通過通訊來共享記憶體」。

掌握這些概念和模式,你就能夠編寫出安全、高效的 Go 並行程式。
<!-- END_MARKDOWN_CELL -->
