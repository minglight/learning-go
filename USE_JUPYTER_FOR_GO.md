# 使用 Jupyter Notebook 學習 Go 語言 - 完整教學

## 概述
本教學將教你如何使用 Jupyter Notebook 來學習 Go 語言，透過 GoNB (Go Notebooks) 實現類似 Python Jupyter 的互動式編程體驗。

## 前置需求
- 已安裝 Go 語言 (建議 1.21+)
- 已安裝 Python 和 Jupyter Lab/Notebook
- 終端機或命令提示字元

## 步驟 1: 安裝 GoNB

### 1.1 安裝 GoNB
```bash
go install github.com/janpfeifer/gonb@latest
```

### 1.2 驗證安裝
```bash
gonb --version
```

## 步驟 2: 安裝 GoNB Kernel 到 Jupyter

### 2.1 安裝 Kernel
```bash
gonb --install
```

成功訊息會顯示：
```
Go (gonb) kernel configuration installed in "/Users/username/Library/Jupyter/kernels/gonb/kernel.json".
```

### 2.2 驗證 Kernel 安裝
```bash
jupyter kernelspec list
```

應該會看到：
```
Available kernels:
  gonb             /Users/username/Library/Jupyter/kernels/gonb
  python3          /path/to/python3/kernel
```

## 步驟 3: 啟動 Jupyter Lab

### 3.1 啟動服務
```bash
jupyter lab
```

### 3.2 開啟瀏覽器
- 自動開啟瀏覽器或手動前往：`http://localhost:8888`
- 複製並貼上終端機顯示的 token URL

## 步驟 4: 建立 Go Notebook

### 4.1 建立新 Notebook
1. 在 Jupyter Lab 中點擊右上角的 **"New"** 按鈕
2. 選擇 **"Go (gonb)"** kernel
3. 這會建立一個新的 `.ipynb` 檔案

### 4.2 基本操作
- **新增程式碼單元格**：點擊 **"+"** 按鈕 → 選擇 **"Code"**
- **新增 Markdown 單元格**：點擊 **"+"** 按鈕 → 選擇 **"Markdown"**
- **執行程式碼**：按 `Shift + Enter` 或點擊播放按鈕 ▶️
- **儲存**：按 `Ctrl+S` 或點擊儲存按鈕

## 步驟 5: 在 Cursor 中使用

### 5.1 安裝 Jupyter 擴展
1. 在 Cursor 中按 `Cmd+Shift+X` 打開擴展面板
2. 搜尋並安裝 **"Jupyter"** 擴展 (Microsoft 官方)

### 5.2 建立 .ipynb 檔案
```json
{
 "cells": [
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "# Go 學習筆記\n",
    "\n",
    "這是一個使用 GoNB kernel 的 Jupyter notebook。"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": null,
   "metadata": {},
   "outputs": [],
   "source": [
    "package main\n",
    "\n",
    "import \"fmt\"\n",
    "\n",
    "func main() {\n",
    "    fmt.Println(\"Hello from Go in Jupyter!\")\n",
    "}"
   ]
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Go (gonb)",
   "language": "go",
   "name": "gonb"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 4
}
```

### 5.3 在 Cursor 中執行
1. 打開 `.ipynb` 檔案
2. 選擇 **"Go (gonb)"** kernel
3. 點擊播放按鈕或按 `Shift + Enter` 執行程式碼

## 步驟 6: 實用快捷鍵

### 6.1 單元格操作
- **`A`**：在當前單元格上方新增程式碼單元格
- **`B`**：在當前單元格下方新增程式碼單元格
- **`M`**：將當前單元格轉換為 Markdown
- **`Y`**：將當前單元格轉換為程式碼
- **`DD`**：刪除當前單元格

### 6.2 執行操作
- **`Shift + Enter`**：執行當前單元格並移動到下一個
- **`Ctrl + Enter`**：執行當前單元格但不移動
- **`Alt + Enter`**：執行當前單元格並在下方新增新單元格

## 步驟 7: 範例程式碼

### 7.1 基本語法
```go
package main

import "fmt"

func main() {
    // 變數宣告
    name := "Go"
    version := 1.21
    fmt.Printf("Language: %s, Version: %.1f\n", name, version)
    
    // 陣列和切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Numbers:", numbers)
    
    // 迴圈
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("Sum: %d\n", sum)
}
```

### 7.2 函數定義
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    result := calculateSquare(5)
    fmt.Printf("Square of 5: %.0f\n", result)
}

func calculateSquare(x float64) float64 {
    return math.Pow(x, 2)
}
```

## 步驟 8: 故障排除

### 8.1 常見問題

**問題：`gonb --kernel` 需要參數**
- 解決：這個命令是給 Jupyter 內部使用的，不需要手動執行

**問題：找不到 GoNB kernel**
- 解決：重新執行 `gonb --install`

**問題：程式碼執行失敗**
- 解決：確保每個程式碼單元格都有 `package main` 和 `func main()`

### 8.2 重新安裝
如果遇到問題，可以重新安裝：
```bash
# 移除舊的 kernel
jupyter kernelspec remove gonb

# 重新安裝
gonb --install
```

## 優點

1. **互動式學習**：可以即時執行和測試程式碼
2. **筆記整合**：Markdown 和程式碼混合使用
3. **版本控制**：.ipynb 檔案可以加入 Git
4. **多環境支援**：瀏覽器和 IDE 都可以使用
5. **即時反饋**：程式碼錯誤會立即顯示

## 總結

使用 Jupyter Notebook 學習 Go 語言提供了：
- 更好的學習體驗
- 即時程式碼測試
- 結構化的筆記系統
- 與傳統 IDE 的無縫整合

現在你可以開始使用 Jupyter Notebook 來學習 Go 語言了！








