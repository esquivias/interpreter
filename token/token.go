package token

// Type will be used as token.Type by other packages; avoid stutter by calling this Type and not TokenType.
type Type string

// Token data structure
type Token struct {
	Type    Type   // string;
	Literal string // string; has the advantage of being easy to debug
}

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// Define the possible Token.Type as constants
const (
	//
	// Special Types
	//

	// ILLEGAL an illegal or unknown token type
	ILLEGAL = "ILLEGAL"

	// EOF is end of file
	EOF = "EOF"

	//
	// Identifiers & Literals
	//

	// IDENT is an identifier type
	IDENT = "IDENT"

	// INT is an integer type
	INT = "INT"

	//
	// Operators
	//

	// ASSIGN is an operator type
	ASSIGN = "="

	// PLUS is an operator type
	PLUS = "+"

	//
	// Delimiters
	//

	// COMMA is a delimiter type
	COMMA = ","

	// SEMICOLON is a delimiter type
	SEMICOLON = ";"

	// LPAREN  is a delimiter type
	LPAREN = "("

	// RPAREN  is a delimiter type
	RPAREN = ")"

	// LBRACE  is a delimiter type
	LBRACE = "{"

	// RBRACE  is a delimiter type
	RBRACE = "}"

	//
	// Keywords
	//

	// FUNCTION is a keyword type
	FUNCTION = "FUNCTION"

	// LET is a keyword type
	LET = "LET"
)

// LookupIdent returns a keyword's constant if found, or IDENT if not, as the token.Type
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
