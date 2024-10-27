package command

import (
	"bufio"
	ccwcerr "ccwc/error"
	"fmt"
	"os"
	"strings"
)

type (
	Wc interface {
		getNumberOfLines() int
		getNumberOfWords() int
		getNumberOfBytes() int
		getNumberOfCharacters() int
		Exec()
	}

	wcFromFile struct {
		options  string
		filePath string
	}
)

func NewWcFromFile(options string, filePath string) Wc {
	return &wcFromFile{options, filePath}
}

func (w *wcFromFile) getNumberOfLines() int {
	open, err := getInputFileOrStdIn(w.filePath)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	fileScanner := bufio.NewScanner(open)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

func (w *wcFromFile) getNumberOfWords() int {
	open, err := getInputFileOrStdIn(w.filePath)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	fileScanner := bufio.NewScanner(open)
	fileScanner.Split(bufio.ScanWords)
	wordCount := 0
	for fileScanner.Scan() {
		wordCount++
	}
	return wordCount
}

func (w *wcFromFile) getNumberOfBytes() int {
	open, err := getInputFileOrStdIn(w.filePath)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	fileScanner := bufio.NewScanner(open)
	fileScanner.Split(bufio.ScanBytes)
	byteCount := 0
	for fileScanner.Scan() {
		byteCount++
	}
	return byteCount
}

func (w *wcFromFile) getNumberOfCharacters() int {
	open, err := getInputFileOrStdIn(w.filePath)
	if err != nil {
		panic(ccwcerr.NewFileOperationIssue(err))
	}
	defer open.Close()
	fileScanner := bufio.NewScanner(open)
	fileScanner.Split(bufio.ScanRunes)
	byteCount := 0
	for fileScanner.Scan() {
		byteCount++
	}
	return byteCount
}

func (w *wcFromFile) Exec() {
	fmt.Println(w.filePath)
	if len(w.options) == 0 || strings.Contains(w.options, "c") {
		fmt.Printf("# Bytes: %d \t", w.getNumberOfBytes())
	}
	if len(w.options) == 0 || strings.Contains(w.options, "l") {
		fmt.Printf("# Lines: %d \t", w.getNumberOfLines())
	}
	if len(w.options) == 0 || strings.Contains(w.options, "w") {
		fmt.Printf("# Words: %d \t", w.getNumberOfWords())
	}
	if len(w.options) == 0 || strings.Contains(w.options, "m") {
		fmt.Printf("# Characters: %d \t", w.getNumberOfCharacters())
	}
	fmt.Println()
}

func getInputFileOrStdIn(filePath string) (*os.File, error) {
	if filePath == "" {
		return os.Stdin, nil
	}
	open, err := os.Open(filePath)
	return open, err
}
