package lexer

import (
	"slices"
	"strings"
	"unicode"
)

type TokenType uint8

const (
	Keyword TokenType = iota
	Identifer
	Symbol
	Integer
	Illegal
)

var keywords = []string{"clear", "incr", "decr", "while", "not", "do", "end"}

type Token struct {
	Type    TokenType
	Literal string
}

func Tokenize(data string) []Token {
	var tokens []Token

	data = strings.ReplaceAll(data, "\n", "")
	data = strings.ReplaceAll(data, "\r", "")

	i := 0
Out:
	for i < len(data) {
		switch data[i] {
		case ' ':
			i++
			continue
		case ';':
			i++
			tokens = append(tokens, Token{Symbol, ";"})
			continue
		default:
			for _, keyword := range keywords {
				if strings.HasPrefix(data[i:], keyword) && slices.Contains([]string{" ", ";"}, string(data[i+len(keyword)])) {
					tokens = append(tokens, Token{Keyword, keyword})
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
				tokens = append(tokens, Token{Integer, num})
				continue
			}

			if unicode.IsLetter(rune(data[i])) {
				s := ""
				for unicode.IsNumber(rune(data[i])) || unicode.IsLetter(rune(data[i])) || data[i] == '_' {
					s += string(data[i])
					i++
				}
				tokens = append(tokens, Token{Identifer, s})
				continue
			}

			tokens = append(tokens, Token{Illegal, string(data[i])})
			i++
		}
	}

	return tokens
}
