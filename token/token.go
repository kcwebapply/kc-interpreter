package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar,xyy
	INT   = "INT"

	//型
	TYPE = "TYPE"

	// 演算子
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ATSTERISK = "*"
	SLASK     = "/"

	LT = "<"
	GT = ">"

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

	// true or false
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	EQ     = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"int":    TYPE,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
