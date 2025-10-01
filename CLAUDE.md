# learning-go

這是一個學習Go的Project, 主要有三個任務, 當我要求時你要根據我的要求做不同的事情
- 產生jupyter筆記
- 產生面試考題
- 改考卷

**重要** : 不需要把修改檔案的diff印在console

### 參考資訊 
- 要觀看目錄 ( ./目錄.md)
- 要看筆記 example (./ch2/ch2_basic_types_declarations.ipynb)
- 要看考題 example (./ch2/ch2_interview_questions.ipynb)
- 要看考題答案example (./ch2/ch2_interview_answers_review.ipynb)

### Terms & Condition
- `章` or `Chapter` : 依照目錄, 這本筆記總共有15個chapter
- `節` : 依照目錄, 每一章的內部第一排縮排為節
- `小節` : 依照目錄, 每一章的內部第二排縮排為節
- `Section` : 由一個markdown block & n個code block組成, 如果目錄有小節的話, 每一個小節都需要自己一個section. 
- `Markdown block` : 說明每一章,節, 小節要陳述的內容, 每個陳述都要有code or signature example
- `Code block` : 對應到這個section的markdown說明的code example, 每個example都要使用comment說明, 如果有Print出東西的話, 也要直接在comment描述這行會印出什麼結果, 我不想還需要執行才看的到結果.
- `Code block重點` 每一個demo的use case要獨立使用一個code block, 不要這個section的所有example都放到同一個blcok裡面, 弄的很長最後才在main一次呼叫這樣子太難看的懂. 
- 請只用 NotebookEdit/Write 建檔，不要將內容貼在訊息中


## 三個任務
### 1. 產生/補充/整理jupyter筆記 : 
- 依照那一章`目錄`的順序產生jupyter學習筆記; 每個markdown block的標題要跟目錄寫的一模一樣.
- 如果`目錄`的起始跟結束頁面差超過20頁, 那麼你需要先規劃, 依照頁數跟`節` (注意是節而非小節)來拆成不同的task, 分成不同的檔案(使用part1, part2...), 但不要切在小節上. 舉例來說, 如果這一章有25頁, 你就找一個`節`的結尾是13. 那就拆成 1-12, 13-25; 另一個例子, 如果只有15頁, 就不需要分part. 
- 除了每一章需要說明之外，每一(小)節也都至少要有從初階到進階的詳細說明(包含example) in markdown block, 然後接著code blocks做詳細的範例, 每一段展示code都要有comment說明; 
- 寫code block要注意分段, 不要一個example寫完所有概念markdown提到的所有概念, 而是每個概念都分開來寫example, 然後使用每個example都要使用 /*{說明}*/ 開頭講一下現在要幹嘛, 這樣我才不需要一直scroll.
- 如果有特別注意的點, 或是容易出錯的點也都要說明在markdown的部分跟 code block的部分
- 你可以把每一(小)節拆成一個task來做
- 不要寫這個chapter更進階的概念, 才會提到的困難的概念, 例如ch7筆記提到ch10才有的概念是不需要的.
- 在生成(修改)完的最後, 要加一個task去確認完整的順序是不是跟目錄一模一樣, 不然就要調整

### 2. 產生面試考題
在同一個chapter folder下產生.ipynb file, 要分兩個部分
- `練習` : 簡單概念的練習, 可以在code block的comment直接寫題目, 然後讓我可以打字肌肉練習這個章節的概念. 
- `題目` : 要寫一些面試常出現的考題, 你使用markdown寫題目, 題目要有標題跟內容, 然後要留 code block讓我寫答案.
- 如果這個chapter的筆記有多個part, 那你要針對每個part產生不同的檔案. 

### 3. 改考卷
當我寫完之後我會請你改考卷, 你要另外開一個jupyter file,  一開始要先寫summary, 說明哪裡寫錯或寫得不好,最後再一題一題講解copy我寫不好或寫錯的那一題的題目跟我寫的答案, 然後新增你的說明跟正確解答


## 資料夾結構與命名規範
- **章節資料夾**: 根據 `./目錄.md` 建立 `ch1` ~ `ch14`。
- **每章標準檔案命名**:
  - **章節總筆記**: `chN_note.ipynb`, 可以加上這章在做什麼in file name. ex: `ch2_basic_types_declarations.ipynb`
  - **面試題庫**: `chN_interview_questions.ipynb`。
  - **改卷與講解**: `chN_interview_answers_review.ipynb`。
  - **Go 範例程式**: 放在同一章節資料夾底下的 `.go` 檔（例如 `pointer.go`）。若該章程式較多，可酌情新增 `examples/` 子資料夾，但預設直接放在章節根目錄。

- **寫入位置規則**:
  - 僅將與第 N 章相關的內容寫入 `./chN/`。

- **命名規則**:
  - 全部檔名使用小寫，單字以底線 `_` 分隔。
  - 一律以 `chN_` 作為章節前綴（例如 `ch2_interview_questions.ipynb`）。
  - Notebook 副檔名 `.ipynb`，Go 程式檔副檔名 `.go`。

- **參考範例**:
  - 章節 2 筆記：`./ch2/ch2_basic_types_declarations.ipynb`
  - 章節 2 題目：`./ch2/ch2_interview_questions.ipynb`
  - 章節 2 改卷：`./ch2/ch2_interview_answers_review.ipynb`

- **指示 AI 時的寫入位置**:
  - 「產生 chN 面試考題」→ 寫入 `./chN/chN_interview_questions.ipynb`。
  - 「改 chN 考卷」→ 在 `./chN/chN_interview_answers_review.ipynb` 新增一個新的改卷段落。
  - 「針對 chN產生筆記」→ 建立/更新 `./chN/chN_<chapter_subject>.ipynb`；`chapter_subject` 請以英文精簡描述主題。


# AI Instructions & Guidelines

## General Principles
- Do what has been asked; nothing more, nothing less
- NEVER create files unless they're absolutely necessary for achieving your goal
- ALWAYS prefer editing an existing file to creating a new one
- NEVER proactively create documentation files (*.md) or README files. Only create documentation files if explicitly requested by the User

## Jupyter Notebook Management

### 🛠️ Tool Selection Rules

| 任務類型 | 使用工具 | 說明 |
|----------|----------|------|
| **新建 notebook** | `Write` | 一次性創建完整結構 |
| **修改單一 cell** | `NotebookEdit` | 替換或修改特定 cell |
| **修改多個 cells** | `Write` | 重新創建整個 notebook |
| **修復損壞檔案** | `Write` | 萃取內容後重建 |

### ✨ Best Practices

#### 1. **Meaningful Cell IDs 命名規範**

所有 notebook cells 都應該有語義化的 ID，便於理解和維護：

```markdown
命名模式：
- chapter-intro           # 章節介紹
- section-1-basics        # 第一節：基礎概念
- code-pointer-basic      # 指標基礎程式碼範例
- section-2-slice         # 第二節：切片
- code-slice-demo         # 切片示範程式碼
- section-3-map           # 第三節：映射
- code-map-examples       # 映射範例
- chapter-summary         # 章節總結
- practice-exercises      # 練習題
```

**ID 命名原則**：
- 使用 kebab-case (小寫 + 連字符)
- 包含內容類型前綴：`chapter-`, `section-`, `code-`, `practice-`
- 描述性名稱，一看就知道內容
- 按邏輯順序編號：`section-1`, `section-2`

#### 2. **NotebookEdit 僅用於 Replace/Edit**

**✅ 推薦用法**：
```json
// 修改現有 cell 內容
{
  "notebook_path": "/path/to/notebook.ipynb",
  "cell_id": "code-pointer-basic",
  "new_source": "/* 更新的程式碼 */\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    var x int = 42\n    var p *int = &x  // p 指向 x 的位址\n    fmt.Println(\"x =\", x)      // 印出: x = 42\n    fmt.Println(\"&x =\", &x)    // 印出: &x = 0xc000018098\n    fmt.Println(\"p =\", p)      // 印出: p = 0xc000018098\n    fmt.Println(\"*p =\", *p)    // 印出: *p = 42\n}"
}

// 替換特定章節說明
{
  "notebook_path": "/path/to/notebook.ipynb",
  "cell_id": "section-1-basics",
  "new_source": "## 1.1 指標基礎概念\n\n更新的說明內容..."
}
```

**❌ 避免的用法**：
```json
// 不要用 insert 模式
{
  "edit_mode": "insert",  // 會導致順序問題
  "cell_type": "markdown",
  "new_source": "新內容"
}
```

#### 3. **預規劃 Notebook 結構**

在創建 notebook 前，先規劃完整結構：

```markdown
規劃範例 - Chapter 3: Pointers
1. chapter-intro          : 章節介紹和學習目標
2. section-1-basics       : 指標基礎概念
3. code-basic-example     : 基礎指標範例
4. section-2-operations   : 指標操作
5. code-operations-demo   : 操作示範
6. section-3-functions    : 指標與函式
7. code-functions-example : 函式參數指標範例
8. section-4-arrays       : 指標與陣列
9. code-arrays-demo       : 陣列指標示範
10. chapter-summary       : 章節總結
```

### ⚠️ 絕對禁止的操作
1. **使用 NotebookEdit insert** - 不同task不知道彼此做了什麼, 所以會導致 cell 順序顛倒
2. **混用 Write 和 NotebookEdit** - 容易產生結構衝突
3. **手動轉義 JSON 屬性名** - `"metadata"` 不要寫成 `\"metadata\"`
4. **在損壞檔案上直接修復** - 必須重新創建

**執行流程**：
1. **規劃階段**: 列出所有 sections 和對應 cell IDs
2. **準備階段**: 準備所有 markdown 和 code 內容
3. **檢查階段**: 檢查輸出格式
  - [ ] 所有引號格式一致（無 `\"` 轉義字符）
  - [ ] JSON 結構完整（brackets 和 braces 配對）
  - [ ] metadata 部分格式與 cells 部分一致
4. **創建階段**: 使用 Write 工具一次性創建完整結構
  - 如果遇到Output Token Maximum超過的API Error, 就把筆記拆成兩份.
5. **維護階段**: 只用 NotebookEdit replace 模式做小幅修正

### 📋 正確的 Jupyter Notebook JSON 結構

```json
{
  "cells": [
    {
      "cell_type": "markdown",
      "id": "unique-cell-id",
      "metadata": {},
      "source": [
        "# 章節標題\n",
        "\n",
        "這是 markdown 內容"
      ]
    },
    {
      "cell_type": "code",
      "id": "code-cell-id",
      "metadata": {},
      "execution_count": null,
      "outputs": [],
      "source": [
        "/* Go 程式範例 */\n",
        "package main\n",
        "\n",
        "import \"fmt\"\n",
        "\n",
        "func main() {\n",
        "    fmt.Println(\"Hello, World!\")\n",
        "}"
      ]
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

### 🔥 Critical JSON Formatting Rules

**根本問題**: 混合使用轉義和非轉義引號格式會導致 JSON 解析失敗

**❌ 錯誤的混合格式**:
```json
"cells": [...],           // ✓ 正確的非轉義格式
\"metadata\": {           // ✗ 錯誤的轉義格式
  \"kernelspec\": {...}   // ✗ 造成解析錯誤
}
```

**✅ 正確的一致格式**:
```json
"cells": [...],
"metadata": {
  "kernelspec": {...}
}
```

**Prevention Rules**:
- 整個 JSON 檔案必須使用一致的引號格式
- 使用 Write tool 時檢查是否有 `\"` 轉義字符出現
- metadata 部分必須與 cells 部分使用相同格式

## Notebook 修復流程

### 🚨 診斷損壞的 Notebook

1. **識別錯誤**: JSON 解析錯誤通常出現在特定行列位置
2. **萃取內容**: 即使 JSON 損壞，內容通常還是可讀的
3. **重建策略**: 完全重新創建，不要嘗試修復

### 🔧 修復標準流程

```markdown
1. **診斷階段**
   - 嘗試 Read 檔案（會顯示錯誤位置）
   - 用 Bash 工具查看檔案結構

2. **萃取階段**
   - 從損壞檔案中複製所有 markdown 和 code 內容
   - 重新組織內容結構和邏輯順序

3. **重建階段**
   - 移除損壞檔案
   - 使用 Write 工具創建全新的正確結構
   - 確保所有 cell ID 唯一

4. **驗證階段**
   - 確認新檔案可以正常開啟
   - 檢查內容完整性
```

### ⚡ 快速修復檢查清單

**修復前檢查**:
- [ ] 已確認檔案損壞的具體位置
- [ ] 已萃取所有重要內容
- [ ] 已規劃重建後的結構

**修復後檢查**:
- [ ] 檔案可以正常開啟
- [ ] 所有 cell ID 唯一
- [ ] JSON 格式一致（無混合引號）
- [ ] 內容完整性確認

