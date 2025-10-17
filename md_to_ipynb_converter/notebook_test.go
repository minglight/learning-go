package main

import (
	"encoding/json"
	"testing"
)

func TestNewNotebook(t *testing.T) {
	nb := NewNotebook()

	if nb.NBFormat != 4 {
		t.Errorf("Expected nbformat 4, got %d", nb.NBFormat)
	}

	if nb.NBFormatMinor != 4 {
		t.Errorf("Expected nbformat_minor 4, got %d", nb.NBFormatMinor)
	}

	if nb.Metadata.Kernelspec.Name != "gophernotes" {
		t.Errorf("Expected gophernotes kernel, got %s", nb.Metadata.Kernelspec.Name)
	}

	if nb.Metadata.Kernelspec.Language != "go" {
		t.Errorf("Expected go language, got %s", nb.Metadata.Kernelspec.Language)
	}

	if len(nb.Cells) != 0 {
		t.Errorf("Expected empty cells, got %d cells", len(nb.Cells))
	}
}

func TestNotebook_AddMarkdownCell(t *testing.T) {
	nb := NewNotebook()
	nb.AddMarkdownCell("test-id", "# Title\n\nContent")

	if len(nb.Cells) != 1 {
		t.Fatalf("Expected 1 cell, got %d", len(nb.Cells))
	}

	cell := nb.Cells[0]
	if cell.CellType != "markdown" {
		t.Errorf("Expected markdown cell, got %s", cell.CellType)
	}

	if cell.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", cell.ID)
	}

	if cell.ExecutionCount != nil {
		t.Errorf("Markdown cell should not have ExecutionCount")
	}

	if cell.Outputs != nil {
		t.Errorf("Markdown cell should not have Outputs")
	}
}

func TestNotebook_AddCodeCell(t *testing.T) {
	nb := NewNotebook()
	nb.AddCodeCell("code-id", "package main\n\nfunc main() {}")

	if len(nb.Cells) != 1 {
		t.Fatalf("Expected 1 cell, got %d", len(nb.Cells))
	}

	cell := nb.Cells[0]
	if cell.CellType != "code" {
		t.Errorf("Expected code cell, got %s", cell.CellType)
	}

	if cell.ExecutionCount != nil {
		t.Errorf("ExecutionCount should be nil")
	}

	if cell.Outputs == nil {
		t.Errorf("Outputs should not be nil")
	}

	if len(cell.Outputs) != 0 {
		t.Errorf("Outputs should be empty array, got %d items", len(cell.Outputs))
	}
}

func TestNotebook_ToJSON(t *testing.T) {
	nb := NewNotebook()
	nb.AddMarkdownCell("cell-0", "# Test")

	jsonData, err := nb.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	// 驗證 JSON 格式
	var parsed map[string]any
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v", err)
	}

	// 檢查必要欄位
	if _, ok := parsed["cells"]; !ok {
		t.Error("Missing 'cells' field")
	}

	if _, ok := parsed["metadata"]; !ok {
		t.Error("Missing 'metadata' field")
	}

	if _, ok := parsed["nbformat"]; !ok {
		t.Error("Missing 'nbformat' field")
	}

	if _, ok := parsed["nbformat_minor"]; !ok {
		t.Error("Missing 'nbformat_minor' field")
	}
}

func TestSplitLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single line without newline",
			input:    "single line",
			expected: []string{"single line"},
		},
		{
			name:     "single line with newline",
			input:    "single line\n",
			expected: []string{"single line\n"},
		},
		{
			name:     "two lines",
			input:    "line1\nline2",
			expected: []string{"line1\n", "line2"},
		},
		{
			name:     "two lines with trailing newline",
			input:    "line1\nline2\n",
			expected: []string{"line1\n", "line2\n"},
		},
		{
			name:     "multiple lines",
			input:    "line1\nline2\nline3",
			expected: []string{"line1\n", "line2\n", "line3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitLines(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("splitLines(%q) length = %d, want %d", tt.input, len(result), len(tt.expected))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("splitLines(%q)[%d] = %q, want %q", tt.input, i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestNotebook_MultipleCells(t *testing.T) {
	nb := NewNotebook()
	nb.AddMarkdownCell("md-1", "# Title")
	nb.AddCodeCell("code-1", "var x = 1")
	nb.AddMarkdownCell("md-2", "## Section")
	nb.AddCodeCell("code-2", "var y = 2")

	if len(nb.Cells) != 4 {
		t.Errorf("Expected 4 cells, got %d", len(nb.Cells))
	}

	expectedTypes := []string{"markdown", "code", "markdown", "code"}
	expectedIDs := []string{"md-1", "code-1", "md-2", "code-2"}

	for i, cell := range nb.Cells {
		if cell.CellType != expectedTypes[i] {
			t.Errorf("Cell %d: expected type %s, got %s", i, expectedTypes[i], cell.CellType)
		}
		if cell.ID != expectedIDs[i] {
			t.Errorf("Cell %d: expected ID %s, got %s", i, expectedIDs[i], cell.ID)
		}
	}
}
