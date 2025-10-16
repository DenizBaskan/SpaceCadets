package lexer

import (
	"slices"
	"strconv"
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
	EOF       TokenType = "EOF"
)

var keywords = []string{"clear", "incr", "decr", "while", "not", "do", "end", "copy", "to"}

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

		case '\r':
			i++

		case ' ':
			i++

		case ';':
			tokens = append(tokens, Token{Symbol, ";", line})
			i++

		default:
			for _, keyword := range keywords {
				if strings.HasPrefix(data[i:], keyword) {
					if i+len(keyword) < len(data) && !slices.Contains([]string{" ", ";"}, string(data[i+len(keyword)])) {
						// character after keyword is not a space or a comma thus not a keyword
						continue
					}

					tokens = append(tokens, Token{Keyword, keyword, line})
					i += len(keyword)
					continue Out
				}
			}

			s := ""

			// gather a string until the end of the file or until a whitespace or ;
			for i < len(data) && data[i] != ';' && data[i] != ' ' {
				s += string(data[i])
				i++
			}

			if validIdentifer(s) {
				tokens = append(tokens, Token{Identifer, s, line})
			} else if _, err := strconv.Atoi(s); err == nil {
				tokens = append(tokens, Token{Integer, s, line})
			} else {
				tokens = append(tokens, Token{Illegal, string(s), line})
			}

			i += len(s)
		}
	}

	return append(tokens, Token{EOF, "", line})
}

func validIdentifer(s string) bool {
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
