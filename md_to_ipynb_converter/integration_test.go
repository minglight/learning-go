package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestIntegration_SimpleConversion(t *testing.T) {
	// 準備測試資料目錄
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "test.md")
	outputPath := filepath.Join(testDir, "test.ipynb")

	// 創建測試輸入
	testInput := `<!-- MARKDOWN_CELL -->
# 第九章 模組、程式包與匯入

本章學習重點

<!-- CODE_CELL -->
` + "```go" + `
/* 示範範例 */
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
// 輸出: Hello, Go!
` + "```" + `
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 版本庫、模組與程式包

這是說明內容。`

	if err := os.WriteFile(inputPath, []byte(testInput), 0644); err != nil {
		t.Fatalf("Failed to create test input: %v", err)
	}

	// 執行轉換
	if err := convert(inputPath, outputPath); err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	// 驗證輸出檔案存在
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatal("Output file was not created")
	}

	// 讀取並驗證 JSON 格式
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	// 驗證結構
	if len(notebook.Cells) != 3 {
		t.Errorf("Expected 3 cells, got %d", len(notebook.Cells))
	}

	if notebook.NBFormat != 4 {
		t.Errorf("Expected nbformat 4, got %d", notebook.NBFormat)
	}

	// 驗證 cell 類型順序
	expectedTypes := []string{"markdown", "code", "markdown"}
	for i, expected := range expectedTypes {
		if i < len(notebook.Cells) && notebook.Cells[i].CellType != expected {
			t.Errorf("Cell %d: expected type %s, got %s", i, expected, notebook.Cells[i].CellType)
		}
	}
}

func TestIntegration_ComplexDocument(t *testing.T) {
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "complex.md")
	outputPath := filepath.Join(testDir, "complex.ipynb")

	// 創建較複雜的測試文件
	complexInput := `<!-- MARKDOWN_CELL -->
# 章節標題

多行內容
包含列表：
- 項目 1
- 項目 2

**粗體文字**

<!-- CODE_CELL -->
` + "```go" + `
package main

import (
    "fmt"
    "time"
)

func main() {
    // 多行程式碼
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}
` + "```" + `
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 小節 1

內容 1

<!-- CODE_CELL -->
` + "```go" + `
var x = 1
` + "```" + `
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## 小節 2

內容 2`

	if err := os.WriteFile(inputPath, []byte(complexInput), 0644); err != nil {
		t.Fatalf("Failed to create test input: %v", err)
	}

	if err := convert(inputPath, outputPath); err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if len(notebook.Cells) != 5 {
		t.Errorf("Expected 5 cells, got %d", len(notebook.Cells))
	}
}

func TestIntegration_EmptyInput(t *testing.T) {
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "empty.md")
	outputPath := filepath.Join(testDir, "empty.ipynb")

	// 創建空文件
	if err := os.WriteFile(inputPath, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create test input: %v", err)
	}

	if err := convert(inputPath, outputPath); err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if len(notebook.Cells) != 0 {
		t.Errorf("Expected 0 cells for empty input, got %d", len(notebook.Cells))
	}
}

func TestIntegration_OnlyMarkdownCells(t *testing.T) {
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "markdown.md")
	outputPath := filepath.Join(testDir, "markdown.ipynb")

	markdownInput := `<!-- MARKDOWN_CELL -->
# Title 1

Content 1

<!-- MARKDOWN_CELL -->
# Title 2

Content 2`

	if err := os.WriteFile(inputPath, []byte(markdownInput), 0644); err != nil {
		t.Fatalf("Failed to create test input: %v", err)
	}

	if err := convert(inputPath, outputPath); err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if len(notebook.Cells) != 2 {
		t.Errorf("Expected 2 cells, got %d", len(notebook.Cells))
	}

	for i, cell := range notebook.Cells {
		if cell.CellType != "markdown" {
			t.Errorf("Cell %d should be markdown, got %s", i, cell.CellType)
		}
	}
}

func TestIntegration_OnlyCodeCells(t *testing.T) {
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "code.md")
	outputPath := filepath.Join(testDir, "code.ipynb")

	codeInput := `<!-- CODE_CELL -->
` + "```go" + `
var x = 1
` + "```" + `
<!-- END_CODE_CELL -->

<!-- CODE_CELL -->
` + "```go" + `
var y = 2
` + "```" + `
<!-- END_CODE_CELL -->`

	if err := os.WriteFile(inputPath, []byte(codeInput), 0644); err != nil {
		t.Fatalf("Failed to create test input: %v", err)
	}

	if err := convert(inputPath, outputPath); err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if len(notebook.Cells) != 2 {
		t.Errorf("Expected 2 cells, got %d", len(notebook.Cells))
	}

	for i, cell := range notebook.Cells {
		if cell.CellType != "code" {
			t.Errorf("Cell %d should be code, got %s", i, cell.CellType)
		}
		if cell.ExecutionCount != nil {
			t.Errorf("Cell %d should have nil execution count", i)
		}
	}
}

func TestIntegration_NonexistentInputFile(t *testing.T) {
	testDir := t.TempDir()
	inputPath := filepath.Join(testDir, "nonexistent.md")
	outputPath := filepath.Join(testDir, "output.ipynb")

	err := convert(inputPath, outputPath)
	if err == nil {
		t.Error("Expected error for nonexistent input file, got nil")
	}
}
