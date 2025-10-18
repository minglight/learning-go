package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s input.md output.ipynb\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if err := convert(inputFile, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 成功轉換: %s -> %s\n", inputFile, outputFile)
}

func convert(inputPath, outputPath string) error {
	// 讀取輸入檔案
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	// 解析
	parser := NewParser(inputFile)
	notebook, err := parser.Parse()
	if err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	// 轉換成 JSON
	jsonData, err := notebook.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	// 寫入輸出檔案
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("📊 共 %d 個 cells\n", len(notebook.Cells))

	return nil
}
