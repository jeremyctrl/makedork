package makedork

import "github.com/jeremyctrl/makedork/pkg/googlesearch"

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}
