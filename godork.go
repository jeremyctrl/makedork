package godork

import "github.com/jeremyctrl/godork/internal/googlesearch"

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}
