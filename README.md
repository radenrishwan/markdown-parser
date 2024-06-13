# Markdown parser in go

`radenrishwan/markdown-parser` is a library in go to parsing markdown with feature:
- parsing into html, or etc (you can implement yourself)
- easy to use
- no dependencies except std

## Documentation

not implemented yet.

### How to use
you can check [here](https://github.com/radenrishwan/markdown-parser/tree/master/example) to see all examples

```go
package main

import (
	"log"

	markdownparser "github.com/radenrishwan/markdown-parser"
)

func main() {
	text := "#Hello, World\n  ##Hello, World\n*An italic text*\n**An bold text**"

	scan := markdownparser.NewScanner(text, markdownparser.DefaultKeyword)

	log.Println(scan.Tokens)

	result := markdownparser.Parsing(scan)

	log.Println(result)
}
```

## License
check #[License](https://github.com/radenrishwan/markdown-parser/blob/master/LICENSE)
