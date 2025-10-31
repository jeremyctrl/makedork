package bingsearch_test

import (
	"testing"

	"github.com/jeremyctrl/makedork/pkg/bingsearch"
)

func TestContains(t *testing.T) {
	d := bingsearch.New().Contains("lorem ipsum")
	want := `contains:"lorem ipsum"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFeed(t *testing.T) {
	d := bingsearch.New().Feed("example.com")
	want := `feed:"example.com"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFileType(t *testing.T) {
	d := bingsearch.New().FileType("pdf")
	want := `filetype:"pdf"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHasFeed(t *testing.T) {
	d := bingsearch.New().HasFeed("example.com")
	want := `hasfeed:"example.com"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestInBody(t *testing.T) {
	d := bingsearch.New().InBody("climate change")
	want := `inbody:"climate change"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestInTitle(t *testing.T) {
	d := bingsearch.New().InTitle("research")
	want := `intitle:"research"`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGroup(t *testing.T) {
	d := bingsearch.New().
		FileType("pdf").
		Group(bingsearch.New().InTitle("report").Or().InBody("summary"))

	want := `filetype:"pdf" (intitle:"report" | inbody:"summary")`
	if got := d.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
