# fhirpath

Evaluates [FHIR Path](https://hl7.org/fhirpath/) against [FHIR](http://hl7.org/fhir/) resources.

This is a work in progress.  44 out of the 686 [official tests](https://hl7.org/fhirpath/tests.html) pass.

## Install

Require `github.com/halprin/fhirpath` in your `go.mod` file or use `go` to add it.
```shell
go get github.com/halprin/fhirpath
```

## API

Start by importing the package.
```go
import "github.com/halprin/fhirpath"
```

### Evaluate

`fhirpath.Evaluate` is the main function of the API.

It takes two arguments.

1. `fhirString` - `string`.  Contains the FHIR JSON that the FHIR path will evaluate against.
2. `fhirPath` - `string`.  The FHIR path that is evaluated.

There are two return values.

1. `[]T` - A slice of values based on the evaluation of the FHIR path against the FHIR JSON.  If the evaluation
   resulted in nothing, an empty slice is returned.  A slice of size 1 or larger is possible depending on whether the
   evaluation matched multiple values.
2. `error` - Optional.  If not `nil`, an error was generated during evaluation.

`fhirpath.Evaluate` is a generic function, so it takes a type parameter.  Upon evaluation, any results that are not the
same as the type parameter are filtered out.  If you want nothing filtered out, use `any` as the type paramter.

#### Example

```go
package main

import (
	"fmt"
	"github.com/halprin/fhirpath"
)

// see https://github.com/halprin/fhirpath/blob/main/sample/patient.json
//go:embed sample/patient.json
var fhirPatient string

func main() {
	result, err := fhirpath.Evaluate[string](fhirPatient, "Patient.identifier.where(system='http://new-republic.gov/galactic-citizen-identifier').value")
	if err != nil {
		panic("FHIR path evaluation failed")
	}

	fmt.Printf("Number of results=%d\n", len(result)) // Number of results=1
	fmt.Printf("First result=%s\n", result[0])        // First result=b531d827-de9a-4e2e-a53b-8621bd29f656
}
```
