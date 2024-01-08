
package token


type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	UNORDERED_LIST = "UL"
	ORDERED_LIST = "OL"
	H1 = "HEADER 1"
	H2 = "HEADER 2"
	H3 = "HEADER 3"
	NEW_LINE = "NEW LINE"
	TEXT = "TEXT"
	BOLD = "BOLD"
	ITALIC = "ITALIC"
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
)
