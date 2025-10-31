package bingsearch

import (
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
