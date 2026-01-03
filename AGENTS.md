# FHIRPath Evaluator Library for Go

This repository contains a Go library for evaluating [FHIR Path](https://hl7.org/fhirpath/) expressions against [FHIR](http://hl7.org/fhir/) resources.

## Architecture Overview

This library uses ANTLR4 for lexical analysis and parsing. The grammar specification is located at `internal/grammar/antlrGen/Fhirpath.g4` and follows the [FHIRPath Normative Release](http://hl7.org/fhirpath/N1) specification.

### Key Components

- Main Entry Point: `evaluate.go` - Contains the `Evaluate[T]` function, the primary API for evaluating FHIRPath expressions.
- Parser: Uses ANTLR-generated parser from the grammar specification.
- Engine: `internal/engine/` - Contains the evaluation engine that walks the ANTLR parse tree.

### Engine Structure

The files in `internal/engine/` correspond directly to rules in the ANTLR grammar (`Fhirpath.g4`). Each file implements the evaluation logic for its corresponding grammar rule. Sometimes there are helper functions in other files than listed above that should be called from other grammar engine rules.

- `dynamicValue.go` - Runtime value representation.
- `engine.go` - Core engine initialization and coordination.
- `utils.go` - Utility functions.

Sometimes there are helper functions in other files than listed above that should be called from other grammar engine rules.

## How to Contribute

When adding new features or fixing bugs:

- Grammar Changes: If modifying FHIRPath language support, update `Fhirpath.g4` first and regenerate the parser. This will never happen in the context of AI.
- Engine Files: Create or modify files in `internal/engine/` that correspond to the grammar rules being implemented.
- Naming Convention: File names should match the grammar rule names (e.g., `memberInvocation` rule `MemberInvocation.go`)
- Testing: Add test cases to `evaluate_test.go` and run (but do not add to) the official FHIRPath test suite in `official_test.go`
