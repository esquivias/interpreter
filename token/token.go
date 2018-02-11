package token

// Type will be used as token.Type by other packages; avoid stutter by calling this Type and not TokenType.
type Type string

// Token data structure
type Token struct {
	Type    Type   // string;
	Literal string // string; has the advantage of being easy to debug
}

var keywords = map[string]Type{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
}

// Define the possible Token.Type as constants
const (

	//
	// Special Types
	//

	// EOF is end of file
	EOF = "EOF"

	// ILLEGAL an illegal or unknown token type
	ILLEGAL = "ILLEGAL"

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

	// ASTERISK is an operator type
	ASTERISK = "*"

	// BANG is an operator type
	BANG = "!"

	// EQ is an operator type
	EQ = "=="

	// GT is an operator type
	GT = ">"

	// LT is an operator type
	LT = "<"

	// MINUS is an operator type
	MINUS = "-"

	// NEQ is a operator type
	NEQ = "!="

	// PLUS is an operator type
	PLUS = "+"

	// SLASH is an operator type
	SLASH = "/"

	//
	// Delimiters
	//

	// COMMA is a delimiter type
	COMMA = ","

	// LBRACE  is a delimiter type
	LBRACE = "{"

	// LPAREN  is a delimiter type
	LPAREN = "("

	// RBRACE  is a delimiter type
	RBRACE = "}"

	// RPAREN  is a delimiter type
	RPAREN = ")"

	// SEMICOLON is a delimiter type
	SEMICOLON = ";"

	//
	// Keywords
	//

	// ELSE is a keyword type
	ELSE = "ELSE"

	// FALSE is a keyword type
	FALSE = "FALSE"

	// FUNCTION is a keyword type
	FUNCTION = "FUNCTION"

	// IF is a keyword type
	IF = "IF"

	// LET is a keyword type
	LET = "LET"

	// RETURN is a keyword type
	RETURN = "RETURN"

	// TRUE is a keyword type
	TRUE = "TRUE"
)

// LookupIdent returns a keyword's constant if found, or IDENT if not, as the token.Type
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
