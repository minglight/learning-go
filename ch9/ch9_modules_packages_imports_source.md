<!-- MARKDOWN_CELL -->
# 第九章 模組、程式包與匯入

本章將學習如何組織Go程式碼，包括模組(modules)、程式包(packages)和匯入(imports)的概念與使用方式。這些是構建大型Go應用程式的基礎。
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
## Repository、Module 與 Package

Go的程式碼組織架構有三個層級，理解它們之間的關係是掌握Go專案結構的關鍵。
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
### 三個層級的定義

**1. Repository（版本庫）**
- 是版本控制系統（如Git）中管理的程式碼集合
- 通常對應一個Git倉庫
- 可以包含一個或多個模組
- 例如：整個 `github.com/gin-gonic/gin` Git倉庫

**2. Module（模組）**
- 是Go相依性管理的基本單位
- 由 `go.mod` 檔案定義和標識
- 包含一個或多個相關的程式包
- 有自己的版本號（使用語意化版本）
- 一個repository可以包含多個module（較少見）

**3. Package（程式包）**
- 是Go程式碼組織的基本單位
- 同一目錄下所有 `.go` 檔案屬於同一個package
- 透過 `package` 關鍵字宣告
- 提供相關功能的集合
- 是 `import` 語句的目標
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
### 三者的關係

```
Repository（版本庫）
└── Module（模組 - 由 go.mod 定義）
    ├── Package A（程式包 - 根目錄）
    ├── Package B（程式包 - 子目錄 /utils）
    └── Package C（程式包 - 子目錄 /internal/config）
```
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 示範三個層級的概念 */
package main

import "fmt"

func main() {
	fmt.Println("=== Go 程式碼組織的三個層級 ===\n")

	fmt.Println("【Repository】")
	fmt.Println("- 定義：Git 倉庫，版本控制的單位")
	fmt.Println("- 範例：https://github.com/gin-gonic/gin")
	fmt.Println("- 內容：整個專案的所有原始碼")

	fmt.Println("\n【Module】")
	fmt.Println("- 定義：相依性管理的單位，由 go.mod 定義")
	fmt.Println("- 範例：github.com/gin-gonic/gin")
	fmt.Println("- 版本：v1.9.1（使用語意化版本）")
	fmt.Println("- 標識：go.mod 檔案中的 module 路徑")

	fmt.Println("\n【Package】")
	fmt.Println("- 定義：程式碼組織的基本單位")
	fmt.Println("- 範例：github.com/gin-gonic/gin/binding")
	fmt.Println("- 宣告：package binding")
	fmt.Println("- 使用：import \"github.com/gin-gonic/gin/binding\"")
}
// 輸出:
// === Go 程式碼組織的三個層級 ===
//
// 【Repository】
// - 定義：Git 倉庫，版本控制的單位
// - 範例：https://github.com/gin-gonic/gin
// - 內容：整個專案的所有原始碼
//
// 【Module】
// - 定義：相依性管理的單位，由 go.mod 定義
// - 範例：github.com/gin-gonic/gin
// - 版本：v1.9.1（使用語意化版本）
// - 標識：go.mod 檔案中的 module 路徑
//
// 【Package】
// - 定義：程式碼組織的基本單位
// - 範例：github.com/gin-gonic/gin/binding
// - 宣告：package binding
// - 使用：import "github.com/gin-gonic/gin/binding"
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### Module 命名：為什麼有些用 github.com 開頭？

Module 路徑的命名決定了它是**本地模組**還是**遠端模組**。

#### 遠端 Module（使用域名開頭）

當你打算發布 module 供他人使用時，module 路徑**必須**使用可訪問的URL格式：

**格式**：`域名/路徑/專案名`

**常見域名**：
- `github.com/使用者/專案` - GitHub託管
- `gitlab.com/使用者/專案` - GitLab託管
- `bitbucket.org/使用者/專案` - Bitbucket託管
- `golang.org/x/工具` - Go官方擴展包
- `gopkg.in/專案.v1` - gopkg.in服務

**為什麼要用域名？**
1. **唯一性**：確保全球範圍內模組名稱不衝突
2. **可追溯性**：Go工具可以自動下載模組
3. **版本管理**：配合Git標籤進行版本控制
4. **安全性**：驗證模組來源的真實性

#### 本地 Module（不使用域名）

如果 module 僅在本地使用，不打算發布，可以使用簡單名稱：

**格式**：任意有效的模組名（通常是專案名）

**範例**：
- `myproject`
- `company-internal-tool`
- `learning-go-exercises`

**限制**：
- 無法被其他專案透過 `go get` 下載
- 僅適合個人學習或內部專案
- 無法發布到公共模組代理
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* Repository 與 Module 的關係：單一模組 vs 多模組 */
package main

import "fmt"

func main() {
	fmt.Println("=== Repository 與 Module 的關係 ===\n")

	fmt.Println("【常見情況：一個 repository = 一個 module】")
	fmt.Println("repo: github.com/gin-gonic/gin")
	fmt.Println("└── module: github.com/gin-gonic/gin")
	fmt.Println("    go.mod 位於倉庫根目錄")

	fmt.Println("\n【較少見：一個 repository = 多個 module】")
	fmt.Println("repo: github.com/golang/tools")
	fmt.Println("├── module: golang.org/x/tools")
	fmt.Println("│   go.mod 位於根目錄")
	fmt.Println("├── module: golang.org/x/tools/gopls")
	fmt.Println("│   go.mod 位於 gopls/ 子目錄")
	fmt.Println("└── module: golang.org/x/tools/cmd/stringer")
	fmt.Println("    go.mod 位於 cmd/stringer/ 子目錄")

	fmt.Println("\n【為什麼需要多 module？】")
	fmt.Println("1. 不同的版本週期")
	fmt.Println("   └─ 子專案需要獨立發版")
	fmt.Println("2. 減少相依性")
	fmt.Println("   └─ 避免引入不必要的相依")
	fmt.Println("3. 大型單一倉庫（Monorepo）")
	fmt.Println("   └─ 管理多個獨立專案")

	fmt.Println("\n【建議】")
	fmt.Println("對於大多數專案，使用單一 module 即可")
	fmt.Println("只有在明確需要時才使用多 module 結構")
}
// 輸出:
// === Repository 與 Module 的關係 ===
//
// 【常見情況：一個 repository = 一個 module】
// repo: github.com/gin-gonic/gin
// └── module: github.com/gin-gonic/gin
//     go.mod 位於倉庫根目錄
//
// 【較少見：一個 repository = 多個 module】
// repo: github.com/golang/tools
// ├── module: golang.org/x/tools
// │   go.mod 位於根目錄
// ├── module: golang.org/x/tools/gopls
// │   go.mod 位於 gopls/ 子目錄
// └── module: golang.org/x/tools/cmd/stringer
//     go.mod 位於 cmd/stringer/ 子目錄
//
// 【為什麼需要多 module？】
// 1. 不同的版本週期
//    └─ 子專案需要獨立發版
// 2. 減少相依性
//    └─ 避免引入不必要的相依
// 3. 大型單一倉庫（Monorepo）
//    └─ 管理多個獨立專案
//
// 【建議】
// 對於大多數專案，使用單一 module 即可
// 只有在明確需要時才使用多 module 結構
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## go.mod

go.mod 檔案是 Go module 的核心，它定義了模組的路徑、Go版本需求和相依性。每個 module 都必須有一個 go.mod 檔案位於模組根目錄。

go.mod 檔案包含五個主要指令：

**1. `module`** - 宣告模組路徑
- 定義此模組的匯入路徑
- 其他專案使用此路徑來匯入

**2. `go`** - 指定Go的最低版本需求
- 影響語言特性的可用性
- 例如：`go 1.21`

**3. `require`** - 宣告模組的直接相依性
- 列出專案需要的外部模組
- 包含版本號（使用語意化版本）
- 例如：`github.com/gin-gonic/gin v1.9.1`

**4. `replace`** - 替換相依性的來源（通常用於開發階段）
- **本地開發**：使用本地版本進行開發測試
  - `replace github.com/company/lib => ../lib`
- **Fork 替換**：使用自己修復的版本
  - `replace github.com/original/lib => github.com/myuser/lib v1.2.3`
- **版本覆蓋**：強制使用特定版本或降級
  - `replace github.com/lib/pkg v2.0.0 => github.com/lib/pkg v1.9.5`
- **路徑重定向**：模組搬家時重定向
  - `replace old.com/pkg => new.com/pkg v1.0.0`

**5. `exclude`** - 排除特定版本的相依性（較少使用）
- **排除有 bug 的版本**：防止使用問題版本
  - `exclude github.com/some/lib v1.5.0`
- **排除安全漏洞版本**：確保安全性
  - `exclude github.com/vulnerable/lib v2.1.0`
- **排除不相容版本**：強制使用相容版本
  - `exclude github.com/pkg/tool v3.0.0`
- 注意：只影響當前模組，不會傳遞到其他專案

**重要特性**：
- `replace` 和 `exclude` 只影響當前模組的建構
- Go 工具會在 `exclude` 後自動選擇其他可用版本
- `replace` 常用於開發階段，發布時通常會移除
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 完整的 go.mod 檔案結構 */
package main

import "fmt"

func main() {
	fmt.Println("=== go.mod 檔案結構 ===\n")

	fmt.Println("【完整 go.mod 範例】")
	fmt.Println("module github.com/myuser/myproject")
	fmt.Println("")
	fmt.Println("go 1.21")
	fmt.Println("")
	fmt.Println("require (")
	fmt.Println("    github.com/gin-gonic/gin v1.9.1")
	fmt.Println("    github.com/go-sql-driver/mysql v1.7.1")
	fmt.Println("    github.com/some/lib v1.6.0")
	fmt.Println(")")
	fmt.Println("")
	fmt.Println("replace (")
	fmt.Println("    // 本地開發")
	fmt.Println("    github.com/company/shared => ../shared")
	fmt.Println("    // 使用 fork 版本")
	fmt.Println("    github.com/old/lib => github.com/myuser/lib v1.0.1")
	fmt.Println(")")
	fmt.Println("")
	fmt.Println("exclude (")
	fmt.Println("    // 已知問題版本")
	fmt.Println("    github.com/some/lib v1.5.0  // memory leak")
	fmt.Println("    // 安全漏洞版本")
	fmt.Println("    github.com/other/pkg v2.0.0  // CVE-2023-12345")
	fmt.Println(")")

	fmt.Println("\n【各指令說明】")
	fmt.Println("1. module: 定義模組路徑")
	fmt.Println("2. go: 指定最低 Go 版本")
	fmt.Println("3. require: 列出相依模組")
	fmt.Println("4. replace: 替換模組來源（開發用）")
	fmt.Println("5. exclude: 排除特定版本（安全用）")
}
// 輸出:
// === go.mod 檔案結構 ===
//
// 【完整 go.mod 範例】
// module github.com/myuser/myproject
//
// go 1.21
//
// require (
//     github.com/gin-gonic/gin v1.9.1
//     github.com/go-sql-driver/mysql v1.7.1
//     github.com/some/lib v1.6.0
// )
//
// replace (
//     // 本地開發
//     github.com/company/shared => ../shared
//     // 使用 fork 版本
//     github.com/old/lib => github.com/myuser/lib v1.0.1
// )
//
// exclude (
//     // 已知問題版本
//     github.com/some/lib v1.5.0  // memory leak
//     // 安全漏洞版本
//     github.com/other/pkg v2.0.0  // CVE-2023-12345
// )
//
// 【各指令說明】
// 1. module: 定義模組路徑
// 2. go: 指定最低 Go 版本
// 3. require: 列出相依模組
// 4. replace: 替換模組來源（開發用）
// 5. exclude: 排除特定版本（安全用）
```
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
```go
/* 使用 go mod init 創建新模組 */
package main

import "fmt"

func main() {
	fmt.Println("=== 創建新模組 ===\n")

	fmt.Println("【指令】")
	fmt.Println("$ go mod init github.com/myuser/myproject")

	fmt.Println("\n【生成的 go.mod】")
	fmt.Println("module github.com/myuser/myproject")
	fmt.Println("")
	fmt.Println("go 1.21")

	fmt.Println("\n【模組路徑選擇】")
	fmt.Println("1. 發布到公共平台：")
	fmt.Println("   $ go mod init github.com/username/projectname")
	fmt.Println("   └─ 必須使用實際可訪問的URL")
	fmt.Println("")
	fmt.Println("2. 僅本地使用：")
	fmt.Println("   $ go mod init myproject")
	fmt.Println("   └─ 可以使用任意名稱")
	fmt.Println("")
	fmt.Println("3. 公司內部專案：")
	fmt.Println("   $ go mod init company.com/team/project")
	fmt.Println("   └─ 使用公司域名")

	fmt.Println("\n【最佳實踐】")
	fmt.Println("✓ 即使是本地專案，也建議使用完整路徑")
	fmt.Println("  └─ 方便未來發布或分享")
	fmt.Println("✓ 路徑應該全部小寫")
	fmt.Println("✓ 使用短橫線分隔單字（my-project）")
}
// 輸出:
// === 創建新模組 ===
//
// 【指令】
// $ go mod init github.com/myuser/myproject
//
// 【生成的 go.mod】
// module github.com/myuser/myproject
//
// go 1.21
//
// 【模組路徑選擇】
// 1. 發布到公共平台：
//    $ go mod init github.com/username/projectname
//    └─ 必須使用實際可訪問的URL
//
// 2. 僅本地使用：
//    $ go mod init myproject
//    └─ 可以使用任意名稱
//
// 3. 公司內部專案：
//    $ go mod init company.com/team/project
//    └─ 使用公司域名
//
// 【最佳實踐】
// ✓ 即使是本地專案，也建議使用完整路徑
//   └─ 方便未來發布或分享
// ✓ 路徑應該全部小寫
// ✓ 使用短橫線分隔單字（my-project）
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 建構 Package

Package 是 Go 程式組織的基本單位。一個 package 包含在同一目錄下的所有 Go 檔案，這些檔案共享相同的 package 宣告。建構 package 涉及匯入匯出、命名、組織等多個面向。
<!-- END_MARKDOWN_CELL -->

<!-- MARKDOWN_CELL -->
### 匯入與匯出

Go 使用大小寫來控制可見性：
- **大寫開頭**：公開(exported)，可被其他 package 存取
- **小寫開頭**：私有(unexported)，只能在同一 package 內存取

這種設計讓 API 的設計更加明確，不需要額外的關鍵字。
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 演示匯入與匯出的概念 */
package main

import "fmt"

// PublicFunction 是公開的函式（大寫開頭）
// 可以被其他 package 匯入和呼叫
func PublicFunction() {
	fmt.Println("這是公開函式，可被其他 package 呼叫")
}

// privateFunction 是私有的函式（小寫開頭）
// 只能在同一 package 內呼叫
func privateFunction() {
	fmt.Println("這是私有函式，只能在同一 package 內呼叫")
}

// PublicVar 是公開的變數
var PublicVar = "公開變數"

// privateVar 是私有的變數
var privateVar = "私有變數"

// PublicStruct 是公開的結構
type PublicStruct struct {
	PublicField  string // 公開欄位
	privateField string // 私有欄位
}

func main() {
	fmt.Println("=== 可見性規則 ===\n")

	// 在同一 package 內，所有識別符都可存取
	PublicFunction()
	privateFunction()

	fmt.Println(PublicVar)
	fmt.Println(privateVar)

	s := PublicStruct{
		PublicField:  "可以存取",
		privateField: "同一 package 內可以存取",
	}
	fmt.Printf("\n結構體：%+v\n", s)

	fmt.Println("\n【重要】")
	fmt.Println("如果其他 package 匯入此 package：")
	fmt.Println("✓ 可以存取：PublicFunction, PublicVar, PublicStruct, PublicField")
	fmt.Println("✗ 無法存取：privateFunction, privateVar, privateField")
}
// 輸出:
// === 可見性規則 ===
//
// 這是公開函式，可被其他 package 呼叫
// 這是私有函式，只能在同一 package 內呼叫
// 公開變數
// 私有變數
//
// 結構體：{PublicField:可以存取 privateField:同一 package 內可以存取}
//
// 【重要】
// 如果其他 package 匯入此 package：
// ✓ 可以存取：PublicFunction, PublicVar, PublicStruct, PublicField
// ✗ 無法存取：privateFunction, privateVar, privateField
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 建立與使用 Package

要建立自己的 package，需要：
1. 建立新目錄
2. 在目錄中建立 Go 檔案，以 `package packagename` 開頭
3. 在其他 package 中匯入並使用

**重要概念**：
- **Import 的是什麼**：import 匯入的是 **package 路徑**，而非檔案或目錄
- **Package 路徑**：module 路徑 + 相對於 module 根目錄的路徑
- **Package 名稱**：在程式碼中使用的識別符，通常是路徑的最後一個元素
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* Package 的建立與使用 */
package main

import "fmt"

func main() {
	fmt.Println("=== Package 的建立與使用 ===\n")

	fmt.Println("【專案結構】")
	fmt.Println("myproject/")
	fmt.Println("├── go.mod                          module: github.com/user/myproject")
	fmt.Println("├── main.go                         package main")
	fmt.Println("├── mathutil/")
	fmt.Println("│   └── calculator.go               package mathutil")
	fmt.Println("└── stringutil/")
	fmt.Println("    └── helper.go                   package stringutil")

	fmt.Println("\n【mathutil/calculator.go 內容】")
	fmt.Println("package mathutil")
	fmt.Println("")
	fmt.Println("// Add 是公開函式")
	fmt.Println("func Add(a, b int) int {")
	fmt.Println("    return a + b")
	fmt.Println("}")

	fmt.Println("\n【在 main.go 中使用】")
	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("import (")
	fmt.Println(`    "fmt"`)
	fmt.Println(`    "github.com/user/myproject/mathutil"    // import 路徑`)
	fmt.Println(")")
	fmt.Println("")
	fmt.Println("func main() {")
	fmt.Println("    result := mathutil.Add(1, 2)            // 使用 package 名稱")
	fmt.Println("    fmt.Println(result)")
	fmt.Println("}")

	fmt.Println("\n【關鍵概念】")
	fmt.Println("Import 路徑：github.com/user/myproject/mathutil")
	fmt.Println("  ├─ Module 路徑：github.com/user/myproject")
	fmt.Println("  └─ 相對路徑：mathutil")
	fmt.Println("")
	fmt.Println("使用時：mathutil.Add()")
	fmt.Println("  └─ mathutil 是 package 名稱（通常與最後的路徑元素相同）")
}
// 輸出:
// === Package 的建立與使用 ===
//
// 【專案結構】
// myproject/
// ├── go.mod                          module: github.com/user/myproject
// ├── main.go                         package main
// ├── mathutil/
// │   └── calculator.go               package mathutil
// └── stringutil/
//     └── helper.go                   package stringutil
//
// 【mathutil/calculator.go 內容】
// package mathutil
//
// // Add 是公開函式
// func Add(a, b int) int {
//     return a + b
// }
//
// 【在 main.go 中使用】
// package main
//
// import (
//     "fmt"
//     "github.com/user/myproject/mathutil"    // import 路徑
// )
//
// func main() {
//     result := mathutil.Add(1, 2)            // 使用 package 名稱
//     fmt.Println(result)
// }
//
// 【關鍵概念】
// Import 路徑：github.com/user/myproject/mathutil
//   ├─ Module 路徑：github.com/user/myproject
//   └─ 相對路徑：mathutil
//
// 使用時：mathutil.Add()
//   └─ mathutil 是 package 名稱（通常與最後的路徑元素相同）
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### Import 的深入理解

讓我們更詳細地了解 import 的機制，特別是當路徑包含多層目錄時如何使用。
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* Import 路徑與 Package 名稱 */
package main

import "fmt"

func main() {
	fmt.Println("=== Import 路徑 vs Package 名稱 ===\n")

	fmt.Println("【專案結構】")
	fmt.Println("ecommerce/")
	fmt.Println("├── go.mod                          module: github.com/shop/ecommerce")
	fmt.Println("└── internal/")
	fmt.Println("    └── payment/")
	fmt.Println("        └── stripe.go               package payment")

	fmt.Println("\n【stripe.go 的內容】")
	fmt.Println("package payment  // Package 名稱")
	fmt.Println("")
	fmt.Println("func ProcessPayment(amount float64) error {")
	fmt.Println("    // 處理支付")
	fmt.Println("    return nil")
	fmt.Println("}")

	fmt.Println("\n【如何 Import 和使用】")
	fmt.Println("import \"github.com/shop/ecommerce/internal/payment\"")
	fmt.Println("         │                                    │")
	fmt.Println("         │                                    └─ 路徑：internal/payment")
	fmt.Println("         └─ Module 路徑：github.com/shop/ecommerce")
	fmt.Println("")
	fmt.Println("使用方式：")
	fmt.Println("err := payment.ProcessPayment(100.0)")
	fmt.Println("       │")
	fmt.Println("       └─ 使用 package 名稱 'payment'（不是 'internal'）")

	fmt.Println("\n【重要規則】")
	fmt.Println("1. Import 語句使用完整路徑")
	fmt.Println("2. 程式碼中使用 package 名稱（package 關鍵字後的名稱）")
	fmt.Println("3. Package 名稱通常是路徑的最後一個元素")
	fmt.Println("4. Package 名稱與目錄名稱可以不同（但不建議）")
}
// 輸出:
// === Import 路徑 vs Package 名稱 ===
//
// 【專案結構】
// ecommerce/
// ├── go.mod                          module: github.com/shop/ecommerce
// └── internal/
//     └── payment/
//         └── stripe.go               package payment
//
// 【stripe.go 的內容】
// package payment  // Package 名稱
//
// func ProcessPayment(amount float64) error {
//     // 處理支付
//     return nil
// }
//
// 【如何 Import 和使用】
// import "github.com/shop/ecommerce/internal/payment"
//          │                                    │
//          │                                    └─ 路徑：internal/payment
//          └─ Module 路徑：github.com/shop/ecommerce
//
// 使用方式：
// err := payment.ProcessPayment(100.0)
//        │
//        └─ 使用 package 名稱 'payment'（不是 'internal'）
//
// 【重要規則】
// 1. Import 語句使用完整路徑
// 2. 程式碼中使用 package 名稱（package 關鍵字後的名稱）
// 3. Package 名稱通常是路徑的最後一個元素
// 4. Package 名稱與目錄名稱可以不同（但不建議）
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 為 Package 命名

Package 命名是 Go 開發中的重要環節，好的命名讓程式碼更易讀易用。

**基本規則**：
1. **預設情況**：package 名稱 = 目錄名稱（路徑的最後一個元素）
2. **全部小寫**：不使用大寫字母或底線
3. **簡短且描述性**：清楚表達 package 的用途
4. **單數形式**：通常使用單數而非複數（如 `user` 而非 `users`）

**何時 package 名稱可以不同於目錄名稱**：

雖然可以讓 package 名稱與目錄名稱不同，但這會造成混淆，**強烈不建議**，除非以下特殊情況：

1. **main package**：可執行程式必須使用 `package main`
2. **測試 package**：使用 `_test` 後綴進行黑盒測試
3. **版本後綴**：主版本 v2+ 的 package
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* Package 命名規則 */
package main

import "fmt"

func main() {
	fmt.Println("=== Package 命名規則 ===\n")

	fmt.Println("【好的 Package 命名】")
	fmt.Println("✓ http       - 處理HTTP協議")
	fmt.Println("✓ json       - JSON編碼/解碼")
	fmt.Println("✓ time       - 時間處理")
	fmt.Println("✓ crypto     - 加密功能")
	fmt.Println("✓ parser     - 解析功能")
	fmt.Println("✓ user       - 使用者管理（單數）")

	fmt.Println("\n【應該避免的命名】")
	fmt.Println("✗ util       - 太泛用，無意義")
	fmt.Println("✗ common     - 不夠具體")
	fmt.Println("✗ helpers    - 複數形式")
	fmt.Println("✗ my_utils   - 使用底線")
	fmt.Println("✗ HttpUtil   - 混合大小寫")
	fmt.Println("✗ stuff      - 語意不清")

	fmt.Println("\n【命名原則】")
	fmt.Println("1. 簡短 - 通常一個單字")
	fmt.Println("2. 小寫 - 全部小寫字母")
	fmt.Println("3. 描述性 - 清楚表達用途")
	fmt.Println("4. 單數 - 使用單數形式")
	fmt.Println("5. 避免縮寫 - 除非是廣為人知的縮寫")
}
// 輸出:
// === Package 命名規則 ===
//
// 【好的 Package 命名】
// ✓ http       - 處理HTTP協議
// ✓ json       - JSON編碼/解碼
// ✓ time       - 時間處理
// ✓ crypto     - 加密功能
// ✓ parser     - 解析功能
// ✓ user       - 使用者管理（單數）
//
// 【應該避免的命名】
// ✗ util       - 太泛用，無意義
// ✗ common     - 不夠具體
// ✗ helpers    - 複數形式
// ✗ my_utils   - 使用底線
// ✗ HttpUtil   - 混合大小寫
// ✗ stuff      - 語意不清
//
// 【命名原則】
// 1. 簡短 - 通常一個單字
// 2. 小寫 - 全部小寫字母
// 3. 描述性 - 清楚表達用途
// 4. 單數 - 使用單數形式
// 5. 避免縮寫 - 除非是廣為人知的縮寫
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 如何組織你的模組

好的模組組織能讓程式碼更易維護和理解。Go 社群有一些約定俗成的最佳實踐。

**常見的專案結構模式**：

1. **扁平結構**（適合小型專案）
   - 所有 package 都在根目錄下
   - 簡單直觀，適合 10 個以下的 package

2. **分層結構**（適合中大型專案）
   - 按功能分組到不同目錄
   - 使用 `cmd/`、`internal/`、`pkg/` 等標準目錄

**標準目錄慣例**：
- **`cmd/`** - 存放可執行程式的 main package
- **`internal/`** - 私有程式碼，無法被外部 import
- **`pkg/`** - 可被外部使用的公開程式庫（較少使用）
- **`api/`** - API 定義（如 OpenAPI/Swagger）
- **`web/`** - Web 靜態資源
- **`configs/`** - 設定檔範本
- **`test/`** - 額外的測試資料和工具

**組織原則**：
- 按**功能領域**而非**技術層次**組織
- 避免過度分層
- 相關功能放在一起
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 專案結構範例 */
package main

import "fmt"

func main() {
	fmt.Println("=== 推薦的專案結構 ===\n")

	fmt.Println("【小型專案（扁平結構）】")
	fmt.Println("myapp/")
	fmt.Println("├── go.mod")
	fmt.Println("├── main.go           // package main")
	fmt.Println("├── user.go           // package main")
	fmt.Println("├── auth.go           // package main")
	fmt.Println("└── database.go       // package main")
	fmt.Println("")
	fmt.Println("特點：所有程式碼在同一個 package")
	fmt.Println("適用：簡單的命令列工具、腳本")

	fmt.Println("\n【中型專案（分層結構）】")
	fmt.Println("myapp/")
	fmt.Println("├── go.mod")
	fmt.Println("├── main.go           // package main")
	fmt.Println("├── user/")
	fmt.Println("│   ├── user.go       // package user")
	fmt.Println("│   └── repo.go       // package user")
	fmt.Println("├── auth/")
	fmt.Println("│   └── auth.go       // package auth")
	fmt.Println("└── database/")
	fmt.Println("    └── db.go         // package database")
	fmt.Println("")
	fmt.Println("特點：按功能領域分 package")
	fmt.Println("適用：一般的應用程式")

	fmt.Println("\n【大型專案（標準佈局）】")
	fmt.Println("myapp/")
	fmt.Println("├── go.mod")
	fmt.Println("├── cmd/")
	fmt.Println("│   ├── server/")
	fmt.Println("│   │   └── main.go   // package main")
	fmt.Println("│   └── worker/")
	fmt.Println("│       └── main.go   // package main")
	fmt.Println("├── internal/")
	fmt.Println("│   ├── user/         // 私有，不可被外部 import")
	fmt.Println("│   ├── auth/")
	fmt.Println("│   └── database/")
	fmt.Println("├── pkg/")
	fmt.Println("│   └── mathutil/     // 公開，可被外部使用")
	fmt.Println("├── api/")
	fmt.Println("│   └── openapi.yaml")
	fmt.Println("└── configs/")
	fmt.Println("    └── config.yaml")
	fmt.Println("")
	fmt.Println("特點：使用標準目錄慣例")
	fmt.Println("適用：複雜的應用、可重用的程式庫")

	fmt.Println("\n【組織原則】")
	fmt.Println("✓ 按功能領域組織（user, order, payment）")
	fmt.Println("✗ 按技術層次組織（models, controllers, services）")
	fmt.Println("")
	fmt.Println("✓ 每個 package 有清晰的職責")
	fmt.Println("✗ 避免 util, common, helpers 等泛用名稱")
}
// 輸出:
// === 推薦的專案結構 ===
//
// 【小型專案（扁平結構）】
// myapp/
// ├── go.mod
// ├── main.go           // package main
// ├── user.go           // package main
// ├── auth.go           // package main
// └── database.go       // package main
//
// 特點：所有程式碼在同一個 package
// 適用：簡單的命令列工具、腳本
//
// 【中型專案（分層結構）】
// myapp/
// ├── go.mod
// ├── main.go           // package main
// ├── user/
// │   ├── user.go       // package user
// │   └── repo.go       // package user
// ├── auth/
// │   └── auth.go       // package auth
// └── database/
//     └── db.go         // package database
//
// 特點：按功能領域分 package
// 適用：一般的應用程式
//
// 【大型專案（標準佈局）】
// myapp/
// ├── go.mod
// ├── cmd/
// │   ├── server/
// │   │   └── main.go   // package main
// │   └── worker/
// │       └── main.go   // package main
// ├── internal/
// │   ├── user/         // 私有，不可被外部 import
// │   ├── auth/
// │   └── database/
// ├── pkg/
// │   └── mathutil/     // 公開，可被外部使用
// ├── api/
// │   └── openapi.yaml
// └── configs/
//     └── config.yaml
//
// 特點：使用標準目錄慣例
// 適用：複雜的應用、可重用的程式庫
//
// 【組織原則】
// ✓ 按功能領域組織（user, order, payment）
// ✗ 按技術層次組織（models, controllers, services）
//
// ✓ 每個 package 有清晰的職責
// ✗ 避免 util, common, helpers 等泛用名稱
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 覆寫程式包的名稱

有時需要為匯入的 package 指定別名，特別是當遇到名稱衝突或需要提高可讀性時。

**使用別名的場景**：
1. **避免名稱衝突** - 兩個 package 有相同名稱
2. **簡化長名稱** - 縮短冗長的 package 名稱
3. **提高可讀性** - 讓程式碼意圖更清楚
4. **版本區分** - 同時使用同一 package 的不同版本

**特殊匯入**：
- **點匯入 (`.`)** - 將 package 內容直接匯入當前命名空間（不推薦）
- **底線匯入 (`_`)** - 僅執行 package 的 init 函式
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* Package 別名範例 */
package main

import (
	"fmt"
	// 使用別名避免名稱衝突
	stdjson "encoding/json"
)

func main() {
	fmt.Println("=== Package 別名 ===\n")

	fmt.Println("【場景一：避免名稱衝突】")
	fmt.Println("import (")
	fmt.Println(`    stdjson "encoding/json"`)
	fmt.Println(`    customjson "company.com/json"`)
	fmt.Println(")")
	fmt.Println("使用：")
	fmt.Println("  stdjson.Marshal(data)")
	fmt.Println("  customjson.Marshal(data)")

	// 實際使用標準庫
	data := map[string]string{"name": "Go"}
	bytes, _ := stdjson.Marshal(data)
	fmt.Printf("\n實際輸出：%s\n", bytes)

	fmt.Println("\n【場景二：簡化長名稱】")
	fmt.Println(`import pb "github.com/company/project/api/protobuf/v1"`)
	fmt.Println("使用：pb.Request{} 而非 v1.Request{}")

	fmt.Println("\n【場景三：提高可讀性】")
	fmt.Println(`import mysql "github.com/go-sql-driver/mysql"`)
	fmt.Println(`import postgres "github.com/lib/pq"`)
	fmt.Println("清楚區分不同的資料庫驅動")

	fmt.Println("\n【特殊匯入：點 (.) 匯入 - 不推薦】")
	fmt.Println(`import . "fmt"`)
	fmt.Println("效果：可直接使用 Println() 而非 fmt.Println()")
	fmt.Println("問題：")
	fmt.Println("  ✗ 不清楚函式來自哪個 package")
	fmt.Println("  ✗ 容易造成名稱衝突")
	fmt.Println("  ✗ 降低程式碼可讀性")
	fmt.Println("使用時機：僅在測試程式碼中偶爾使用")

	fmt.Println("\n【特殊匯入：底線 (_) 匯入】")
	fmt.Println(`import _ "github.com/lib/pq"`)
	fmt.Println("用途：")
	fmt.Println("  ✓ 僅執行 package 的 init() 函式")
	fmt.Println("  ✓ 註冊驅動程式（如資料庫驅動）")
	fmt.Println("  ✓ 不直接使用 package 的函式")
	fmt.Println("")
	fmt.Println("範例：")
	fmt.Println("  import (")
	fmt.Println(`      "database/sql"`)
	fmt.Println(`      _ "github.com/lib/pq"  // 註冊 PostgreSQL 驅動`)
	fmt.Println("  )")
	fmt.Println(`  db, _ := sql.Open("postgres", connStr)`)
}
// 輸出:
// === Package 別名 ===
//
// 【場景一：避免名稱衝突】
// import (
//     stdjson "encoding/json"
//     customjson "company.com/json"
// )
// 使用：
//   stdjson.Marshal(data)
//   customjson.Marshal(data)
//
// 實際輸出：{"name":"Go"}
//
// 【場景二：簡化長名稱】
// import pb "github.com/company/project/api/protobuf/v1"
// 使用：pb.Request{} 而非 v1.Request{}
//
// 【場景三：提高可讀性】
// import mysql "github.com/go-sql-driver/mysql"
// import postgres "github.com/lib/pq"
// 清楚區分不同的資料庫驅動
//
// 【特殊匯入：點 (.) 匯入 - 不推薦】
// import . "fmt"
// 效果：可直接使用 Println() 而非 fmt.Println()
// 問題：
//   ✗ 不清楚函式來自哪個 package
//   ✗ 容易造成名稱衝突
//   ✗ 降低程式碼可讀性
// 使用時機：僅在測試程式碼中偶爾使用
//
// 【特殊匯入：底線 (_) 匯入】
// import _ "github.com/lib/pq"
// 用途：
//   ✓ 僅執行 package 的 init() 函式
//   ✓ 註冊驅動程式（如資料庫驅動）
//   ✓ 不直接使用 package 的函式
//
// 範例：
//   import (
//       "database/sql"
//       _ "github.com/lib/pq"  // 註冊 PostgreSQL 驅動
//   )
//   db, _ := sql.Open("postgres", connStr)
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 程式包註解與 godoc

Go 有內建的文件系統 `godoc`，透過特定格式的註解自動產生文件。良好的文件是專業 package 的標誌。

**Package 文件規則**：
- Package 註解寫在 `package` 語句之前
- 以 "Package packagename" 開頭
- 解釋 package 的用途和使用方式
- 通常寫在 `doc.go` 檔案中（大型 package）

**函式/型態文件規則**：
- 註解直接寫在宣告之前，不留空行
- 以被文件化的名稱開頭
- 使用完整句子，以句號結尾
- 公開的識別符都應該有文件

**查看文件**：
- `go doc package` - 查看 package 文件
- `go doc package.Function` - 查看特定函式文件
- `godoc -http=:6060` - 啟動本地文件伺服器
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* godoc 文件範例 */
package main

import "fmt"

func main() {
	fmt.Println("=== godoc 文件規則 ===\n")

	fmt.Println("【Package 文件】")
	fmt.Println("// Package mathutil 提供數學運算的實用函式。")
	fmt.Println("//")
	fmt.Println("// 此 package 包含常用的數學計算，包括：")
	fmt.Println("//   - 基本算術運算")
	fmt.Println("//   - 統計函式")
	fmt.Println("//   - 數值轉換")
	fmt.Println("package mathutil")

	fmt.Println("\n【函式文件】")
	fmt.Println("// Add 將兩個整數相加並回傳結果。")
	fmt.Println("//")
	fmt.Println("// 範例：")
	fmt.Println("//")
	fmt.Println("//     result := Add(1, 2)  // result = 3")
	fmt.Println("//")
	fmt.Println("func Add(a, b int) int {")
	fmt.Println("    return a + b")
	fmt.Println("}")

	fmt.Println("\n【型態文件】")
	fmt.Println("// Calculator 提供一系列計算方法。")
	fmt.Println("// 使用 NewCalculator 來建立實例。")
	fmt.Println("type Calculator struct {")
	fmt.Println("    // 私有欄位不需要文件")
	fmt.Println("    value int")
	fmt.Println("}")

	fmt.Println("\n【文件規則】")
	fmt.Println("✓ 以被文件化的名稱開頭")
	fmt.Println("✓ 使用完整句子")
	fmt.Println("✓ 第一句應該是概要（會顯示在列表中）")
	fmt.Println("✓ 可以包含範例程式碼")
	fmt.Println("✓ 使用空白註解行分段")

	fmt.Println("\n【查看文件指令】")
	fmt.Println("$ go doc mathutil")
	fmt.Println("$ go doc mathutil.Add")
	fmt.Println("$ godoc -http=:6060  # 啟動本地文件伺服器")
}
// 輸出:
// === godoc 文件規則 ===
//
// 【Package 文件】
// // Package mathutil 提供數學運算的實用函式。
// //
// // 此 package 包含常用的數學計算，包括：
// //   - 基本算術運算
// //   - 統計函式
// //   - 數值轉換
// package mathutil
//
// 【函式文件】
// // Add 將兩個整數相加並回傳結果。
// //
// // 範例：
// //
// //     result := Add(1, 2)  // result = 3
// //
// func Add(a, b int) int {
//     return a + b
// }
//
// 【型態文件】
// // Calculator 提供一系列計算方法。
// // 使用 NewCalculator 來建立實例。
// type Calculator struct {
//     // 私有欄位不需要文件
//     value int
// }
//
// 【文件規則】
// ✓ 以被文件化的名稱開頭
// ✓ 使用完整句子
// ✓ 第一句應該是概要（會顯示在列表中）
// ✓ 可以包含範例程式碼
// ✓ 使用空白註解行分段
//
// 【查看文件指令】
// $ go doc mathutil
// $ go doc mathutil.Add
// $ godoc -http=:6060  # 啟動本地文件伺服器
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### internal Package

`internal` 是 Go 的特殊目錄名稱，用於建立私有 package。放在 `internal/` 目錄下的 package 只能被其父目錄及子目錄的程式碼 import。

**用途**：
- 防止外部專案依賴內部實作
- 保留重構的自由度
- 清楚標示私有 API

**規則**：
- `internal/` 可以出現在任何層級
- 僅其父目錄樹可以 import
- 對外部 module 完全不可見

**範例**：
```
myproject/
├── internal/
│   └── auth/          # 只有 myproject 可以 import
└── pkg/
    └── public/        # 任何人都可以 import
```
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* internal package 說明 */
package main

import "fmt"

func main() {
	fmt.Println("=== internal Package ===\n")

	fmt.Println("【專案結構】")
	fmt.Println("github.com/user/myproject/")
	fmt.Println("├── go.mod")
	fmt.Println("├── main.go")
	fmt.Println("├── internal/")
	fmt.Println("│   ├── auth/         # 私有：僅 myproject 可用")
	fmt.Println("│   └── database/     # 私有：僅 myproject 可用")
	fmt.Println("└── pkg/")
	fmt.Println("    └── mathutil/     # 公開：任何人可用")

	fmt.Println("\n【Import 規則】")
	fmt.Println("✓ 允許：myproject/main.go import myproject/internal/auth")
	fmt.Println("✗ 禁止：other-project import myproject/internal/auth")
	fmt.Println("")
	fmt.Println("原因：internal 目錄對外部 module 不可見")

	fmt.Println("\n【巢狀 internal】")
	fmt.Println("myproject/")
	fmt.Println("├── internal/")
	fmt.Println("│   └── auth/")
	fmt.Println("│       └── internal/")
	fmt.Println("│           └── crypto/  # 僅 auth 可用")
	fmt.Println("")
	fmt.Println("✓ 允許：auth import auth/internal/crypto")
	fmt.Println("✗ 禁止：myproject/main.go import auth/internal/crypto")

	fmt.Println("\n【使用時機】")
	fmt.Println("✓ 內部實作細節")
	fmt.Println("✓ 實驗性功能")
	fmt.Println("✓ 可能變更的 API")
	fmt.Println("✓ 不希望被外部依賴的程式碼")

	fmt.Println("\n【最佳實踐】")
	fmt.Println("1. 預設使用 internal，需要時才公開")
	fmt.Println("2. 公開的 API 要慎重設計（難以變更）")
	fmt.Println("3. internal 給你重構的自由")
}
// 輸出:
// === internal Package ===
//
// 【專案結構】
// github.com/user/myproject/
// ├── go.mod
// ├── main.go
// ├── internal/
// │   ├── auth/         # 私有：僅 myproject 可用
// │   └── database/     # 私有：僅 myproject 可用
// └── pkg/
//     └── mathutil/     # 公開：任何人可用
//
// 【Import 規則】
// ✓ 允許：myproject/main.go import myproject/internal/auth
// ✗ 禁止：other-project import myproject/internal/auth
//
// 原因：internal 目錄對外部 module 不可見
//
// 【巢狀 internal】
// myproject/
// ├── internal/
// │   └── auth/
// │       └── internal/
// │           └── crypto/  # 僅 auth 可用
//
// ✓ 允許：auth import auth/internal/crypto
// ✗ 禁止：myproject/main.go import auth/internal/crypto
//
// 【使用時機】
// ✓ 內部實作細節
// ✓ 實驗性功能
// ✓ 可能變更的 API
// ✓ 不希望被外部依賴的程式碼
//
// 【最佳實踐】
// 1. 預設使用 internal，需要時才公開
// 2. 公開的 API 要慎重設計（難以變更）
// 3. internal 給你重構的自由
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### init 函式：請量避免

`init` 函式在 package 被 import 時自動執行，但它有許多缺點，應該謹慎使用。

**init 函式特性**：
- 在 package 初始化時自動執行
- 不能被呼叫或引用
- 一個 package 可以有多個 init 函式
- 執行順序：相依 package → 當前 package 的變數初始化 → init 函式

**為什麼要避免**：
- 難以測試（無法控制執行）
- 隱藏的相依性
- 增加啟動時間
- 難以理解程式碼流程
- 無法傳遞參數或回傳錯誤

**合理使用場景**（僅限）：
- 註冊驅動程式（如 database/sql 驅動）
- 驗證編譯時常數
- 一次性的全域設定

**替代方案**：
使用明確的初始化函式（如 `New()`, `Init()`）更好
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* init 函式說明 */
package main

import "fmt"

func main() {
	fmt.Println("=== init 函式 ===\n")

	fmt.Println("【init 函式語法】")
	fmt.Println("func init() {")
	fmt.Println("    // 初始化程式碼")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("特性：")
	fmt.Println("- 自動執行，無需呼叫")
	fmt.Println("- 不能有參數或回傳值")
	fmt.Println("- 一個檔案可以有多個 init")

	fmt.Println("\n【執行順序】")
	fmt.Println("1. 相依 package 的初始化")
	fmt.Println("2. 當前 package 的全域變數初始化")
	fmt.Println("3. 當前 package 的 init 函式")
	fmt.Println("4. main.main() 函式（如果是 main package）")

	fmt.Println("\n【為什麼要避免】")
	fmt.Println("✗ 無法測試（不能直接呼叫）")
	fmt.Println("✗ 無法傳遞參數")
	fmt.Println("✗ 無法回傳錯誤")
	fmt.Println("✗ 隱藏的副作用")
	fmt.Println("✗ 增加啟動時間")
	fmt.Println("✗ 難以理解程式碼流程")

	fmt.Println("\n【合理使用場景】")
	fmt.Println("✓ 註冊驅動程式")
	fmt.Println("  import _ \"github.com/lib/pq\"")
	fmt.Println("")
	fmt.Println("✓ 驗證編譯時常數")
	fmt.Println("  func init() {")
	fmt.Println("      if maxSize < minSize {")
	fmt.Println("          panic(\"invalid configuration\")")
	fmt.Println("      }")
	fmt.Println("  }")

	fmt.Println("\n【推薦做法：使用明確的初始化】")
	fmt.Println("// 不好：使用 init")
	fmt.Println("var db *sql.DB")
	fmt.Println("func init() {")
	fmt.Println("    db, _ = sql.Open(\"postgres\", connStr)")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("// 好：使用明確函式")
	fmt.Println("func NewDB(connStr string) (*sql.DB, error) {")
	fmt.Println("    return sql.Open(\"postgres\", connStr)")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("優點：")
	fmt.Println("✓ 可以測試")
	fmt.Println("✓ 可以處理錯誤")
	fmt.Println("✓ 明確的控制流程")
	fmt.Println("✓ 可以傳遞設定")
}
// 輸出:
// === init 函式 ===
//
// 【init 函式語法】
// func init() {
//     // 初始化程式碼
// }
//
// 特性：
// - 自動執行，無需呼叫
// - 不能有參數或回傳值
// - 一個檔案可以有多個 init
//
// 【執行順序】
// 1. 相依 package 的初始化
// 2. 當前 package 的全域變數初始化
// 3. 當前 package 的 init 函式
// 4. main.main() 函式（如果是 main package）
//
// 【為什麼要避免】
// ✗ 無法測試（不能直接呼叫）
// ✗ 無法傳遞參數
// ✗ 無法回傳錯誤
// ✗ 隱藏的副作用
// ✗ 增加啟動時間
// ✗ 難以理解程式碼流程
//
// 【合理使用場景】
// ✓ 註冊驅動程式
//   import _ "github.com/lib/pq"
//
// ✓ 驗證編譯時常數
//   func init() {
//       if maxSize < minSize {
//           panic("invalid configuration")
//       }
//   }
//
// 【推薦做法：使用明確的初始化】
// // 不好：使用 init
// var db *sql.DB
// func init() {
//     db, _ = sql.Open("postgres", connStr)
// }
//
// // 好：使用明確函式
// func NewDB(connStr string) (*sql.DB, error) {
//     return sql.Open("postgres", connStr)
// }
//
// 優點：
// ✓ 可以測試
// ✓ 可以處理錯誤
// ✓ 明確的控制流程
// ✓ 可以傳遞設定
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 循環依賴關係

循環依賴（circular dependency）是指兩個或多個 package 互相 import，Go 編譯器會拒絕編譯這樣的程式碼。

**問題**：
- Package A import Package B
- Package B import Package A
- 編譯器報錯：import cycle not allowed

**解決方案**：
1. **引入第三個 package** - 提取共用介面或型態
2. **重新設計架構** - 檢視職責是否劃分正確
3. **使用介面** - 定義介面打破直接相依
4. **合併 package** - 如果真的緊密相關

**預防**：
- 清晰的分層架構
- 相依關係單向流動
- 使用介面降低耦合
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* 循環依賴說明 */
package main

import "fmt"

func main() {
	fmt.Println("=== 循環依賴 ===\n")

	fmt.Println("【問題範例】")
	fmt.Println("// package user")
	fmt.Println("import \"myapp/order\"")
	fmt.Println("func GetUserOrders(id int) []order.Order { ... }")
	fmt.Println("")
	fmt.Println("// package order")
	fmt.Println("import \"myapp/user\"")
	fmt.Println("func CreateOrder(u user.User) { ... }")
	fmt.Println("")
	fmt.Println("錯誤：import cycle: myapp/user -> myapp/order -> myapp/user")

	fmt.Println("\n【解決方案一：引入第三個 package】")
	fmt.Println("myapp/")
	fmt.Println("├── types/")
	fmt.Println("│   ├── user.go        // type User struct")
	fmt.Println("│   └── order.go       // type Order struct")
	fmt.Println("├── user/")
	fmt.Println("│   └── service.go     // import \"myapp/types\"")
	fmt.Println("└── order/")
	fmt.Println("    └── service.go     // import \"myapp/types\"")
	fmt.Println("")
	fmt.Println("說明：將共用型態提取到 types package")

	fmt.Println("\n【解決方案二：使用介面】")
	fmt.Println("// package order")
	fmt.Println("type UserGetter interface {")
	fmt.Println("    GetUser(id int) User")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("func CreateOrder(userGetter UserGetter) { ... }")
	fmt.Println("")
	fmt.Println("說明：order 定義介面，user 實作介面")

	fmt.Println("\n【解決方案三：重新設計】")
	fmt.Println("檢視職責劃分：")
	fmt.Println("- user 負責使用者相關邏輯")
	fmt.Println("- order 負責訂單相關邏輯")
	fmt.Println("- 也許需要 service 層協調兩者")

	fmt.Println("\n【解決方案四：合併 package】")
	fmt.Println("如果 user 和 order 真的緊密相關：")
	fmt.Println("myapp/")
	fmt.Println("└── domain/")
	fmt.Println("    ├── user.go")
	fmt.Println("    └── order.go")

	fmt.Println("\n【預防循環依賴】")
	fmt.Println("✓ 分層架構：上層可以依賴下層，反之不行")
	fmt.Println("✓ 相依方向：domain ← service ← handler")
	fmt.Println("✓ 使用介面：降低耦合")
	fmt.Println("✓ 職責單一：每個 package 有明確職責")
}
// 輸出:
// === 循環依賴 ===
//
// 【問題範例】
// // package user
// import "myapp/order"
// func GetUserOrders(id int) []order.Order { ... }
//
// // package order
// import "myapp/user"
// func CreateOrder(u user.User) { ... }
//
// 錯誤：import cycle: myapp/user -> myapp/order -> myapp/user
//
// 【解決方案一：引入第三個 package】
// myapp/
// ├── types/
// │   ├── user.go        // type User struct
// │   └── order.go       // type Order struct
// ├── user/
// │   └── service.go     // import "myapp/types"
// └── order/
//     └── service.go     // import "myapp/types"
//
// 說明：將共用型態提取到 types package
//
// 【解決方案二：使用介面】
// // package order
// type UserGetter interface {
//     GetUser(id int) User
// }
//
// func CreateOrder(userGetter UserGetter) { ... }
//
// 說明：order 定義介面，user 實作介面
//
// 【解決方案三：重新設計】
// 檢視職責劃分：
// - user 負責使用者相關邏輯
// - order 負責訂單相關邏輯
// - 也許需要 service 層協調兩者
//
// 【解決方案四：合併 package】
// 如果 user 和 order 真的緊密相關：
// myapp/
// └── domain/
//     ├── user.go
//     └── order.go
//
// 【預防循環依賴】
// ✓ 分層架構：上層可以依賴下層，反之不行
// ✓ 相依方向：domain ← service ← handler
// ✓ 使用介面：降低耦合
// ✓ 職責單一：每個 package 有明確職責
```
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
### 優雅地重命名與重組織你的 API

隨著專案演進，可能需要重命名或重組織 package。Go 提供了一些機制讓這個過程更平滑，避免破壞現有使用者的程式碼。

**重構策略**：
1. **使用型態別名** - 保留舊名稱指向新名稱
2. **保留舊 package** - 在舊位置提供轉接層
3. **文件標記** - 使用 `Deprecated:` 註解
4. **版本控制** - 主版本升級表示不相容變更

**型態別名（Type Alias）**：
```go
// 新的型態定義
type NewUser struct { ... }

// 舊的型態別名（標記為已棄用）
// Deprecated: 使用 NewUser 取代
type User = NewUser
```

**最佳實踐**：
- 提供過渡期（至少一個版本）
- 清楚的文件說明
- 使用語意化版本管理
- 考慮使用者的升級成本
<!-- END_MARKDOWN_CELL -->

<!-- CODE_CELL -->
```go
/* API 重構範例 */
package main

import "fmt"

func main() {
	fmt.Println("=== 優雅地重組織 API ===\n")

	fmt.Println("【場景：重命名 package】")
	fmt.Println("舊結構：")
	fmt.Println("myproject/")
	fmt.Println("└── utils/")
	fmt.Println("    └── string.go")
	fmt.Println("")
	fmt.Println("新結構：")
	fmt.Println("myproject/")
	fmt.Println("├── stringutil/")
	fmt.Println("│   └── string.go      // 新的實作")
	fmt.Println("└── utils/")
	fmt.Println("    └── string.go      // 轉接到 stringutil")

	fmt.Println("\n【舊 package 的轉接實作】")
	fmt.Println("// utils/string.go")
	fmt.Println("package utils")
	fmt.Println("")
	fmt.Println("import \"myproject/stringutil\"")
	fmt.Println("")
	fmt.Println("// Reverse 反轉字串")
	fmt.Println("// Deprecated: 使用 stringutil.Reverse 取代")
	fmt.Println("func Reverse(s string) string {")
	fmt.Println("    return stringutil.Reverse(s)")
	fmt.Println("}")

	fmt.Println("\n【型態別名範例】")
	fmt.Println("// 新的定義")
	fmt.Println("type Configuration struct {")
	fmt.Println("    Host string")
	fmt.Println("    Port int")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("// 舊的別名")
	fmt.Println("// Deprecated: 使用 Configuration 取代")
	fmt.Println("type Config = Configuration")

	fmt.Println("\n【版本升級策略】")
	fmt.Println("v1.x.x → v1.y.y：相容變更")
	fmt.Println("  ✓ 新增函式/方法")
	fmt.Println("  ✓ 新增欄位（使用預設值）")
	fmt.Println("  ✓ 標記為 Deprecated")
	fmt.Println("")
	fmt.Println("v1.x.x → v2.0.0：不相容變更")
	fmt.Println("  • 移除已棄用的功能")
	fmt.Println("  • 修改函式簽名")
	fmt.Println("  • 重組 package 結構")
	fmt.Println("  • 需要修改 module 路徑（加上 /v2）")

	fmt.Println("\n【遷移期間的最佳實踐】")
	fmt.Println("1. 先新增新 API，保留舊 API")
	fmt.Println("2. 標記舊 API 為 Deprecated")
	fmt.Println("3. 提供遷移文件")
	fmt.Println("4. 至少維護一個版本週期")
	fmt.Println("5. 在主版本升級時移除舊 API")
}
// 輸出:
// === 優雅地重組織 API ===
//
// 【場景：重命名 package】
// 舊結構：
// myproject/
// └── utils/
//     └── string.go
//
// 新結構：
// myproject/
// ├── stringutil/
// │   └── string.go      // 新的實作
// └── utils/
//     └── string.go      // 轉接到 stringutil
//
// 【舊 package 的轉接實作】
// // utils/string.go
// package utils
//
// import "myproject/stringutil"
//
// // Reverse 反轉字串
// // Deprecated: 使用 stringutil.Reverse 取代
// func Reverse(s string) string {
//     return stringutil.Reverse(s)
// }
//
// 【型態別名範例】
// // 新的定義
// type Configuration struct {
//     Host string
//     Port int
// }
//
// // 舊的別名
// // Deprecated: 使用 Configuration 取代
// type Config = Configuration
//
// 【版本升級策略】
// v1.x.x → v1.y.y：相容變更
//   ✓ 新增函式/方法
//   ✓ 新增欄位（使用預設值）
//   ✓ 標記為 Deprecated
//
// v1.x.x → v2.0.0：不相容變更
//   • 移除已棄用的功能
//   • 修改函式簽名
//   • 重組 package 結構
//   • 需要修改 module 路徑（加上 /v2）
//
// 【遷移期間的最佳實踐】
// 1. 先新增新 API，保留舊 API
// 2. 標記舊 API 為 Deprecated
// 3. 提供遷移文件
// 4. 至少維護一個版本週期
// 5. 在主版本升級時移除舊 API
```
<!-- END_CODE_CELL -->
