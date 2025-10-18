package main

import "encoding/json"

// Notebook 代表 Jupyter Notebook 的完整結構
type Notebook struct {
	Cells         []Cell           `json:"cells"`
	Metadata      NotebookMetadata `json:"metadata"`
	NBFormat      int              `json:"nbformat"`
	NBFormatMinor int              `json:"nbformat_minor"`
}

// Cell 代表單一 cell
type Cell struct {
	CellType       string       `json:"cell_type"`
	ID             string       `json:"id"`
	Metadata       CellMetadata `json:"metadata"`
	Source         []string     `json:"source"`
	ExecutionCount *int         `json:"execution_count,omitempty"`
	Outputs        []any        `json:"outputs,omitempty"`
}

// CellMetadata cell 的 metadata（目前為空）
type CellMetadata struct{}

// NotebookMetadata notebook 的 metadata
type NotebookMetadata struct {
	Kernelspec   Kernelspec   `json:"kernelspec"`
	LanguageInfo LanguageInfo `json:"language_info"`
}

// Kernelspec kernel 設定
type Kernelspec struct {
	DisplayName string `json:"display_name"`
	Language    string `json:"language"`
	Name        string `json:"name"`
}

// LanguageInfo 語言資訊
type LanguageInfo struct {
	FileExtension string `json:"file_extension"`
	MimeType      string `json:"mimetype"`
	Name          string `json:"name"`
}

// NewNotebook 創建新的 Go Notebook
func NewNotebook() *Notebook {
	return &Notebook{
		Cells: []Cell{},
		Metadata: NotebookMetadata{
			Kernelspec: Kernelspec{
				DisplayName: "Go",
				Language:    "go",
				Name:        "gophernotes",
			},
			LanguageInfo: LanguageInfo{
				FileExtension: ".go",
				MimeType:      "text/x-go",
				Name:          "go",
			},
		},
		NBFormat:      4,
		NBFormatMinor: 4,
	}
}

// AddMarkdownCell 新增 markdown cell
func (nb *Notebook) AddMarkdownCell(id, content string) {
	cell := Cell{
		CellType: "markdown",
		ID:       id,
		Metadata: CellMetadata{},
		Source:   splitLines(content),
	}
	nb.Cells = append(nb.Cells, cell)
}

// AddCodeCell 新增 code cell
func (nb *Notebook) AddCodeCell(id, content string) {
	cell := Cell{
		CellType:       "code",
		ID:             id,
		Metadata:       CellMetadata{},
		Source:         splitLines(content),
		ExecutionCount: nil,
		Outputs:        []any{},
	}
	nb.Cells = append(nb.Cells, cell)
}

// ToJSON 輸出 JSON 格式
func (nb *Notebook) ToJSON() ([]byte, error) {
	return json.MarshalIndent(nb, "", "  ")
}

// splitLines 將字串分割成行（保留換行符）
func splitLines(content string) []string {
	if content == "" {
		return []string{}
	}

	lines := []string{}
	current := ""

	for i, ch := range content {
		current += string(ch)
		if ch == '\n' {
			lines = append(lines, current)
			current = ""
		} else if i == len(content)-1 {
			// 最後一行沒有換行符
			lines = append(lines, current)
		}
	}

	return lines
}
