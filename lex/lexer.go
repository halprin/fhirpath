package lex

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Lexer struct {
	input       string
	inputReader *bufio.Reader
}

func NewLexer(input string) Lexer {
	return Lexer{
		input:       input,
		inputReader: bufio.NewReader(strings.NewReader(input)),
	}
}

func (receiver Lexer) Lex() ([]Token, error) {
	var tokens []Token
	var token Token
	var err error

	for err == nil {
		token, err = receiver.NextToken()
		if err != nil {
			if err == io.EOF {
				err = nil // we don't need to bubble up EOF, just finish the lexing
			}
			break
		}

		tokens = append(tokens, token)
	}

	return tokens, err
}

func (receiver Lexer) NextToken() (Token, error) {
	character, err := receiver.read()
	if err != nil {
		return Token{}, err
	}

	switch character {
	case '.':
		return Token{Type: PERIOD, Literal: string(character)}, nil
	case '(':
		return Token{Type: PARENTHESIS_START, Literal: string(character)}, nil
	case ')':
		return Token{Type: PARENTHESIS_END, Literal: string(character)}, nil
	case '@':
		return Token{Type: AT_SIGN, Literal: string(character)}, nil
	case '\'':
		return Token{Type: QUOTE, Literal: string(character)}, nil
	case '+':
		return Token{Type: PLUS, Literal: string(character)}, nil
	case '-':
		return Token{Type: DASH, Literal: string(character)}, nil
	case '*':
		return Token{Type: STAR, Literal: string(character)}, nil
	case '/':
		return Token{Type: SLASH, Literal: string(character)}, nil
	case '\\':
		return Token{Type: BACK_SLASH, Literal: string(character)}, nil
	default:
		return receiver.checkForTokenClass(character)
	}
}

func (receiver Lexer) read() (rune, error) {
	character, _, err := receiver.inputReader.ReadRune()
	return character, err
}

func (receiver Lexer) checkForTokenClass(character rune) (Token, error) {
	if isWhitespace(character) {
		return Token{Type: WHITE_SPACE, Literal: string(character)}, nil
	} else if isNumeric(character) {
		return Token{Type: NUMERIC, Literal: string(character)}, nil
	} else if isAlpha(character) {
		return Token{Type: ALPHA, Literal: string(character)}, nil
	}

	return Token{}, errors.New("no token found")
}

func isWhitespace(character rune) bool {
	return character == ' ' || character == '\t' || character == '\n' || character == '\r'
}

func isAlpha(character rune) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
}

func isNumeric(character rune) bool {
	return character >= '0' && character <= '9'
}
