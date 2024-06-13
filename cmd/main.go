package main

import "log"

func main() {
	text := "#Hello, World\n  ##Hello, World\n*An italic text*"

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

	BOLD           = "BOLD"
	ITALIC         = "ITALIC"
	BLOCKQUOTE     = "BLOCKQUOTE"
	UNORDERED_LIST = "UNORDERED_LIST"
	ORDERED_LIST   = "ORDERED_LIST"

	CODE            = "CODE"
	HORIZONTAL_RULE = "HORIZONTAL_RULE"
	// LINK            = "LINK" // TODO: implement later
	// IMAGE           = "IMAGE"

	EOL  = "END OF LINE"
	TEXT = "TEXT"
	NL   = "NEWLINE"

	// TODO: add extended syntax later
)

var keyword = map[string]string{
	"#":     "<h1>",
	"##":    "<h2>",
	"###":   "<h3>",
	"####":  "<h4>",
	"#####": "<h5>",

	"**":           "<b>",
	"*":            "<i>",
	">":            "<blockquote>",
	"-":            "<ul>",
	"ORDERED_LIST": ORDERED_LIST, // TODO: gonna try something later
	"`":            "<code>",
	"---":          "<hr>",
	"\n":           "<br>",
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
		switch current {
		case '\n':
			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  NL,
				Line:  scanner.Line,
				Value: "\n",
			})

			scanner.Line += 1
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
					Value: scanner.Keyword["######"],
				})

				scanner.Current += 4
			} else if scanner.peek(3) == "###" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H4,
					Line:  scanner.Line,
					Value: scanner.Keyword["####"],
				})

				scanner.Current += 3
			} else if scanner.peek(2) == "##" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H3,
					Line:  scanner.Line,
					Value: scanner.Keyword["###"],
				})

				scanner.Current += 2
			} else if scanner.peek(1) == "#" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H2,
					Line:  scanner.Line,
					Value: scanner.Keyword["##"],
				})

				scanner.Current += 1
			} else {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  H1,
					Line:  scanner.Line,
					Value: scanner.Keyword["#"],
				})
			}

			scanner.text = false
		case '*':
			// check next character * again or not
			if scanner.peek(1) == "*" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  BOLD,
					Line:  scanner.Line,
					Value: scanner.Keyword["**"],
				})

				scanner.Current += 1
				scanner.text = false
				continue
			}

			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  ITALIC,
				Line:  scanner.Line,
				Value: scanner.Keyword["*"],
			})
			scanner.text = false
		case '-':
			// check next character -- or not
			if scanner.peek(2) == "--" {
				scanner.Tokens = append(scanner.Tokens, Token{
					Type:  HORIZONTAL_RULE,
					Line:  scanner.Line,
					Value: scanner.Keyword["-"],
				})

				scanner.Current += 2
				scanner.text = false
				continue
			}

			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  ORDERED_LIST,
				Line:  scanner.Line,
				Value: scanner.Keyword["-"],
			})

			scanner.text = false
		case '`':
			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  CODE,
				Line:  scanner.Line,
				Value: scanner.Keyword["`"],
			})

			scanner.text = false
		default:
			store += string(current) // store value if not a TOKENS
			scanner.text = true
		}

		if scanner.text == false && store != "" {
			scanner.Tokens = append(scanner.Tokens, Token{
				Type:  TEXT,
				Line:  scanner.Line,
				Value: store,
			})

			store = ""
		}

		log.Printf("current: %d - len: %d", scanner.Current, len(scanner.Source))

		if scanner.Current-1 == len(scanner.Source) {
			break
		}
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
	if scanner.Current+lot > len(scanner.Source) {
		return ""
	}

	return string(scanner.Source[scanner.Current+lot])
}
