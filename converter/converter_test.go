package converter

import (
	"testing"
)

func TestConvertHeaders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "H1",
			input:    "# Heading 1",
			expected: `=<span style="color:#280071;">Heading 1</span>=`,
		},
		{
			name:     "H2",
			input:    "## Heading 2",
			expected: `==<span style="color:#5B2A86;">Heading 2</span>==`,
		},
		{
			name:     "H3",
			input:    "### Heading 3",
			expected: `===<span style="color:#7B4AA3;">Heading 3</span>===`,
		},
		{
			name:     "No Header",
			input:    "Just text",
			expected: "Just text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertHeaders(tt.input)
			if got != tt.expected {
				t.Errorf("ConvertHeaders() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertBoldItalic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Bold **",
			input:    "**bold**",
			expected: "'''bold'''",
		},
		{
			name:     "Bold __",
			input:    "__bold__",
			expected: "'''bold'''",
		},
		{
			name:     "Italic *",
			input:    "*italic*",
			expected: "''italic''",
		},
		{
			name:     "Italic _",
			input:    "_italic_",
			expected: "''italic''",
		},
		{
			name:     "Highlight ==",
			input:    "==highlight==",
			expected: `<mark style="background-color:#eacbbb">highlight</mark>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertBoldItalic(tt.input)
			if got != tt.expected {
				t.Errorf("ConvertBoldItalic() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "External Link",
			input:    "[Google](https://google.com)",
			expected: "[https://google.com Google]",
		},
		{
			name:     "No Link",
			input:    "Just text",
			expected: "Just text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertLinks(tt.input)
			if got != tt.expected {
				t.Errorf("ConvertLinks() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertLists(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Unordered List",
			input:    "- Item 1",
			expected: "* Item 1",
		},
		{
			name:     "Ordered List",
			input:    "1. Item 1",
			expected: "# Item 1",
		},
		{
			name:     "Nested List",
			input:    "  - Nested",
			expected: "** Nested",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertLists(tt.input)
			if got != tt.expected {
				t.Errorf("ConvertLists() = %v, want %v", got, tt.expected)
			}
		})
	}
}
