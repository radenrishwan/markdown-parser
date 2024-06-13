package main

import (
	"errors"
	"log"
)

func main() {
	text := "#Hello, World\n  ##Hello, World\n*An italic text*\n**An bold text**"

	scan := NewScanner(text, keyword)

	scan.Scan()

	log.Println(scan.Tokens)
}

// markdown syntax
const (
	H1 = "HEADING 1"
	H2 = "HEADING 2"
	H3 = "HEADING 3"
	H4 = "HEADING 4"
	H5 = "HEADING 5"

	CH1 = "CLOSE HEADING 1"
	CH2 = "CLOSE HEADING 2"
	CH3 = "CLOSE HEADING 3"
	CH4 = "CLOSE HEADING 4"
	CH5 = "CLOSE HEADING 5"

	BOLD    = "BOLD"
	CBOLD   = "CLOSE BOLD"
	ITALIC  = "ITALIC"
	CITALIC = "CLOSE ITALIC"

	BLOCKQUOTE  = "BLOCKQUOTE"
	CBLOCKQUOTE = "CLOSE BLOCKQUOTE"

	UNORDERED_LIST = "UNORDERED_LIST" // TODO: implement later
	ORDERED_LIST   = "ORDERED_LIST"

	CODE  = "CODE"
	CCODE = "CLOSE CODE"

	HORIZONTAL_RULE = "HORIZONTAL_RULE"
	// LINK            = "LINK" // TODO: implement later
	// IMAGE           = "IMAGE"

	EOL   = "END OF LINE"
	TEXT  = "TEXT"
	CTEXT = "CLOSE TEXT"
	NL    = "NEWLINE"

	// TODO: add extended syntax later
)

var keyword = map[string]string{
	H1: "<h1>",
	H2: "<h2>",
	H3: "<h3>",
	H4: "<h4>",
	H5: "<h5>",

	CH1: "</h1>",
	CH2: "</h2>",
	CH3: "</h3>",
	CH4: "</h4>",
	CH5: "</h5>",

	BOLD:  "<b>",
	CBOLD: "</b>",

	ITALIC:  "<i>",
	CITALIC: "</i>",

	BLOCKQUOTE:  "<blockquote>",
	CBLOCKQUOTE: "</blockquote>",

	UNORDERED_LIST: "<ul>",
	ORDERED_LIST:   "<ol>", // TODO: gonna try something later

	CODE:  "<code>",
	CCODE: "</code>",

	HORIZONTAL_RULE: "<hr>",
	NL:              "<br>",

	TEXT:  "<p>",
	CTEXT: "</p>",
}

type Token struct {
	Type  string
	Value string
	Line  int
}

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
	Keyword map[string]string
	text    bool
}

func NewScanner(source string, keyword map[string]string) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
		Keyword: keyword,
		text:    false,
	}
}

func (scanner *Scanner) Scan() []Token {
	var store string

	for !scanner.isEOL() {
		scanner.Start = scanner.Current

		current := scanner.advance()
		log.Println("Current char: ", string(current))
		log.Printf("current: %d \n", scanner.Current)

		if scanner.text == false && store != "" {
			temp := scanner.Tokens[len(scanner.Tokens)-1]

			scanner.Tokens = append(scanner.Tokens[0:len(scanner.Tokens)-1], Token{
				Type:  TEXT,
				Line:  scanner.Line,
				Value: store,
			}, temp)

			store = ""
		}

		switch current {
		case '\n':
			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  NL,
				Line:  scanner.Line,
				Value: scanner.Keyword[NL],
			})

			scanner.Line += 1
			scanner.text = false
		case ' ':
			if scanner.text == true {
				store += string(current) // store value if not a TOKENS

			}
		case '#':
			// check next character
			if scanner.peek(4) == "####" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H5,
					Line:  scanner.Line,
					Value: scanner.Keyword[H5],
				})

				scanner.Current += 4
			} else if scanner.peek(3) == "###" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H4,
					Line:  scanner.Line,
					Value: scanner.Keyword[H4],
				})

				scanner.Current += 3
			} else if scanner.peek(2) == "##" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H3,
					Line:  scanner.Line,
					Value: scanner.Keyword[H3],
				})

				scanner.Current += 2
			} else if scanner.peek(1) == "#" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H2,
					Line:  scanner.Line,
					Value: scanner.Keyword[H2],
				})

				scanner.Current += 1
			} else {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H1,
					Line:  scanner.Line,
					Value: scanner.Keyword[H1],
				})
			}

			scanner.text = false
		case '*':
			// check next character * again or not
			if scanner.peek(1) == "*" {
				// check if close or open
				c, _ := scanner.checkPrevToken(len(scanner.Tokens) - 1)
				if c.Type == BOLD {
					scanner.Tokens = append(scanner.Tokens, Token{
						Type:  CBOLD,
						Line:  scanner.Line,
						Value: scanner.Keyword[CBOLD],
					})
				} else {
					scanner.Tokens = append(scanner.Tokens, Token{
						Type:  BOLD,
						Line:  scanner.Line,
						Value: scanner.Keyword[BOLD],
					})
				}

				scanner.Current += 1
				scanner.text = false
				continue
			}

			c, _ := scanner.checkPrevToken(len(scanner.Tokens) - 1)
			if c.Type == ITALIC {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  CITALIC,
					Line:  scanner.Line,
					Value: scanner.Keyword[CITALIC],
				})
			} else {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  ITALIC,
					Line:  scanner.Line,
					Value: scanner.Keyword[ITALIC],
				})
			}

			scanner.text = false
		case '-':
			// check next character -- or not
			if scanner.peek(2) == "--" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  HORIZONTAL_RULE,
					Line:  scanner.Line,
					Value: scanner.Keyword[HORIZONTAL_RULE],
				})

				scanner.Current += 2
				scanner.text = false
				continue
			}

			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  ORDERED_LIST,
				Line:  scanner.Line,
				Value: scanner.Keyword[ORDERED_LIST],
			})

			scanner.text = false
		case '`':
			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  CODE,
				Line:  scanner.Line,
				Value: scanner.Keyword[CODE],
			})

			scanner.text = false
		default:
			store += string(current) // store value if not a TOKENS
			scanner.text = true
		}

		if scanner.Current-1 == len(scanner.Source) {
			break
		}
	}

	// need to check in case store is not empty
	if scanner.text == false && store != "" {
		temp := scanner.Tokens[len(scanner.Tokens)-1]

		scanner.Tokens = append(scanner.Tokens[0:len(scanner.Tokens)-1], Token{
			Type:  TEXT,
			Line:  scanner.Line,
			Value: store,
		}, temp)

		store = ""
	}

	return scanner.Tokens
}

// check if scanner already EOL
func (scanner *Scanner) isEOL() bool {
	return scanner.Current == len(scanner.Source)
}

// get current string
func (scanner *Scanner) advance() byte {
	scanner.Current++

	return scanner.Source[scanner.Current-1]
}

func (scanner *Scanner) peek(lot int) string {
	// check if source is shorter than lot
	log.Printf("check peek: %d - len: %d", scanner.Current+lot, len(scanner.Source))
	if scanner.Current+lot > len(scanner.Source) {
		return ""
	}

	return string(scanner.Source[scanner.Current : scanner.Current+lot])
}

func (scanner *Scanner) checkPrevToken(index int) (Token, error) {
	if len(scanner.Tokens) < index {
		return Token{}, errors.New("out of index")
	}

	return scanner.Tokens[index], nil
}
