package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"slices"
)

func (p *Parser) advance()             { p.index++ }
func (p *Parser) current() lexer.Token { return p.tokens[p.index] }

func (p *Parser) expect(arg any) lexer.Token {
	token := p.current()

	switch v := arg.(type) {
	case lexer.TokenType:
		if token.Type != v {
			panic(fmt.Sprintf("line: %d: expected token type %s got %s", token.Line, v, token.Type))
		}
	case string:
		if token.Literal != v {
			panic(fmt.Sprintf("line: %d: expected literal %s got %s", token.Line, v, token.Literal))
		}
	default:
		panic("argument to expect() is neither a token type or literal")
	}

	return token
}

func (p *Parser) parseStatement() ast.Node {
	// statements start with a keyword
	token := p.expect(lexer.Keyword)
	p.advance()

	if slices.Contains([]string{"incr", "decr", "clear"}, token.Literal) {
		next := p.expect(lexer.Identifer)
		p.advance()

		p.expect(";")
		p.advance()

		switch token.Literal {
		case "incr":
			return &ast.Incr{Var: next.Literal}
		case "decr":
			return &ast.Decr{Var: next.Literal}
		case "clear":
			return &ast.Clear{Var: next.Literal}
		}
	} else if token.Literal == "while" {
		return p.parseWhile()
	} else if token.Literal == "copy" {
		return p.parseCopy()
	}

	panic("unknown keyword: " + token.Literal)
}

func (p *Parser) parseWhile() *ast.While {
	node := ast.While{}

	token := p.expect(lexer.Identifer)
	node.Var = token.Literal
	p.advance()

	p.expect("not")
	p.advance()

	token = p.expect("0")
	p.advance()

	p.expect("do")
	p.advance()

	p.expect(";")
	p.advance()

	for p.current().Literal != "end" {
		subnode := p.parseStatement()
		node.Body = append(node.Body, subnode)

		if p.index >= len(p.tokens) {
			panic("program ended without termination of the while loop")
		}
	}
	p.advance()

	p.expect(";")
	p.advance()

	return &node
}

func (p *Parser) parseCopy() *ast.Copy {
	node := ast.Copy{}

	token := p.expect(lexer.Identifer)
	node.Src = token.Literal
	p.advance()

	p.expect("to")
	p.advance()

	token = p.expect(lexer.Identifer)
	node.Dst = token.Literal
	p.advance()

	p.expect(";")
	p.advance()

	return &node
}
