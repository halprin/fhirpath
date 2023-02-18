package lex

import (
	"bufio"
	"bytes"
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
	default:
		return receiver.checkForMultiCharacterToken(character)
	}
}

func (receiver Lexer) read() (rune, error) {
	character, _, err := receiver.inputReader.ReadRune()
	return character, err
}

func (receiver Lexer) checkForMultiCharacterToken(character rune) (Token, error) {
	if isWhitespace(character) {
		literal, err := receiver.scanUntilNot(isWhitespace, character)
		return Token{Type: WHITE_SPACE, Literal: literal}, err
	} else if isNumeric(character) {
		literal, err := receiver.scanUntilNot(isNumeric, character)
		return Token{Type: NUMBER, Literal: literal}, err
	} else if isAlpha(character) {
		//if we start with an alpha, scan all the alpha-numeric stuff
		literal, err := receiver.scanUntilNot(func(innerCharacter rune) bool {
			return isAlpha(innerCharacter) || isNumeric(innerCharacter)
		}, character)

		if literal == "true" {
			return Token{Type: TRUE, Literal: literal}, err
		} else if literal == "false" {
			return Token{Type: FALSE, Literal: literal}, err
		} else if literal == "and" {
			return Token{Type: AND, Literal: literal}, err
		} else if literal == "or" {
			return Token{Type: OR, Literal: literal}, err
		}

		return Token{Type: ALPHA_NUMERIC, Literal: literal}, err
	}

	return Token{}, errors.New("no token found")
}

func (receiver Lexer) scanUntilNot(characterClass func(rune) bool, initialCharacter rune) (string, error) {
	var buffer bytes.Buffer
	var character rune
	var err error

	buffer.WriteRune(initialCharacter)

	for {
		character, err = receiver.read()
		if err != nil {
			break
		} else if !characterClass(character) {
			err = receiver.inputReader.UnreadRune()
			break
		}

		buffer.WriteRune(character)
	}

	if err == io.EOF && buffer.Len() > 0 {
		//we encountered the end, but we've read something means we have something valid; hide the error because it will be encountered in the next read
		err = nil
	}

	return buffer.String(), err
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
