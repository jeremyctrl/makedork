package bingsearch

import (
	"fmt"
	"strconv"
	"strings"
)

type BingSearch struct {
	tags []string
}

func New() *BingSearch {
	return &BingSearch{}
}

func joinTag(tag, value string, quotes bool) string {
	if quotes {
		return tag + strconv.Quote(value)
	}

	return tag + value
}

func (b *BingSearch) String() string {
	return strings.Join(b.tags, " ")
}

func (b *BingSearch) Or() *BingSearch {
	b.tags = append(b.tags, "|")
	return b
}

func (b *BingSearch) And() *BingSearch {
	b.tags = append(b.tags, "&")
	return b
}

func (b *BingSearch) Group(q *BingSearch) *BingSearch {
	b.tags = append(b.tags, fmt.Sprintf("(%s)", q.String()))
	return b
}

// C

func (b *BingSearch) Contains(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("contains:", value, true))
	return b
}

// F

func (b *BingSearch) Feed(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("feed:", value, true))
	return b
}

func (b *BingSearch) FileType(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("filetype:", value, true))
	return b
}

// I

func (b *BingSearch) InBody(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("inbody:", value, true))
	return b
}

func (b *BingSearch) InTitle(value string) *BingSearch {
	b.tags = append(b.tags, joinTag("intitle:", value, true))
	return b
}
