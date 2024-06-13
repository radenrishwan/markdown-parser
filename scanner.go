package markdownparser

import "errors"

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
	result := &Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
		Keyword: keyword,
		text:    false,
	}

	result.Scan()

	return result
}

func (scanner *Scanner) Scan() []Token {
	var store string

	for !scanner.isEOL() {
		scanner.Start = scanner.Current

		current := scanner.advance()
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
