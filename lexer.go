package main

import (
	"unicode"
)

type TokenType int

const (
	TokenInt TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenLParen     // Left Parenthesis
	TokenRParen     // Right Parenthesis
	TokenEqual      // Equal sign (=)
	TokenSemicolon  // Semicolon (;)
	TokenIdentifier // Identifiers (variable names)
	TokenFunction   // Function names (e.g., "sin", "cos")
	TokenExponent   // Exponent (^)
	TokenNegation   // Negation (!)
	TokenEOF
	TokenIllegal
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input        string
	position     int
	currentToken Token
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.nextToken()
	return lexer
}

func (l *Lexer) nextToken() {
	if l.position >= len(l.input) {
		l.currentToken = Token{Type: TokenEOF, Value: ""}
		return
	}

	currentChar := l.input[l.position]

	switch currentChar {
	case '+':
		l.currentToken = Token{Type: TokenPlus, Value: "+"}
	case '-':
		l.currentToken = Token{Type: TokenMinus, Value: "-"}
	case '*':
		l.currentToken = Token{Type: TokenMultiply, Value: "*"}
	case '/':
		l.currentToken = Token{Type: TokenDivide, Value: "/"}
	case '(':
		l.currentToken = Token{Type: TokenLParen, Value: "("}
	case ')':
		l.currentToken = Token{Type: TokenRParen, Value: ")"}
	case '=':
		l.currentToken = Token{Type: TokenEqual, Value: "="}
	case ';':
		l.currentToken = Token{Type: TokenSemicolon, Value: ";"}
	case '!':
		l.currentToken = Token{Type: TokenNegation, Value: "!"}
	case '^':
		l.currentToken = Token{Type: TokenExponent, Value: "^"}
	default:
		if unicode.IsDigit(rune(currentChar)) {
			start := l.position
			for l.position < len(l.input) && unicode.IsDigit(rune(l.input[l.position])) {
				l.position++
			}
			l.currentToken = Token{Type: TokenInt, Value: l.input[start:l.position]}
			return
		} else if unicode.IsLetter(rune(currentChar)) {
			start := l.position
			for l.position < len(l.input) && (unicode.IsLetter(rune(l.input[l.position])) || unicode.IsDigit(rune(l.input[l.position]))) {
				l.position++
			}
			l.currentToken = Token{Type: TokenIdentifier, Value: l.input[start:l.position]}
			return
		} else if unicode.IsSpace(rune(currentChar)) {
			l.position++
			l.nextToken()
			return
		} else {
			// Invalid token
			l.currentToken = Token{Type: TokenIllegal, Value: string(currentChar)}
			l.position++
			return
		}
	}

	l.position++
}
