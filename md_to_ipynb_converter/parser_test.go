package main

import (
	"strings"
	"testing"
)

func TestParser_ParseSimpleMarkdown(t *testing.T) {
	input := `<!-- MARKDOWN_CELL -->
# Test Title

This is a test.

<!-- CODE_CELL -->
` + "```go" + `
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
` + "```" + `
<!-- END_CODE_CELL -->`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 2 {
		t.Errorf("Expected 2 cells, got %d", len(notebook.Cells))
	}

	// 檢查第一個 cell (markdown)
	if notebook.Cells[0].CellType != "markdown" {
		t.Errorf("First cell should be markdown, got %s", notebook.Cells[0].CellType)
	}

	if notebook.Cells[0].ID != "cell-0" {
		t.Errorf("First cell ID should be cell-0, got %s", notebook.Cells[0].ID)
	}

	// 檢查第二個 cell (code)
	if notebook.Cells[1].CellType != "code" {
		t.Errorf("Second cell should be code, got %s", notebook.Cells[1].CellType)
	}

	if notebook.Cells[1].ID != "cell-1" {
		t.Errorf("Second cell ID should be cell-1, got %s", notebook.Cells[1].ID)
	}
}

func TestParser_ParseEmptyCells(t *testing.T) {
	input := `<!-- MARKDOWN_CELL -->

<!-- CODE_CELL -->
` + "```go" + `
` + "```" + `
<!-- END_CODE_CELL -->`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// 空的 cells 應該被忽略
	if len(notebook.Cells) != 0 {
		t.Errorf("Expected 0 cells, got %d", len(notebook.Cells))
	}
}

func TestParser_ParseMultipleCells(t *testing.T) {
	input := `<!-- MARKDOWN_CELL -->
## Section 1

<!-- CODE_CELL -->
` + "```go" + `
var x = 1
` + "```" + `
<!-- END_CODE_CELL -->

<!-- MARKDOWN_CELL -->
## Section 2

<!-- CODE_CELL -->
` + "```go" + `
var y = 2
` + "```" + `
<!-- END_CODE_CELL -->`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 4 {
		t.Errorf("Expected 4 cells, got %d", len(notebook.Cells))
	}

	expectedTypes := []string{"markdown", "code", "markdown", "code"}
	for i, cell := range notebook.Cells {
		if cell.CellType != expectedTypes[i] {
			t.Errorf("Cell %d: expected type %s, got %s", i, expectedTypes[i], cell.CellType)
		}
	}
}

func TestParser_RemovesCodeFences(t *testing.T) {
	input := `<!-- CODE_CELL -->
` + "```go" + `
package main
func main() {}
` + "```" + `
<!-- END_CODE_CELL -->`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 1 {
		t.Fatalf("Expected 1 cell, got %d", len(notebook.Cells))
	}

	// 檢查是否移除了 ```go 和 ```
	content := strings.Join(notebook.Cells[0].Source, "")
	if strings.Contains(content, "```go") {
		t.Error("Code cell should not contain ```go")
	}
	if strings.Contains(content, "```") {
		t.Error("Code cell should not contain closing ```")
	}
	if !strings.Contains(content, "package main") {
		t.Error("Code cell should contain actual code")
	}
}

func TestParser_PreservesMarkdownContent(t *testing.T) {
	input := `<!-- MARKDOWN_CELL -->
# Title

This is **bold** and *italic*.

- List item 1
- List item 2

Inline ` + "`code`" + ` example.`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 1 {
		t.Fatalf("Expected 1 cell, got %d", len(notebook.Cells))
	}

	content := strings.Join(notebook.Cells[0].Source, "")
	if !strings.Contains(content, "# Title") {
		t.Error("Should preserve title")
	}
	if !strings.Contains(content, "**bold**") {
		t.Error("Should preserve bold formatting")
	}
	if !strings.Contains(content, "- List item 1") {
		t.Error("Should preserve list items")
	}
}

func TestParser_HandleMultilineCode(t *testing.T) {
	input := `<!-- CODE_CELL -->
` + "```go" + `
package main

import (
    "fmt"
    "strings"
)

func main() {
    x := 1
    y := 2
    fmt.Println(x + y)
}
` + "```" + `
<!-- END_CODE_CELL -->`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 1 {
		t.Fatalf("Expected 1 cell, got %d", len(notebook.Cells))
	}

	content := strings.Join(notebook.Cells[0].Source, "")
	if !strings.Contains(content, "package main") {
		t.Error("Should contain package declaration")
	}
	if !strings.Contains(content, "import") {
		t.Error("Should contain import statement")
	}
	if !strings.Contains(content, "func main()") {
		t.Error("Should contain main function")
	}
}

func TestParser_CellIDIncrement(t *testing.T) {
	input := `<!-- MARKDOWN_CELL -->
First

<!-- MARKDOWN_CELL -->
Second

<!-- MARKDOWN_CELL -->
Third`

	parser := NewParser(strings.NewReader(input))
	notebook, err := parser.Parse()

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(notebook.Cells) != 3 {
		t.Fatalf("Expected 3 cells, got %d", len(notebook.Cells))
	}

	expectedIDs := []string{"cell-0", "cell-1", "cell-2"}
	for i, cell := range notebook.Cells {
		if cell.ID != expectedIDs[i] {
			t.Errorf("Cell %d: expected ID %s, got %s", i, expectedIDs[i], cell.ID)
		}
	}
}
