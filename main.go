package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/russross/blackfriday"
)

const HEADERDELIMITER = "---"
const INPUT_FILE_EXTENSION = ".md"
const OUTPUT_FILE_EXTENSION = ".html"

type Headers struct {
	Title  string
	Author string
	Date   time.Time
	Tags   []string
}

type Post struct {
	Headers Headers
	Body    string
}

func ParsePostFile(file []byte) Post {
	fileParts := bytes.Split(file, []byte(HEADERDELIMITER))

	if len(fileParts) > 2 {
		log.Fatalf("'%s' must only be used as header delimeter", HEADERDELIMITER)
	}

	headerPart := fileParts[0]
	bodyPart := fileParts[1]

	headers := parseHeaders(headerPart)
	body := parseBody(bodyPart)

	return Post{
		headers,
		body,
	}
}

func parseHeaders(headers []byte) Headers {
	var headerLines []string

	scanner := bufio.NewScanner(bytes.NewReader(headers))

	for scanner.Scan() {
		headerLines = append(headerLines, scanner.Text())
	}

	if len(headerLines) < 4 {
		log.Println("required headers are incomplete")
		return Headers{}
	}

	title := extractHeaderField(headerLines[0], "title")
	author := extractHeaderField(headerLines[1], "author")
	date, _ := time.Parse("2006-01-02", extractHeaderField(headerLines[2], "date"))
	tags := strings.Split(extractHeaderField(headerLines[3], "tags"), ",")

	return Headers{
		title,
		author,
		date,
		tags,
	}
}

func parseBody(body []byte) string {
	return string(blackfriday.MarkdownBasic(body))
}

func extractHeaderField(header string, fieldName string) string {
	if !strings.Contains(header, fieldName) {
		log.Printf("does not contain field %s", fieldName)
	}

	return strings.TrimSpace(strings.TrimPrefix(header, fieldName+":"))
}

func main() {
	inputFilePath := path.Clean(os.Args[1])
	inputFileExtension := path.Ext(inputFilePath)
	if inputFileExtension != INPUT_FILE_EXTENSION {
		log.Fatalf("wanted markdown file got: %s", inputFileExtension)
	}
	outPutFilePath := strings.Replace(inputFilePath, INPUT_FILE_EXTENSION, OUTPUT_FILE_EXTENSION, 1)

	log.Printf("reading file %s", inputFilePath)
	content, _ := os.ReadFile(inputFilePath)

	log.Println("parsing file")
	post := ParsePostFile(content)

	log.Printf("writing output to %s", outPutFilePath)
	log.Printf("--title: %s", post.Headers.Title)
	log.Printf("--author: %s", post.Headers.Author)
	log.Printf("--date: %s", post.Headers.Date.Format(time.RFC3339))
	out, err := os.OpenFile(outPutFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("could not open file")
	}

	writeWithLayout(out, post)
}

func writeWithLayout(out io.Writer, post Post) {
	const layoutTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta name="description" content="{{.Headers.Title}}" />
	<meta name="author" content="{{.Headers.Author}}" />
	<meta name="viewport" content="user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0, width=device-width" />
	<title>{{.Headers.Title}}</title>
</head>
<body>
{{.Body}}
</body>
</html>`

	tmpl, _ := template.New("post").Parse(layoutTemplate)

	if err := tmpl.Execute(out, post); err != nil {
		log.Fatal("error writing to file:", err)
	}

}
