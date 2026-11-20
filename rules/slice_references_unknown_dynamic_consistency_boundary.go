package rules

import (
	"context"
	"fmt"

	"github.com/thenativeweb/esdm/diag"
	"github.com/thenativeweb/esdm/model"
)

const ruleIDSliceReferencesUnknownDynamicConsistencyBoundary = "esdm/structure/slice-references-unknown-dynamic-consistency-boundary"

type sliceReferencesUnknownDynamicConsistencyBoundaryRule struct{}

func newSliceReferencesUnknownDynamicConsistencyBoundaryRule() *sliceReferencesUnknownDynamicConsistencyBoundaryRule {
	return &sliceReferencesUnknownDynamicConsistencyBoundaryRule{}
}

func (*sliceReferencesUnknownDynamicConsistencyBoundaryRule) Meta() Meta {
	return Meta{
		ID:          ruleIDSliceReferencesUnknownDynamicConsistencyBoundary,
		Extension:   "event-modeling",
		Severity:    diag.SeverityError,
		Description: "The dynamic consistency boundary a Slice names must be declared in the Slice's Bounded Context. Event Modeling needs no Aggregate, so the boundary is where a Slice's command is decided; without it the command has no home.",
	}
}

func (*sliceReferencesUnknownDynamicConsistencyBoundaryRule) Check(ctx context.Context, m *model.Model, report diag.Reporter) {
	for _, slice := range sortedByName(m.Extensions.EventModeling.Slices) {
		domain, boundedContext := sliceBoundedContext(slice.Scope())
		node := slice.DynamicConsistencyBoundary()
		name, _ := node.Text()
		if name == "" {
			continue
		}
		_, ok := m.LookupDynamicConsistencyBoundary(domain, boundedContext, name)
		if ok {
			continue
		}
		sliceName, _ := slice.Name().Text()
		report.Report(diag.Diagnostic{
			Message:  fmt.Sprintf("slice %q references dynamic-consistency-boundary %q which is not declared in bounded-context %q", sliceName, name, boundedContext),
			Location: node.Location(),
		})
	}
}
