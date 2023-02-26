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

	TRUE
	FALSE

	PARENTHESIS_START
	PARENTHESIS_END

	AND
	OR

	AT_SIGN

	QUOTE

	PLUS
	DASH
	STAR
	SLASH

	BACK_SLASH
)
