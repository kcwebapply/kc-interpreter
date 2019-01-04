package parser

import (
	"github.com/kcwebapply/kc-interpreter/lexar"
	"github.com/kcwebapply/kc-interpreter/token"
)

type Parser struct {
	l *lexar.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexar.Lexer) *Parser {
	p := &Parser{l: lexer}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
