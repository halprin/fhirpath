#!/bin/sh

rm ./*.interp
rm ./*.tokens
rm ./fhirpath*.go

alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package antlrgen *.g4