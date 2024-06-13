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
