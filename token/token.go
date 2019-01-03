package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar,xyy
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	//デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	//pattern
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
