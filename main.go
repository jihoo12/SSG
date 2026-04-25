package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type PageData struct {
	Date    string
	Contents []template.HTML
}

func main() {
	// 1. check argument
	if len(os.Args) < 2 {
		fmt.Println("need argument new or build")
		return
	}

	command := os.Args[1]
	today := time.Now().Format("2006-01-02") // date (YYYY-MM-DD)
	mdFileName := today + ".md"
	htmlFileName := today + ".html"

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
    source, err := os.ReadFile(mdFile)
    if err != nil {
        fmt.Printf("❌ Failed to read markdown file: %s\n", mdFile)
        return
    }

    md := goldmark.New()
    var htmlSections []template.HTML

    // Convert markdown source to string
    contentStr := string(source)

    // Split content into sections using "---" as a delimiter.
    // If no delimiter is found, the entire file is treated as a single section.
    var rawSections []string
    if strings.Contains(contentStr, "\n---\n") {
        rawSections = strings.Split(contentStr, "\n---\n")
    } else {
        rawSections = []string{contentStr}
    }

    for i, section := range rawSections {
        // Skip empty sections or sections containing only whitespace
        if strings.TrimSpace(section) == "" {
            continue
        }

        var buf bytes.Buffer
        if err := md.Convert([]byte(section), &buf); err != nil {
            fmt.Printf("❌ Error converting section %d: %v\n", i, err)
            continue
        }
        htmlSections = append(htmlSections, template.HTML(buf.String()))
    }

    // Ensure the template file "layout.html" exists in the current directory
    tmpl, err := template.ParseFiles("layout.html")
    if err != nil {
        fmt.Printf("❌ Failed to load template: %v\n", err)
        return
    }

    outFile, err := os.Create(htmlFile)
    if err != nil {
        fmt.Printf("❌ Failed to create output file: %v\n", err)
        return
    }
    defer outFile.Close()

    data := PageData{
        Date:     date,
        Contents: htmlSections,
    }

    // Execute template and inject data; log any execution errors to the terminal
    err = tmpl.Execute(outFile, data)
    if err != nil {
        fmt.Printf("❌ Failed to execute template (data injection): %v\n", err)
        return
    }

    fmt.Printf("✅ Build successful: %d sections generated -> %s\n", len(htmlSections), htmlFile)
}