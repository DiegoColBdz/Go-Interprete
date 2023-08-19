package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una expresión: ")
	input, _ := reader.ReadString('\n')

	lexer := NewLexer(input)
	fmt.Println("Tokens:")
	for {
		token := lexer.currentToken
		if token.Type == TokenEOF {
			break
		}
		fmt.Printf("%s es Token %s\n", token.Value, tokenTypeToString(token.Type))
		lexer.nextToken()
	}
}

func tokenTypeToString(tokenType TokenType) string {
	switch tokenType {
	case TokenInt:
		return "Int"
	case TokenPlus:
		return "Plus"
	case TokenMinus:
		return "Minus"
	case TokenMultiply:
		return "Multiply"
	case TokenDivide:
		return "Divide"
	case TokenLParen:
		return "Left Parenthesis"
	case TokenRParen:
		return "Right Parenthesis"
	case TokenEqual:
		return "Equal"
	case TokenSemicolon:
		return "Semicolon"
	case TokenIdentifier:
		return "Identifier"
	case TokenFunction:
		return "Function"
	case TokenExponent:
		return "Exponent"
	case TokenNegation:
		return "Negation"
	case TokenEOF:
		return "EOF"
	case TokenIllegal:
		return "Illegal"
	default:
		return "Unknown"
	}
}
