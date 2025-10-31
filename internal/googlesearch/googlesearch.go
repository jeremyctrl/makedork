package googlesearch

import (
	"fmt"
	"strconv"
	"strings"
)

type GoogleSearch struct {
	tags []string
}

func New() *GoogleSearch {
	return &GoogleSearch{}
}

func joinTag(tag, value string, quotes bool) string {
	if quotes {
		return tag + strconv.Quote(value)
	}

	return tag + value
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

// A
func (g *GoogleSearch) AllInAnchor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinanchor:", value, true))
	return g
}
func (g *GoogleSearch) AllInPostAuthor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinpostauthor:", value, true))
	return g
}
func (g *GoogleSearch) AllInText(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allintext:", value, true))
	return g
}
func (g *GoogleSearch) AllInTitle(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allintitle:", value, true))
	return g
}
func (g *GoogleSearch) AllInURL(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinurl:", value, true))
	return g
}
func (g *GoogleSearch) After(date string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("after:", date, false))
	return g
}

// B
func (g *GoogleSearch) Before(date string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("before:", date, false))
	return g
}
