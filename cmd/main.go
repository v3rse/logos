package main

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	logos "github.com/v3rse/logos"
)

const INPUT_FILE_EXTENSION = ".md"
const OUTPUT_FILE_EXTENSION = ".html"

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
	post := logos.NewPost(content)

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

func writeWithLayout(out io.Writer, post logos.Post) {
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
