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

### 🎯 核心原則：使用 MD 作為中間格式

**為了避免 JSON token 限制問題，所有 Jupyter Notebook 創建都應該遵循以下流程：**

1. **撰寫階段**：使用 Markdown 格式撰寫內容（省 60-70% token）
2. **轉換階段**：使用 Go 轉換器生成 .ipynb 文件
3. **維護階段**：只要修改, 都要先確認同一個folder有沒有筆記的raw file(md format), 如果有的話都是修改MD file; 如果沒有MD file, 就先把本來資料寫入MD file之後, 再使用go轉換器

### 🛠️ Tool Selection Rules

| 任務類型 | 使用工具 | 說明 |
|----------|----------|------|
| **新建 notebook** | `Write` → MD → `Bash` 轉換 | 先寫 .md，再用轉換器生成 .ipynb |
| **修改 notebook** | `Edit` MD → `Bash` 轉換 | 修改 .md 源文件，重新轉換 |
| **修復損壞檔案** | `Write` Markdown → `Bash` 轉換 | 萃取內容寫成 .md，重新轉換 |

### 📝 Markdown 源文件格式

詳細的格式規範請參考 [./md_to_ipynb_converter/README.md](./md_to_ipynb_converter/README.md)

**快速參考：**

- **Markdown Cell**: `<!-- MARKDOWN_CELL -->` 和 `<!-- END_MARKDOWN_CELL -->` 包圍
- **Code Cell**: `<!-- CODE_CELL -->` 和 `<!-- END_CODE_CELL -->` 包圍
- **Code Fence**: 使用 ` ```go ` （會被自動移除）
- **所有 cells 都必須有結束標記**

範例：
```markdown
<!-- MARKDOWN_CELL -->
# 標題

說明內容...
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
package main
// ...
```
<!-- END_CODE_CELL -->
```

### 🚀 標準創建 Notebook 流程

#### Step 1: 撰寫 Markdown 源文件

##### 🎯 方法 A: 一次撰寫完成（適合小型章節）

```bash
# 使用 Write 工具一次撰寫完整內容
Write → chN/chN_topic_source.md
```

**優勢**：
- Markdown 格式比 JSON 省 60-70% token
- 最簡單直接的方式
- 人類可讀，易於維護和修改

**適用情境**：
- 章節內容較少（< 15,000 tokens）
- 可以一次性完成所有內容

##### 🔄 方法 B: 分段撰寫（適合大型章節，推薦！）

當章節內容龐大，單次 Write 會超過 token 限制時，使用分段撰寫策略：

```bash
# 第一段：寫入檔案開頭 + 前面幾節
Write → chN/chN_topic_source.md
# 包含：文件開頭、第一節、第二節等

# 第二段：追加中間幾節
Edit → chN/chN_topic_source.md
# 在文件末尾 append 第三節、第四節等

# 第三段：追加剩餘內容
Edit → chN/chN_topic_source.md
# 在文件末尾 append 最後幾節
```

**分段策略**：
1. **規劃階段**：使用 TodoWrite 規劃這一章要拆成幾個部分（建議按「節」來分）
2. **第一次 Write**：寫入前 30-40% 的內容
3. **後續 Edit (append)**：每次追加 30-40% 的內容
4. **確保連貫性**：每段的開頭和結尾要確保格式正確（cell 標記完整）

**優勢**：
- 每次操作的 token 量可控（< 32,000 tokens）
- 避免單次操作超過限制
- 逐步建立完整文件，降低出錯風險
- 仍然保留 Markdown 源文件的所有優點

**注意事項**：
- 每次 Edit 都是在文件**末尾追加**內容
- 確保每段的最後一個 cell 有正確的結束標記 `<!-- END_MARKDOWN_CELL -->` 或 `<!-- END_CODE_CELL -->`
- 下一段的開頭要從新的 cell 標記開始（`<!-- MARKDOWN_CELL -->` 或 `<!-- CODE_CELL -->`）

**範例**：
```markdown
# 第一次 Write 的結尾
<!-- CODE_CELL -->
```go
// 第二節的最後一個範例
```
<!-- END_CODE_CELL -->

# 第二次 Edit 的開頭（append 到文件末尾）
<!-- MARKDOWN_CELL -->
## 第三節：...
<!-- END_MARKDOWN_CELL -->
```

#### Step 2: 執行轉換

```bash
# 使用 Bash 工具執行轉換器
Bash → ./md_to_ipynb_converter/md2ipynb chN/chN_topic_source.md chN/chN_topic.ipynb
```

**轉換器位置**：`/Users/hank/Workspace/hank/learning-go/md_to_ipynb_converter/md2ipynb`

#### Step 3: 驗證結果

```bash
# 1. 驗證 JSON 格式
Bash → python3 -m json.tool chN/chN_topic.ipynb > /dev/null && echo "✅ JSON 正確"

# 2. 預覽 notebook 結構
Read → chN/chN_topic.ipynb (limit: 50)
```

#### Step 4: 清理源文件（可選）

```bash
# 如果不需要保留 Markdown 源文件
Bash → rm chN/chN_topic_source.md
```

### ✨ 小幅修改現有 Notebook

如果只需要修改單一 cell，可以直接使用 NotebookEdit：

```json
{
  "notebook_path": "/path/to/notebook.ipynb",
  "cell_id": "cell-id",
  "new_source": "更新的內容..."
}
```

**注意**：
- 只用於小幅修改（1-2 個 cells）
- 大幅修改應該回到 Markdown 源文件重新轉換
- 絕對不要使用 `insert` 模式

### ⚠️ 重要注意事項

1. **優先使用 Markdown 流程** - 除非是非常小的修改，否則都應該使用 Markdown → 轉換的流程
2. **避免直接 Write .ipynb** - JSON 格式容易超過 token 限制且難以維護
3. **保留源文件選項** - 建議保留 `_source.md` 文件方便未來修改
4. **測試轉換器** - 轉換器已經過完整測試（20 個測試全通過）

### 🔧 修復損壞的 Notebook

如果遇到 Notebook 損壞的情況：

1. **Read 檔案萃取內容** - 即使 JSON 損壞，內容通常還是可讀的
2. **重新組織成 Markdown** - 將內容整理成 Markdown 源文件格式
3. **使用轉換器重建** - 用轉換器生成新的 .ipynb 文件

```bash
# 範例修復流程
Read → chN/broken.ipynb                    # 萃取內容
Write → chN/chN_topic_source.md            # 重新組織成 Markdown
Bash → ./md_to_ipynb_converter/md2ipynb \              # 重新轉換
        chN/chN_topic_source.md \
        chN/chN_topic.ipynb
```

### 📚 轉換器詳細文檔

更多詳細資訊請參考：[./md_to_ipynb_converter/README.md](./md_to_ipynb_converter/README.md)

