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
