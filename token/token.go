package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // ADD, FOOBAR, X, Y, ...
	INT   = "INT"   // 123456

	// operators
	ASSIGN="="
	PLUS="+"

	// delimiters
	COMMA=","
	SEMICOLON = ";"
	LPAREN="("
	RPAREN=")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
