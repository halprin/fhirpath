package engine

import (
	"github.com/halprin/fhirpath/context"
	"strconv"
	"strings"

	"github.com/halprin/fhirpath/internal/grammar"
)

// NumberLiteral converts the text to either a float64 (if the number is a decimal) or an int (if the number isn't a decimal).
func (receiver *engine) NumberLiteral(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	literal := node.Text()

	if strings.ContainsRune(literal, '.') {
		//this is a decimal number
		number, err := strconv.ParseFloat(literal, 64)
		if err != nil {
			return nil, err
		}

		return NewDynamicValue(number), nil
	}

	number, err := strconv.Atoi(literal)
	if err != nil {
		return nil, err
	}

	return NewDynamicValue(number), nil
}
