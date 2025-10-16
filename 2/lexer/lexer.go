package lexer

import (
	"slices"
	"strconv"
	"unicode"
)

type TokenType string

const (
	Keyword   TokenType = "KEYWORD"
	Identifer TokenType = "IDENTIFIER"
	Symbol    TokenType = "SYMBOL"
	Integer   TokenType = "INTEGER"
	Illegal   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
)

var keywords = []string{"clear", "incr", "decr", "while", "not", "do", "end", "copy", "to"}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

// quite a dodgy function but works pretty well
func Tokenize(data string) []Token {
	var (
		tokens []Token
		i      = 0
		line   = 1
	)

	for i < len(data) {
		switch data[i] {
		case '\n':
			line++
			i++

		case '\r', ' ':
			i++

		case ';':
			tokens = append(tokens, Token{Symbol, ";", line})
			i++

		case '#':
			// skip the comment
			for i < len(data) && data[i] != '\n' {
				i++
			}

		default:
			// collect a string

			str := ""
			for i < len(data) && !slices.Contains([]byte{'\n', '\r', ' ', ';'}, data[i]) {
				str += string(data[i])
				i++
			}

			if slices.Contains(keywords, str) {
				tokens = append(tokens, Token{Keyword, str, line})
			} else if isIdentifier(str) {
				tokens = append(tokens, Token{Identifer, str, line})
			} else if _, err := strconv.Atoi(str); err == nil {
				tokens = append(tokens, Token{Integer, str, line})
			} else {
				tokens = append(tokens, Token{Illegal, string(str), line})
			}
		}
	}

	return append(tokens, Token{EOF, "", line})
}

func isIdentifier(s string) bool {
	if s[0] != '_' && !unicode.IsLetter(rune(s[0])) {
		return false
	}

	for _, v := range s {
		if v != '_' && unicode.IsLetter(v) && unicode.IsNumber(v) {
			return false
		}
	}

	return true
}
