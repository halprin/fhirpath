package lex

type Token struct {
	Type    int
	Literal string
}

const (
	PERIOD int = iota

	ALPHA
	NUMERIC

	WHITE_SPACE

	PARENTHESIS_START
	PARENTHESIS_END

	AT_SIGN

	QUOTE

	PLUS
	DASH
	STAR
	SLASH

	BACK_SLASH
)
