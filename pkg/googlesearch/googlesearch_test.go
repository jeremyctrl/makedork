package googlesearch_test

import (
	"testing"

	"github.com/jeremyctrl/makedork/pkg/googlesearch"
)

func TestSite(t *testing.T) {
	d := googlesearch.New().Site("example.com")
	want := `site:"example.com"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestInText(t *testing.T) {
	d := googlesearch.New().InText("text")
	want := `intext:"text"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFileType(t *testing.T) {
	d := googlesearch.New().FileType("pdf")
	want := `filetype:"pdf"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCache(t *testing.T) {
	d := googlesearch.New().Cache("www.google.com")
	want := `cache:"www.google.com"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRelated(t *testing.T) {
	d := googlesearch.New().Related("www.google.com")
	want := `related:"www.google.com"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGroup(t *testing.T) {
	d := googlesearch.New().
		Site("youtube.com").
		Group(googlesearch.New().InURL("dQw4w9WgXcQ").Or().InURL("2yJgwwDcgV8"))

	want := `site:"youtube.com" (inurl:"dQw4w9WgXcQ" | inurl:"2yJgwwDcgV8")`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
