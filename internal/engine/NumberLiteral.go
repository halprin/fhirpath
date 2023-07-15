package engine

import (
	"strconv"
	"strings"

	"github.com/halprin/fhirpath/internal/grammar"
)

// NumberLiteral converts the text to either a float64 (if the number is a decimal) or an int (if the number isn't a decimal).
func (receiver *engine) NumberLiteral(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	literal := node.Text()

	if strings.ContainsRune(literal, '.') {
		//this is a decimal number
		return strconv.ParseFloat(literal, 64)
	}

	return strconv.Atoi(literal)
}
