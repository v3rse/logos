package logos

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewPost(t *testing.T) {
	t.Run("should have headers", func(t *testing.T) {
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

		got := NewPost(file)

		assertPostHeaders(t, want, got)
	})

	t.Run("should have body", func(t *testing.T) {
		file := []byte(`title: Hello Blog
author: Nana
date: 2021-09-09
tags: opinion, thoughts
---
Hello Blog
`)

		want := "<p>Hello Blog</p>\n"

		got := NewPost(file)

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
