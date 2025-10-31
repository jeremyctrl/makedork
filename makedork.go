package makedork

import (
	"github.com/jeremyctrl/makedork/pkg/bingsearch"
	"github.com/jeremyctrl/makedork/pkg/googlesearch"
)

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}

func NewBingSearch() *bingsearch.BingSearch {
	return bingsearch.New()
}
