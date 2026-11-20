package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownReadModel = "esdm/structure/slice-references-unknown-read-model"

type sliceReferencesUnknownReadModelRule struct{}

func newSliceReferencesUnknownReadModelRule() *sliceReferencesUnknownReadModelRule {
	return &sliceReferencesUnknownReadModelRule{}
}

func (*sliceReferencesUnknownReadModelRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownReadModel,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The read model a Slice names must be declared in the Slice's Bounded Context. A Slice that shows or consumes a read model the model does not declare shows nothing.",
	}
}

func (*sliceReferencesUnknownReadModelRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		node := slice.ReadModel()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupReadModel(domain, boundedContext, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references read-model %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
			Location: node.Location(),
		})
	}
}
