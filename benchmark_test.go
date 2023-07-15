package fhirpath

import (
	_ "embed"
	"testing"
)

//go:embed sample/bundle_order.json
var fhirBundleOrder string

func BenchmarkEvaluate_Advanced(b *testing.B) {
	for runIndex := 0; runIndex < b.N; runIndex++ {
		Evaluate[string](fhirBundleOrder, "Bundle.entry.where(resource.resourceType='ServiceRequest').resource.code.coding.code")
	}
}
