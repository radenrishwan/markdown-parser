package main

import (
	"log"

	markdownparser "github.com/radenrishwan/markdown-parser"
)

func main() {
	text := "#Hello, World\n  ##Hello, World\n*An italic text*\n**An bold text**"

	keyword := map[string]string{}
	for key, value := range markdownparser.DefaultKeyword {
		keyword[key] = value
	}

	keyword[markdownparser.H1] = "<h1 class=\"text-md text-slate-200 font-semibold\""

	scan := markdownparser.NewScanner(text, keyword)

	log.Println(scan.Tokens)

	result := markdownparser.Parsing(scan)

	log.Println(result)
}
