package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownQuery = "esdm/structure/slice-references-unknown-query"

type sliceReferencesUnknownQueryRule struct{}

func newSliceReferencesUnknownQueryRule() *sliceReferencesUnknownQueryRule {
	return &sliceReferencesUnknownQueryRule{}
}

func (*sliceReferencesUnknownQueryRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownQuery,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The query a Slice names must be declared in the Slice's Bounded Context. The query is how a screen reads the Slice's read model; without it the screen has no way to ask.",
	}
}

func (*sliceReferencesUnknownQueryRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		node := slice.Query()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupQuery(domain, boundedContext, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references query %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
			Location: node.Location(),
		})
	}
}
