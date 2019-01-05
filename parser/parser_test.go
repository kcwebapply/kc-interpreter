package parser

import (
	"fmt"
	"testing"

	"github.com/kcwebapply/kc-interpreter/ast"
	"github.com/kcwebapply/kc-interpreter/lexar"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x = 1+a;
  let y = 5;
  unko=8;
  let foobar = 838383;
  `

	l := lexar.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("parseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// Test let statements parsing.
func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not * ast.Letstatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s.got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Token literal not %s,got %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d erros", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q, msg", msg)
	}
	t.FailNow()
}

// Test Return statement parsing
func TestReturnsStatement(t *testing.T) {
	input := `
  return 5;
  return "ok";
  return kc;
  `

	l := lexar.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not * ast.ReturnStatement. got=%T", returnStmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("s.TokenLiteral not 'let'. got=%q", returnStmt.TokenLiteral())
		}

		fmt.Println(returnStmt.ReturnValue)

	}

}
