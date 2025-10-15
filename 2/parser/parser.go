package parser

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"strconv"
)

// this whole file is quite messy

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) expectLiteral(literal string) lexer.Token {
	token := p.peek()

	if token.Literal != literal {
		panic(fmt.Sprintf("expected literal %s got %s", literal, token.Literal))
	}

	return token
}

func (p *Parser) expectType(typ lexer.TokenType) lexer.Token {
	token := p.peek()

	if token.Type != typ {
		panic(fmt.Sprintf("expected type %d got %d", typ, token.Type))
	}

	return token
}

func (p *Parser) ended() bool {
	return p.pos >= len(p.tokens)
}

func (p *Parser) peek() lexer.Token {
	if p.ended() {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.pos]
}

func (p *Parser) parseNode() (ast.Node, error) {
	token := p.expectType(lexer.Keyword)
	p.advance()

	switch token.Literal {
	case "while":
		return p.parseWhile()
	case "incr":
		return p.parseIncr()
	case "decr":
		return p.parseDecr()
	case "clear":
		return p.parseClear()
	}

	return nil, errors.New("unknown keyword")
}

func (p *Parser) parseWhile() (*ast.While, error) {
	node := ast.While{}

	token := p.peek()
	if token.Type == lexer.Identifer {
		node.Var = &ast.Identifer{Name: token.Literal}
	} else if token.Type == lexer.Integer {
		num, err := strconv.Atoi(token.Literal)
		if err != nil {
			panic(err)
		}
		node.Var = &ast.IntegerLiteral{Num: num}
	} else {
		panic("expected identifier or integer")
	}
	p.advance()

	p.expectLiteral("not")
	p.advance()

	token = p.peek()
	if token.Type == lexer.Identifer {
		node.Not = &ast.Identifer{Name: token.Literal}
	} else if token.Type == lexer.Integer {
		num, err := strconv.Atoi(token.Literal)
		if err != nil {
			panic(err)
		}
		node.Not = &ast.IntegerLiteral{Num: num}
	} else {
		panic("expected identifier or integer")
	}
	p.advance()

	p.expectLiteral("do")
	p.advance()

	p.expectLiteral(";")
	p.advance()

	for p.peek().Literal != "end" {
		subnode, err := p.parseNode()
		if err != nil {
			panic(err)
		}

		node.Body = append(node.Body, subnode)
	}
	p.advance()

	p.expectLiteral(";")
	p.advance()

	return &node, nil
}

func (p *Parser) parseIncr() (*ast.Incr, error) {
	token := p.expectType(lexer.Identifer)
	p.advance()

	p.expectLiteral(";")
	p.advance()

	return &ast.Incr{Var: token.Literal}, nil
}

func (p *Parser) parseDecr() (*ast.Decr, error) {
	token := p.expectType(lexer.Identifer)
	p.advance()

	p.expectLiteral(";")
	p.advance()

	return &ast.Decr{Var: token.Literal}, nil
}

func (p *Parser) parseClear() (*ast.Clear, error) {
	token := p.expectType(lexer.Identifer)
	p.advance()

	p.expectLiteral(";")
	p.advance()

	return &ast.Clear{Var: token.Literal}, nil
}

func ParseNodes(tokens []lexer.Token) []ast.Node {
	nodes := []ast.Node{}
	parser := Parser{tokens, 0}

	for !parser.ended() {
		node, err := parser.parseNode()
		if err != nil {
			panic(err)
		}

		nodes = append(nodes, node)
	}

	return nodes
}
