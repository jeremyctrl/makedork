package googlesearch

import "strings"

type GoogleSearch struct {
	tags []string
}

func New() *GoogleSearch {
	return &GoogleSearch{}
}

func (g *GoogleSearch) String() string {
	return strings.Join(g.tags, " ")
}
