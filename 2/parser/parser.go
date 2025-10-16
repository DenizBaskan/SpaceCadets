package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
)

type Parser struct {
	tokens []lexer.Token
	index  int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, index: 0}
}

func (p *Parser) Parse() []ast.Node {
	nodes := []ast.Node{}

	for p.index < len(p.tokens) && p.tokens[p.index].Type != lexer.EOF {
		node := p.parseStatement()
		nodes = append(nodes, node)
	}

	return nodes
}
