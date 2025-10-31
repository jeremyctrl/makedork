<div align="center">

# makedork

[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Fluent, composable builder for crafting search dorks in Go.

</div>

`makedork` is a fluent builder library for composing advanced search dorks programmatically via a concise, chainable API.

## Example

```go
package main

import (
	"fmt"
	"github.com/jeremyctrl/makedork"
)

func main() {
	q := makedork.NewGoogleSearch().
		Site("example.com").
		InText("golang").
		FileType("pdf")

	fmt.Println(q.String())
	// site:"example.com" intext:"golang" filetype:"pdf"
}
```