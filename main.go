package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/yuin/goldmark"
)

type PageData struct {
	Date    string
	Content template.HTML
}

func main() {
	// 1. check argument
	if len(os.Args) < 2 {
		fmt.Println("사용법: go run main.go [new|build]")
		return
	}

	command := os.Args[1]
	today := time.Now().Format("2006-01-02") // date (YYYY-MM-DD)
	mdFileName := "test/" + today + ".md"
	htmlFileName := "test/" + today + ".html"

	switch command {
	case "new":
		createNewPost(mdFileName, today)
	case "build":
		buildPost(mdFileName, htmlFileName, today)
	default:
		fmt.Printf("unknown command: %s (use new or build)\n", command)
	}
}

// new: create new markdown file
func createNewPost(fileName string, date string) {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("file already exist: %s\n", fileName)
		return
	}

	initialContent := ""
	err := os.WriteFile(fileName, []byte(initialContent), 0644)
	if err != nil {
		fmt.Println("failed to create new file:", err)
		return
	}
	fmt.Printf("new file was created: %s\n", fileName)
}

// build: turn makrdown to html
func buildPost(mdFile string, htmlFile string, date string) {
	// read markdown
	source, err := os.ReadFile(mdFile)
	if err != nil {
		fmt.Printf("failed to find markdown file: %s\n", mdFile)
		return
	}

	// Goldmark 
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		fmt.Println("failed:", err)
		return
	}

	// save html
	outFile, _ := os.Create(htmlFile)
	defer outFile.Close()

	data := PageData{
		Date:    date,
		Content: template.HTML(buf.String()),
	}

	tmpl.Execute(outFile, data)
	fmt.Printf("success: %s -> %s\n", mdFile, htmlFile)
}