# SSG
This Go-based tool is a lightweight static site generator designed to create and build daily blog posts or notes. It automates the process of generating Markdown files and converting them into structured HTML using sections.

---

## 🚀 Features

* **Quick Scaffolding**: Generate a new Markdown file named after the current date (`YYYY-MM-DD.md`) with a single command.
* **Section-Based Parsing**: Automatically splits Markdown content into multiple sections using the `---` horizontal rule delimiter.
* **Template Integration**: Injects parsed content and metadata (like the date) into a `layout.html` file using Go's `html/template` engine.
* **Markdown Support**: Powered by the `goldmark` library for fast and compliant Markdown-to-HTML conversion.

---

## 🛠️ Installation

1.  Ensure you have **Go** installed on your system.
2.  Install the required dependency:
    ```bash
    go get github.com/yuin/goldmark
    ```

---

## 📖 Usage

The tool operates using two primary commands: `new` and `build`.

### 1. Create a New Post
To create a blank Markdown file for today:
```bash
go run main.go new
```
* **Result**: Creates a file named `2026-04-25.md` (based on the current date).

### 2. Build the HTML Page
To convert your Markdown content into an HTML file:
```bash
go run main.go build
```
* **Input**: Reads `YYYY-MM-DD.md`.
* **Processing**: Each section separated by `---` is treated as an individual piece of HTML.
* **Output**: Generates `YYYY-MM-DD.html` using the structure defined in `layout.html`.

---

## 📄 File Structure

* **`main.go`**: The core logic for file management and Markdown processing.
* **`layout.html`**: The required base template. It must handle the `PageData` struct:
    * `.Date`: The post date string.
    * `.Contents`: A slice of HTML sections.
* **`YYYY-MM-DD.md`**: Your content source.

---

## 🧩 Data Structure

The tool passes the following data to your `layout.html`:

```go
type PageData struct {
    Date     string          // e.g., "2026-04-25"
    Contents []template.HTML // Slice containing HTML for each section
}
```