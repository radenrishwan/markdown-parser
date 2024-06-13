package markdownparser

func Parsing(scanner *Scanner) string {
	var result string

	for index, token := range scanner.Tokens {
		switch token.Type {
		case TEXT:
			// check if before text isn't token
			before := scanner.Tokens[index-1].Type

			if before != TEXT {
				continue
			}

			// if before == H1 || before == H2 || before == H3 || before == H4 || before == H5 {
			// 	continue
			// }

			result += scanner.Source
		case H1, H2, H3, H4, H5:
			result += scanner.Keyword[token.Type]

			// adding text and close tag
			if index < len(scanner.Tokens)-1 {
				next := scanner.Tokens[index+1]
				if next.Type == TEXT {
					result += next.Value
				}
			}

			switch token.Type {
			case H1:
				result += scanner.Keyword[CH1]
			case H2:
				result += scanner.Keyword[CH2]
			case H3:
				result += scanner.Keyword[CH3]
			case H4:
				result += scanner.Keyword[CH4]
			case H5:
				result += scanner.Keyword[CH5]
			}
		case BOLD, ITALIC, BLOCKQUOTE, CODE, HORIZONTAL_RULE, NL:
			result += scanner.Keyword[token.Type]

			if index < len(scanner.Tokens)-1 {
				next := scanner.Tokens[index+1]
				if next.Type == TEXT {
					result += next.Value
				}
			}
		case CBOLD, CITALIC, CBLOCKQUOTE, CCODE:
			result += scanner.Keyword[token.Type]
		}
	}

	return result
}
