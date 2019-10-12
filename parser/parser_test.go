package parser

import (
	"bytes"
	"strings"
	"testing"
)

func TestIsHeading(t *testing.T) {
	tests := []struct {
		expected bool
		example  string
	}{
		{true, "# yay"},
		{true, "## yay"},
		{true, "### yay"},
		{true, "#### yay"},
		{true, "##### yay"},
		{true, "## yay#"},
		{false, "nay"},
		{false, "nay#"},
		{false, "nay##"},
		{false, "n#ay"},
	}

	for i, tt := range tests {
		result := IsHeading(tt.example)
		if tt.expected != result {
			t.Fatalf("tests[%d] %q - output wrong. expected=%t, got=%t",
				i, tt.example, tt.expected, result)
		}
	}
}

func TestWriteLinkText(t *testing.T) {
	tests := []struct {
		expected string
		example  string
	}{
		{"* [one]", "one"},
		{"* [one two]", "one two"},
		{"* [one two three]", "one two three"},
		{"* [Capital]", "Capital"},
		{"* [ space]", " space"},
		{"* [space space]", "space space"},
		{"* [big    Space]", "big    Space"},
	}

	for i, tt := range tests {
		var b bytes.Buffer
		b = WriteLinkText(b, tt.example)
		result := b.String()
		if tt.expected != result {
			t.Fatalf("tests[%d] %q - output wrong. expected=%q, got=%q",
				i, tt.example, tt.expected, result)
		}
	}
}

func TestWriteLink(t *testing.T) {
	tests := []struct {
		expected string
		example  string
	}{
		{"(#one)", "one"},
		{"(#one-two)", "one two"},
		{"(#one-two-three)", "one two three"},
		{"(#capital)", "Capital"},
		{"(#-space)", " space"},
		{"(#space-space)", "space space"},
		{"(#big----space)", "big    Space"},
	}

	for i, tt := range tests {
		var b bytes.Buffer
		split := strings.Split(tt.example, " ")
		b = WriteLink(b, split)
		result := b.String()
		if tt.expected != result {
			t.Fatalf("tests[%d] %q - output wrong. expected=%q, got=%q",
				i, tt.example, tt.expected, result)
		}
	}
}

func TestWriteDepth(t *testing.T) {
	tests := []struct {
		expected string
		example  string
	}{
		{"", "#"},
		{"", "##"},
		{"  ", "###"},
		{"    ", "####"},
		{"      ", "#####"},
	}

	for i, tt := range tests {
		var b bytes.Buffer
		b = WriteDepth(b, tt.example)
		result := b.String()
		if tt.expected != result {
			t.Fatalf("tests[%d] %q - output wrong. expected=%q, got=%q",
				i, tt.example, tt.expected, result)
		}
	}
}
