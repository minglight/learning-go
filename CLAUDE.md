# learning-go

這是一個學習Go的Project, 主要有三個任務, 當我要求時你要根據我的要求做不同的事情
- 產生jupyter筆記
- 產生面試考題
- 改考卷

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



