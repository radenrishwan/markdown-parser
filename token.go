package markdownparser

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

// if you want to change some tag, you can copy or change default Keyword
var DefaultKeyword = map[string]string{
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
