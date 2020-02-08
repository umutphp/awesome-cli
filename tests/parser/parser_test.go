package parser_test

import (
	"awesome/internal/package/parser"

	"testing"
)

func TestSplit(t *testing.T) {
	result := parser.Split("Text before subtext after", "before", "after")

	if result != " subtext " {
		t.Errorf("Split is incorrect, got: %s, want: %s.", result, " subtext ")
	}
}

func TestParseContentFromLine(t *testing.T) {
	name, url, desc, _ := parser.ParseContentFromLine("[Name](url) - Description.")

	if name != "Name" {
		t.Errorf("Name is incorrect, got: %s, want: %s.", name, "Name")
	}

	if url != "url" {
		t.Errorf("URL is incorrect, got: %s, want: %s.", url, "url")
	}

	if desc != "Description." {
		t.Errorf("Description is incorrect, got: %s, want: %s.", desc, "Description.")
	}
}

func TestLineToTitle(t *testing.T) {
	result := parser.LineToTitle(" # Title ")

	if result != "Title" {
		t.Errorf("Title is incorrect, got: %s, want: %s.", result, "Title")
	}
}
