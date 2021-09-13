package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParsePostFile(t *testing.T) {
	t.Run("should parse headers", func(t *testing.T) {
		file := []byte(`title: Hello Blog
author: Nana
date: 2021-09-09
tags: opinion, thoughts
---
`)

		want := Post{
			Headers{
				"Hello Blog",
				"Nana",
				time.Date(2021, time.September, 9, 0, 0, 0, 0, time.UTC),
				[]string{"opinion", "thoughts"},
			},
			`some text`,
		}

		got := ParsePostFile(file)

		assertPostHeaders(t, want, got)
	})

	t.Run("should parse body", func(t *testing.T) {
		file := []byte(`---
Hello Blog
`)

		want := "<p>Hello Blog</p>\n"

		got := ParsePostFile(file)

		assertPostBody(t, want, got.Body)
	})
}

func assertPostHeaders(t testing.TB, want, got Post) {
	t.Helper()

	if reflect.DeepEqual(want.Headers, got.Headers) {
		t.Errorf("post headers: want %v got %v", want.Headers, got.Headers)
	}
}

func assertPostBody(t testing.TB, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("post body: want %#v got %#v", want, got)
	}
}

func convertToSimpleString(data []byte) string {
	return strings.Trim(string(data), "\n")
}
