# Markdown to Jupyter Notebook Converter

這是一個用 Go 編寫的轉換工具，將特殊格式的 Markdown 文件轉換成 Jupyter Notebook (.ipynb) 格式。

## 🎯 用途

當撰寫大型 Jupyter Notebook 內容時，直接生成 JSON 格式會遇到 token 限制問題。使用此轉換器可以：
1. 先用 Markdown 格式撰寫內容（省 60-70% token）
2. 使用轉換器自動生成 .ipynb 文件

## 📋 Markdown 格式規範

### 基本結構

使用 HTML 註解標記來區分不同類型的 cells：

```markdown
<!-- MARKDOWN_CELL -->
# 章節標題

這裡是 Markdown 內容，可以包含：
- 列表
- **粗體** 和 *斜體*
- 程式碼片段 `inline code`
- 連結、圖片等所有標準 Markdown 語法
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 程式範例說明 */
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
// 輸出: Hello, World!
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 下一個小節

繼續內容...
<!-- END_MARKDOWN_CELL -->
```

### 重要規則

1. **Markdown Cell 標記**
   - 開始標記：`<!-- MARKDOWN_CELL -->`
   - 結束標記：`<!-- END_MARKDOWN_CELL -->`
   - **必須**有結束標記

2. **Code Cell 標記**
   - 開始標記：`<!-- CODE_CELL -->`
   - 結束標記：`<!-- END_CODE_CELL -->`
   - Code fence 使用：` ```go ` 和 ` ``` `
   - 轉換時會自動移除 code fence 標記
   - **必須**有結束標記

3. **一致性規則**
   - 所有 cells 都必須有明確的開始和結束標記
   - 沒有標記的內容會被忽略
   - 空白或只有空格的 cells 會被自動忽略

## 🚀 使用方式

### 編譯（已完成）

轉換器已經編譯好，位於：
```bash
/Users/hank/Workspace/hank/learning-go/converter/md2ipynb
```

### 基本用法

```bash
# 語法
./converter/md2ipynb input.md output.ipynb

# 範例：轉換 ch9 筆記
./converter/md2ipynb ch9/ch9_modules_source.md ch9/ch9_modules_packages_imports.ipynb

# 範例：轉換 ch10 筆記
./converter/md2ipynb ch10/ch10_concurrency_source.md ch10/ch10_concurrency.ipynb
```

### 輸出說明

成功轉換後會顯示：
```
📊 共 N 個 cells
✅ 成功轉換: input.md -> output.ipynb
```

## 📝 AI 撰寫 Notebook 的工作流程

### Step 1: 撰寫 Markdown 源文件

使用 `Write` 或 `Edit` 工具撰寫 Markdown 格式的源文件：

```
chN/
├── chN_topic_source.md        # Markdown 源文件
└── chN_topic.ipynb             # 轉換後的 notebook（尚未生成）
```

**撰寫策略：**
- 由於是純文字 Markdown，可以分段撰寫
- 每次 `Write` 或 `Edit` 時 token 消耗較少
- 如果內容過多，可以分 Part 1, Part 2 寫入不同的 `.md` 文件

**範例：**
```markdown
1. Write → ch9/ch9_modules_source.md (撰寫前半部分)
2. Edit → ch9/ch9_modules_source.md (追加後半部分)
```

### Step 2: 執行轉換

使用 `Bash` 工具執行轉換：

```bash
cd /Users/hank/Workspace/hank/learning-go
./converter/md2ipynb chN/chN_topic_source.md chN/chN_topic.ipynb
```

### Step 3: 驗證結果

1. **檢查 JSON 格式**
```bash
python3 -m json.tool chN/chN_topic.ipynb > /dev/null && echo "✅ JSON 正確"
```

2. **查看 Notebook 資訊**
```bash
python3 -c "
import json
with open('chN/chN_topic.ipynb', 'r') as f:
    nb = json.load(f)
    print(f'Cells: {len(nb[\"cells\"])}')
    for i, cell in enumerate(nb['cells']):
        print(f'{i+1}. [{cell[\"cell_type\"]}] {cell[\"id\"]}')
"
```

3. **預覽前幾個 cells**
```bash
Read → chN/chN_topic.ipynb (limit: 50)
```

### Step 4: 清理（可選）

如果不需要保留 Markdown 源文件：
```bash
rm chN/chN_topic_source.md
```

## 🔧 技術細節

### 生成的 Notebook 結構

```json
{
  "cells": [
    {
      "cell_type": "markdown",
      "id": "cell-0",
      "metadata": {},
      "source": ["# 標題\n", "\n", "內容..."]
    },
    {
      "cell_type": "code",
      "id": "cell-1",
      "metadata": {},
      "source": ["package main\n", "func main() {}"],
      "execution_count": null,
      "outputs": []
    }
  ],
  "metadata": {
    "kernelspec": {
      "display_name": "Go",
      "language": "go",
      "name": "gophernotes"
    },
    "language_info": {
      "file_extension": ".go",
      "mimetype": "text/x-go",
      "name": "go"
    }
  },
  "nbformat": 4,
  "nbformat_minor": 4
}
```

### Cell ID 命名

- 自動生成格式：`cell-0`, `cell-1`, `cell-2`, ...
- 按照出現順序遞增
- 可以在生成後手動修改為更有意義的 ID

## ✅ 測試覆蓋率

轉換器包含完整的測試：

**Unit Tests (14 個):**
- Notebook 結構測試
- Parser 解析測試
- 字串處理測試

**Integration Tests (6 個):**
- 簡單文件轉換
- 複雜文件轉換
- 邊界情況處理
- 錯誤處理

執行測試：
```bash
cd converter
go test -v
```

## 🎯 範例

請參考 [example.md](example.md) 和生成的 [example.ipynb](example.ipynb)。

## 📊 優勢總結

| 特性 | 說明 |
|------|------|
| **Token 效率** | Markdown 比 JSON 省 60-70% token |
| **易於撰寫** | 熟悉的 Markdown 語法 |
| **可維護性** | 人類可讀的源文件 |
| **可重複使用** | 適用於所有章節的 notebook |
| **類型安全** | Go 的強類型系統 |
| **完整測試** | 20 個測試全覆蓋 |

## ⚠️ 注意事項

1. **Code fence 必須使用 ` ```go `**
   - 不要使用其他語言標記
   - 轉換器會自動移除 fence 標記

2. **標記必須獨立一行**
   ```markdown
   <!-- MARKDOWN_CELL -->    ✅ 正確
   文字 <!-- MARKDOWN_CELL --> ❌ 錯誤
   ```

3. **所有 cells 都必須有結束標記**
   ```markdown
   <!-- MARKDOWN_CELL -->
   markdown content
   <!-- END_MARKDOWN_CELL -->    ✅ 必須有

   <!-- CODE_CELL -->
   ```go
   code here
   ```
   <!-- END_CODE_CELL -->        ✅ 必須有
   ```

4. **空行會被保留**
   - Markdown 中的空行會出現在 notebook 中
   - Code 中的空行也會保留

## 🔗 相關文件

- [CLAUDE.md](../CLAUDE.md) - 專案整體指南
- [目錄.md](../目錄.md) - 學習 Go 書籍目錄
