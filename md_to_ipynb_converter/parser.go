package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// CellType 定義 cell 類型
type CellType int

const (
	Unknown CellType = iota
	MarkdownCell
	CodeCell
)

// Parser Markdown 解析器
type Parser struct {
	scanner *bufio.Scanner
	cellID  int
}

// NewParser 創建新的解析器
func NewParser(r io.Reader) *Parser {
	return &Parser{
		scanner: bufio.NewScanner(r),
		cellID:  0,
	}
}

// Parse 解析 markdown 並返回 Notebook
func (p *Parser) Parse() (*Notebook, error) {
	notebook := NewNotebook()

	var currentType CellType
	var currentContent strings.Builder

	codeFenceRegex := regexp.MustCompile("^```go\\s*$")
	endCodeFenceRegex := regexp.MustCompile("^```\\s*$")
	inCodeFence := false

	for p.scanner.Scan() {
		line := p.scanner.Text()

		// 檢查標記
		if strings.TrimSpace(line) == "<!-- MARKDOWN_CELL -->" {
			// 儲存前一個 cell（如果有的話）
			if err := p.saveCell(notebook, currentType, currentContent.String()); err != nil {
				return nil, err
			}
			currentType = MarkdownCell
			currentContent.Reset()
			continue
		}

		if strings.TrimSpace(line) == "<!-- END_MARKDOWN_CELL -->" {
			// 儲存 markdown cell
			if err := p.saveCell(notebook, currentType, currentContent.String()); err != nil {
				return nil, err
			}
			currentType = Unknown
			currentContent.Reset()
			continue
		}

		if strings.TrimSpace(line) == "<!-- CODE_CELL -->" {
			// 儲存前一個 cell（如果有的話）
			if err := p.saveCell(notebook, currentType, currentContent.String()); err != nil {
				return nil, err
			}
			currentType = CodeCell
			currentContent.Reset()
			inCodeFence = false
			continue
		}

		if strings.TrimSpace(line) == "<!-- END_CODE_CELL -->" {
			// 儲存 code cell
			if err := p.saveCell(notebook, currentType, currentContent.String()); err != nil {
				return nil, err
			}
			currentType = Unknown
			currentContent.Reset()
			inCodeFence = false
			continue
		}

		// 處理 code fence
		if currentType == CodeCell {
			if codeFenceRegex.MatchString(line) {
				inCodeFence = true
				continue // 跳過 ```go 這一行
			}
			if endCodeFenceRegex.MatchString(line) && inCodeFence {
				inCodeFence = false
				continue // 跳過 ``` 這一行
			}
		}

		// 累積內容
		if currentType != Unknown {
			if currentContent.Len() > 0 {
				currentContent.WriteString("\n")
			}
			currentContent.WriteString(line)
		}
	}

	// 處理最後一個 cell
	if err := p.saveCell(notebook, currentType, currentContent.String()); err != nil {
		return nil, err
	}

	if err := p.scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return notebook, nil
}

// saveCell 儲存當前 cell 到 notebook
func (p *Parser) saveCell(notebook *Notebook, cellType CellType, content string) error {
	if cellType == Unknown || strings.TrimSpace(content) == "" {
		return nil
	}

	cellID := fmt.Sprintf("cell-%d", p.cellID)
	p.cellID++

	switch cellType {
	case MarkdownCell:
		notebook.AddMarkdownCell(cellID, content)
	case CodeCell:
		notebook.AddCodeCell(cellID, content)
	default:
		return fmt.Errorf("unknown cell type: %d", cellType)
	}

	return nil
}
