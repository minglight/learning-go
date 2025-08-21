# learning-go

這是一個學習Go的Project, 主要有三個任務, 當我要求時你要根據我的要求做不同的事情
1. 產生jupyter筆記 : 
- 分章節產生jupyter學習筆記
- 要完全根據指定的那一章的目錄產生筆記. 
- 除了每一章需要說明之外，每一小節也都至少要有從初階到進階的詳細說明(包含example) in markdown block, 然後接著code blocks做詳細的範例, 每一段展示code都要有comment說明
- 如果有特別注意的點, 或是容易出錯的點也都要說明在markdown的部分跟 code block的部分
- 你可以把每一小節拆成一個task來做

參考資訊 
- 要觀看目錄 ( ./目錄.md)
- 要看筆記 example (./ch2/ch2_basic_types_declarations.ipynb)
- 要看考題 example (./ch2/ch2_interview_questions.ipynb)
- 要看考題答案example (./ch2/ch2_interview_answers_review.ipynb)
ex of 說明 in markdown : 
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


### 資料夾結構與命名規範

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


**注意事項**
- 每一個章節跟小節都只需要照`目錄.md`來做, 稱為一個section, 分別是一個markdown block跟一個code block; 然後如果有注意事項可以寫在適合的section內部, 不需要另外開一個section來寫.
- 每一個section的markdown block都要有詳細文字說明 + code example, 然後 code block可以再寫一次example.這樣子印象會比較深刻
- 可以在每一個小節內, 補充你覺得在這個chapter適合的內容(可以多加markdown + code block as小小節), 但要注意不要寫後面chapter才會提到的困難的概念. 
