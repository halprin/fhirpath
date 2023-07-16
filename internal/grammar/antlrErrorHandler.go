package grammar

import (
	"errors"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type AntlrErrorHanderExtractor struct {
	antlr.DefaultErrorListener
	errors []error
}

func (receiver *AntlrErrorHanderExtractor) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, message string, _ antlr.RecognitionException) {
	newError := errors.New(message)
	receiver.errors = append(receiver.errors, newError)
}

func (receiver *AntlrErrorHanderExtractor) Error() error {

	if len(receiver.errors) == 0 {
		return nil
	}

	errorMessageBuilder := strings.Builder{}

	for errorIndex, err := range receiver.errors {
		errorMessageBuilder.WriteString(err.Error())

		if errorIndex == len(receiver.errors)-1 {
			//don't append an ending newline
			continue
		}

		errorMessageBuilder.WriteRune('\n')
	}

	return errors.New(errorMessageBuilder.String())
}
