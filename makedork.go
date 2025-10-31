package makedork

import "github.com/jeremyctrl/makedork/internal/googlesearch"

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}
