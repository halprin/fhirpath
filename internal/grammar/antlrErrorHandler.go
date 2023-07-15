package grammar

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorHanderAppender struct {
	antlr.DefaultErrorListener
}

func (d *ErrorHanderAppender) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, message string, _ antlr.RecognitionException) {
	log.Println("ERROR")
	log.Println(message)
}
