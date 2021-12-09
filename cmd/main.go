package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	logos "github.com/v3rse/logos"
)

const INPUT_FILE_EXTENSION = ".md"
const OUTPUT_FILE_EXTENSION = ".html"
const LAYOUT_TEMPLATE = `<!DOCTYPE html>
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

func writeWithLayout(out io.Writer, page logos.Page) {
	tmpl, _ := template.New("page").Parse(LAYOUT_TEMPLATE)

	if err := tmpl.Execute(out, page); err != nil {
		log.Fatal("error writing to file:", err)
	}

}

func WritePage(inputFilePath string) {
	inputFileExtension := path.Ext(inputFilePath)
	if inputFileExtension != INPUT_FILE_EXTENSION {
		log.Printf("skipping '%s' has extension: %s", inputFilePath, inputFileExtension)
		return
	}
	outPutFilePath := strings.Replace(inputFilePath, INPUT_FILE_EXTENSION, OUTPUT_FILE_EXTENSION, 1)

	log.Printf("reading file %s", inputFilePath)
	content, _ := os.ReadFile(inputFilePath)

	log.Println("--parsing file")
	page := logos.NewPage("/", content)

	out, err := os.OpenFile(outPutFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("--could not open file")
	}

	writeWithLayout(out, page)
}

func isEmpty(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}

func assert(predicate bool, message string) {
	if !predicate {
		log.Fatalf(message)
	}
}

func usage() string {
	return "usage: logos <input-dir> <output-dir>"
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanDir(inputDirPath) {
	dirEntries, err := os.ReadDir(inputDirPath)
	check(err)

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			log.Printf("%s: %s", "file", entry.Name())
		}

		log.Printf("%s: %s", "file", entry.Name())
	}
}

func main() {
	// args
	assert(len(os.Args) >= 3, usage())
	assert(!isEmpty(os.Args[1]), fmt.Sprintf("expected input directory. \n%s", usage()))
	assert(!isEmpty(os.Args[2]), fmt.Sprintf("expected output directory. \n%s", usage()))

	inputDirPath := path.Clean(os.Args[1])
	outputDirPath := path.Clean(os.Args[2])

	log.Printf("out: %s", outputDirPath)

	// scan input root directory
	// collect pages and routes
	// generate nav (root files are at root of nav e.g. about us etc, content should be in folders)
	// write pages to output directory with layout + nav
	// write index with layout and nav
}
