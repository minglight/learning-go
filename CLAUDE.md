# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go language learning repository that uses Jupyter Notebooks with GoNB (Go Notebooks) for interactive learning. The repository contains structured learning materials covering Go fundamentals from basic types to advanced concepts like pointers, interfaces, and error handling.

## Architecture & Structure

### Learning Materials Organization
- **Chapter-based structure**: Each chapter (`ch1_setup_env/`, `ch2/`, `ch3/`, etc.) contains specific Go concepts
- **Jupyter Notebooks**: Primary learning medium using `.ipynb` files with GoNB kernel
- **Supplementary materials**: ChatGPT-generated examples in `chatgpt/` directory

### Key Components
- **Interactive Learning**: Jupyter notebooks allow code execution and immediate feedback
- **Bilingual Content**: Materials are primarily in Traditional Chinese with English code comments
- **Comprehensive Coverage**: From basic types (ch2) to composite types (ch3), control structures (ch4), functions (ch5), and pointers (ch6)

## Development Environment

### GoNB Jupyter Setup ( already set )
Based on `USE_JUPYTER_FOR_GO.md`, this repository uses GoNB for Go development in Jupyter:

```bash
# Install GoNB
go install github.com/janpfeifer/gonb@latest

# Install kernel to Jupyter
gonb --install

# Start Jupyter Lab
jupyter lab
```

### IDE Integration
- **Cursor**: Jupyter extension support for `.ipynb` files
- **VS Code**: Alternative IDE with Jupyter support

##
Key Reference files 
./目錄.md : 放著Learning go的每個章節的目錄

## Common Commands
- 我會要求你分章節寫筆記, 你要幫我放在每個章節的chapter下, 用.ipynb的格式output出來

- 當我要求的時候, 你要幫我看完我的內容跟大綱之後, 在同一個chapter folder下另外開一個jupyer file, 並且要寫一些面試常出現的考題, 你使用markdown寫題目, 題目要有標題跟內容, 然後要留 code block讓我寫答案. 

- 當我寫完之後我會請你改考卷, 你要另外開一個jupyter file, copy我寫不好或寫錯的那一題的題目跟我寫的答案, 然後新增你的說明跟正確解答 




### Learning Progression
1. **Basic Types** (ch2): Variables, constants, type conversion, rune handling
2. **Composite Types** (ch3): Arrays, slices, maps, structs
3. **Control Structures** (ch4): Loops, conditionals, switch statements
4. **Functions** (ch5): Function declarations, multiple returns, closures
5. **Pointers** (ch6): Memory management, pointer operations

### Jupyter Notebook Conventions
- Markdown cells for explanations in Chinese
- Code cells with complete, runnable Go examples
- Output cells showing execution results
- Educational structure with step-by-step progression

## Testing & Validation

Since this is a learning repository without formal tests:
- Verify code by running individual Go files
- Use Jupyter notebook execution to validate examples
- Cross-reference with Go official documentation

## Important Notes

- **No build system**: This is a learning repository without Make files or Go modules
- **Interactive focus**: Designed for experiential learning through Jupyter notebooks
- **Chinese documentation**: Primary language for explanations and comments
- **GoNB dependency**: Requires GoNB installation for full functionality
