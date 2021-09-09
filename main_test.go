package main

import (
	"strings"
	"testing"
)

func TestReadPostFile(t *testing.T) {
	path := "post.md"
	want := "Hello Blog"
	got := convertToSimpleString(ReadPostFile(path))

	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func convertToSimpleString(data []byte) string {
	return strings.Trim(string(data), "\n")
}
