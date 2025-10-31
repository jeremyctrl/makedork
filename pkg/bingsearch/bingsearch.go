package bingsearch

import (
	"fmt"
	"strconv"
	"strings"
)

type BingSearch struct {
	tags []string
}

// New returns a new BingSearch builder instance.
func New() *BingSearch {
	return &BingSearch{}
}

func joinTag(tag, value string, quotes bool) string {
	if quotes {
		return tag + strconv.Quote(value)
	}

	return tag + value
}

// String returns the current Bing search query as a string.
func (b *BingSearch) String() string {
	return strings.Join(b.tags, " ")
}

// Or appends a logical OR operator to the current query.
func (b *BingSearch) Or() *BingSearch {
	b.tags = append(b.tags, "|")
	return b
}

// And appends a logical AND operator to the current query.
func (b *BingSearch) And() *BingSearch {
	b.tags = append(b.tags, "&")
	return b
}

// Group adds a grouped subquery wrapped in parentheses.
func (b *BingSearch) Group(q *BingSearch) *BingSearch {
	b.tags = append(b.tags, fmt.Sprintf("(%s)", q.String()))
	return b
}

// C

// Contains adds a contains: filter to find pages that link to a specified file type.
func (b *BingSearch) Contains(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("contains:", value, true))
	return b
}

// F

// Feed adds a feed: filter to find RSS or Atom feeds related to the specified URL.
func (b *BingSearch) Feed(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("feed:", value, true))
	return b
}

// FileType adds a filetype: filter to restrict results to a given file extension.
func (b *BingSearch) FileType(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("filetype:", value, true))
	return b
}

// H

// HasFeed adds a hasfeed: filter to find pages that contain links to RSS or Atom feeds related to the value.
func (b *BingSearch) HasFeed(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("hasfeed:", value, true))
	return b
}

// I

// InBody adds an inbody: filter matching the specified term in the page body.
func (b *BingSearch) InBody(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("inbody:", value, true))
	return b
}

// InTitle adds an intitle: filter matching terms in the page title.
func (b *BingSearch) InTitle(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("intitle:", value, true))
	return b
}
