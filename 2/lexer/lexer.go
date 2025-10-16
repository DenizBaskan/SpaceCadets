package lexer

import (
	"slices"
	"strings"
	"unicode"
)

type TokenType string

const (
	Keyword   TokenType = "KEYWORD"
	Identifer TokenType = "IDENTIFIER"
	Symbol    TokenType = "SYMBOL"
	Integer   TokenType = "INTEGER"
	Illegal   TokenType = "ILLEGAL"
)

var keywords = []string{"clear", "incr", "decr", "while", "not", "do", "end"}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

func Tokenize(data string) []Token {
	var tokens []Token

	var (
		i    = 0
		line = 1
	)
Out:
	for i < len(data) {
		switch data[i] {
		case '\n':
			line++
			i++
			continue
		case '\r':
			i++
			continue
		case ' ':
			i++
			continue
		case ';':
			i++
			tokens = append(tokens, Token{Symbol, ";", line})
			continue
		default:
			for _, keyword := range keywords {
				if strings.HasPrefix(data[i:], keyword) && slices.Contains([]string{" ", ";"}, string(data[i+len(keyword)])) {
					tokens = append(tokens, Token{Keyword, keyword, line})
					i += len(keyword)
					continue Out
				}
			}
			// if we have reached this point, this is NOT a keyword

			if unicode.IsNumber(rune(data[i])) {
				num := ""
				for unicode.IsNumber(rune(data[i])) {
					num += string(data[i])
					i++
				}
				tokens = append(tokens, Token{Integer, num, line})
				continue
			}

			if unicode.IsLetter(rune(data[i])) {
				s := ""
				for unicode.IsNumber(rune(data[i])) || unicode.IsLetter(rune(data[i])) || data[i] == '_' {
					s += string(data[i])
					i++
				}
				tokens = append(tokens, Token{Identifer, s, line})
				continue
			}

			tokens = append(tokens, Token{Illegal, string(data[i]), line})
			i++
		}
	}

	return tokens
}
