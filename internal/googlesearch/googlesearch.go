package googlesearch

import (
	"fmt"
	"strings"
)

type GoogleSearch struct {
	tags []string
}

func New() *GoogleSearch {
	return &GoogleSearch{}
}

func (g *GoogleSearch) String() string {
	return strings.Join(g.tags, " ")
}

func (g *GoogleSearch) Or() *GoogleSearch {
	g.tags = append(g.tags, "|")
	return g
}

func (g *GoogleSearch) And() *GoogleSearch {
	g.tags = append(g.tags, "&")
	return g
}

func (g *GoogleSearch) Group(q *GoogleSearch) *GoogleSearch {
	g.tags = append(g.tags, fmt.Sprintf("(%s)", q.String()))
	return g
}
