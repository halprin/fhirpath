package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
	"strconv"
	"strings"
)

func (receiver *engine) NumberLiteral(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	//remove the start and end quotes
	literal := node.Text()
	
	if strings.ContainsRune(literal, '.') {
		return strconv.ParseFloat(literal, 64)
	}
	
	return strconv.Atoi(literal)
}
