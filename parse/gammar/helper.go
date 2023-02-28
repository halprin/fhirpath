package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"strings"
)

func concatTokenLiterals(tokens []lex.Token) string {
	var builder strings.Builder

	for _, token := range tokens {
		builder.WriteRune(token.Literal)
	}

	return builder.String()
}
