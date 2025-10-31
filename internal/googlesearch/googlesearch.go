package googlesearch

import (
	"fmt"
	"strconv"
	"strings"
)

type GoogleSearch struct {
	tags []string
}

// New returns a new GoogleSearch builder instance.
func New() *GoogleSearch {
	return &GoogleSearch{}
}

func joinTag(tag, value string, quotes bool) string {
	if quotes {
		return tag + strconv.Quote(value)
	}

	return tag + value
}

// String returns the current Google search query as a string.
func (g *GoogleSearch) String() string {
	return strings.Join(g.tags, " ")
}

// Or appends a logical OR operator to the current query.
func (g *GoogleSearch) Or() *GoogleSearch {
	g.tags = append(g.tags, "|")
	return g
}

// And appends a logical AND operator to the current query.
func (g *GoogleSearch) And() *GoogleSearch {
	g.tags = append(g.tags, "&")
	return g
}

// Group adds a grouped subquery wrapped in parentheses.
func (g *GoogleSearch) Group(q *GoogleSearch) *GoogleSearch {
	g.tags = append(g.tags, fmt.Sprintf("(%s)", q.String()))
	return g
}

// A

// AllInAnchor adds an allinanchor: filter matching all specified anchor text.
func (g *GoogleSearch) AllInAnchor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinanchor:", value, true))
	return g
}

// AllInPostAuthor adds an allinpostauthor: filter matching all specified post authors.
func (g *GoogleSearch) AllInPostAuthor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinpostauthor:", value, true))
	return g
}

// AllInText adds an allintext: filter matching all specified text terms.
func (g *GoogleSearch) AllInText(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allintext:", value, true))
	return g
}

// AllInTitle adds an allintitle: filter matching all specified title terms.
func (g *GoogleSearch) AllInTitle(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allintitle:", value, true))
	return g
}

// AllInURL adds an allinurl: filter matching all specified URL terms.
func (g *GoogleSearch) AllInURL(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("allinurl:", value, true))
	return g
}

// After adds an after: filter to restrict results to those published after the specified date.
func (g *GoogleSearch) After(date string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("after:", date, false))
	return g
}

// B

// Before adds a before: filter to restrict results to those published before the specified date.
func (g *GoogleSearch) Before(date string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("before:", date, false))
	return g
}

// C

// Cache adds a cache: filter to search Google's cached version of a page.
func (g *GoogleSearch) Cache(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("cache:", value, true))
	return g
}

// E

// Exact appends a quoted, exact-match phrase to the query.
func (g *GoogleSearch) Exact(value string) *GoogleSearch {
	g.tags = append(g.tags, strconv.Quote(value))
	return g
}

// F

// FileType adds a filetype: filter to restrict results to a given file extension.
func (g *GoogleSearch) FileType(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("filetype:", value, true))
	return g
}

// I

// InAnchor adds an inanchor: filter matching terms in anchor text.
func (g *GoogleSearch) InAnchor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inanchor:", value, true))
	return g
}

// InPostAuthor adds an inpostauthor: filter matching posts by the specified author.
func (g *GoogleSearch) InPostAuthor(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inpostauthor:", value, true))
	return g
}

// InText adds an intext: filter matching the specified term in the page body.
func (g *GoogleSearch) InText(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("intext:", value, true))
	return g
}

// InTitle adds an intitle: filter matching terms in the page title.
func (g *GoogleSearch) InTitle(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("intitle:", value, true))
	return g
}

// InURL adds an inurl: filter matching terms in the page URL.
func (g *GoogleSearch) InURL(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("inurl:", value, true))
	return g
}

// L

// Link adds a link: filter to find pages linking to the specified URL.
func (g *GoogleSearch) Link(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("link:", value, true))
	return g
}

// P

// Plain appends an unquoted, plain text term to the query.
func (g *GoogleSearch) Plain(value string) *GoogleSearch {
	g.tags = append(g.tags, value)
	return g
}

// R

// Related adds a related: filter to find pages related to the specified site.
func (g *GoogleSearch) Related(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("related:", value, true))
	return g
}

// S

// Site adds a site: filter to restrict results to the specified domain.
func (g *GoogleSearch) Site(value string) *GoogleSearch {
	g.tags = append(g.tags, joinTag("site:", value, true))
	return g
}
