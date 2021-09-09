package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Post struct {
	title  string
	author string
	date   time.Time
	tags   []string
	body   []string
}

func ReadPostFile(path string) []byte {
	file, _ := os.ReadFile(path)

	return file
}

func ParsePostFile(file []byte) Post {
	scanner := bufio.NewScanner(bytes.NewReader(file))
	headers := []string{}
	var body []string

	for scanner.Scan() {
		if scanner.Text() == "---" {
			for scanner.Scan() {
				body = append(body, scanner.Text())
			}
		}

		headers = append(headers, scanner.Text())
	}

	title := extractHeaderField(headers[0], "title")
	author := extractHeaderField(headers[1], "author")
	date, _ := time.Parse("2006-01-02", extractHeaderField(headers[2], "date"))
	log.Print(extractHeaderField(headers[2], "date"))
	tags := strings.Split(extractHeaderField(headers[3], "tags"), ",")

	return Post{
		title,
		author,
		date,
		tags,
		body,
	}
}

func extractHeaderField(header string, fieldName string) string {
	if !strings.Contains(header, fieldName) {
		log.Printf("does not contain field %s", fieldName)
	}

	return strings.TrimSpace(strings.TrimPrefix(header, fieldName+":"))
}

func main() {
	input := os.Args[1]

	output := ParsePostFile(ReadPostFile(input))

	fmt.Printf("%v", output)
}
