package logos

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

const HEADERDELIMITER = "---"

type Headers struct {
	Title  string
	Author string
	Date   time.Time
	Tags   []string
}

type Page struct {
	Route   string
	Headers Headers
	Body    string
}

func NewPage(route string, file []byte) Page {
	fileParts := bytes.Split(file, []byte(HEADERDELIMITER))

	if len(fileParts) > 2 {
		log.Fatalf("'%s' must only be used as header delimeter", HEADERDELIMITER)
	}

	headerPart := fileParts[0]
	bodyPart := fileParts[1]

	headers := parseHeaders(headerPart)
	body := parseBody(bodyPart)

	return Page{
		route,
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
