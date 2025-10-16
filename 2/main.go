package main

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must specify a source file path")
		os.Exit(1)
	}

	path := os.Args[1]

	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	tokens := lexer.Tokenize(string(b))
	fmt.Println("Created token stream")

	p := parser.NewParser(tokens)
	nodes := p.Parse()
	fmt.Println("Parsing complete")

	env := make(map[string]int)
	ast.ExecuteNodes(nodes, env)
}
