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

// C
func (g *GoogleSearch) Cache(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("cache:", value, true))
	return g
}

// E
func (g *GoogleSearch) Exact(value string) *GoogleSearch {
	g.tags = append(g.tags, strconv.Quote(value))
	return g
}

// F
func (g *GoogleSearch) FileType(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("filetype:", value, true))
	return g
}

// I
func (g *GoogleSearch) InAnchor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inanchor:", value, true))
	return g
}
func (g *GoogleSearch) InPostAuthor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inpostauthor:", value, true))
	return g
}
func (g *GoogleSearch) InText(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("intext:", value, true))
	return g
}
func (g *GoogleSearch) InTitle(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("intitle:", value, true))
	return g
}
func (g *GoogleSearch) InURL(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inurl:", value, true))
	return g
}

// L
func (g *GoogleSearch) Link(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("link:", value, true))
	return g
}

// R
func (g *GoogleSearch) Related(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("related:", value, true))
	return g
}

// S
func (g *GoogleSearch) Site(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("site:", value, true))
	return g
}
