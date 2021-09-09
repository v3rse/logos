package main

import (
	"strings"
	"testing"
	"time"
)

func TestReadPostFile(t *testing.T) {
	path := "post.md"
	want := `title: Hello Blog
author: Nana
date: 2021-09-09
tags: opinion, thoughts
---
Hello Blog`
	got := convertToSimpleString(ReadPostFile(path))

	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func TestParsePostFile(t *testing.T) {
	file := []byte(`title: Hello Blog
author: Nana
date: 2021-09-09
tags: opinion, thoughts
---
Hello Blog
These are my thoughts`)

	want := Post{
		title:  "Hello Blog",
		author: "Nana",
		date:   time.Date(2021, time.September, 9, 0, 0, 0, 0, time.UTC),
		tags:   []string{"opinion", "thoughts"},
		body:   []string{"Hello Blog", "These are my thoughts"},
	}

	got := ParsePostFile(file)

	if !assertPost(t, want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func assertPost(t testing.TB, want, got Post) bool {
	t.Helper()
	switch {
	case want.title != got.title:
		return false
	case want.author != got.author:
		return false
	}

	return true
}

func convertToSimpleString(data []byte) string {
	return strings.Trim(string(data), "\n")
}
